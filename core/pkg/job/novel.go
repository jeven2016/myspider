package job

import (
	"context"
	"core/pkg/common"
	"core/pkg/job/parser"
	"core/pkg/log"
	"core/pkg/model"
	"core/pkg/service"
	"core/pkg/stream"
	"core/pkg/stream/message"
	"encoding/json"
	"github.com/reugn/go-streams/flow"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// NovelJob
// 1. load the book home page
// 2. make sure how many chapters to analyse and generate each page link into stream for next step
type NovelJob struct {
}

func (ch *NovelJob) Validate(ctx context.Context) error {
	return nil
}

func (ch *NovelJob) Run(ctx context.Context) {
	bookParser := ctx.Value(common.ParserCtx).(parser.Parser)

	//construct source
	redisClient, source, err := CreateSource(ctx,
		message.BookUrlStream, message.BookUrlStreamConsumer)
	if err != nil {
		return
	}

	// fetch url from stream
	paramsConvertFlow := flow.NewMap(ch.ConvertParams, common.JobParallelism)

	sink := stream.NewRedisStreamSink(ctx, redisClient, message.ChapterUrlStream)

	//get the messages and map them into model and save
	source.Via(paramsConvertFlow).
		Via(FlatMany[*message.ChapterMessage](bookParser)).
		To(sink)

}

func (ch *NovelJob) ConvertParams(jsonData string) *parser.ParseParams {
	//convert the message from map to struct
	var scMessage = &message.NovelMessage{}
	crtErr := json.Unmarshal([]byte(jsonData), scMessage)
	if crtErr != nil {
		log.Logger().Error("unknown SiteCatalogHomeMessage", zap.Any("SiteCatalogHomeMessage", jsonData))
		return &parser.ParseParams{Url: ""}
	}

	site := service.GetSiteByKey(scMessage.SiteKey)
	if site == nil {
		log.Logger().Error("no site with this key exists", zap.String("siteKey", scMessage.SiteKey))
		return nil
	}
	catalogObjectId, crtErr := primitive.ObjectIDFromHex(scMessage.CatalogId)
	if crtErr != nil {
		log.SugaredLogger().Errorf("unable convert '%s' to objectId", scMessage.CatalogId)
		return &parser.ParseParams{Url: ""}
	}
	novel := &model.Novel{SiteCatalogId: catalogObjectId, SiteCatalogName: scMessage.CatalogName}

	//pass params to next step
	var params = &parser.ParseParams{Url: scMessage.NovelLink, SiteKey: scMessage.SiteKey, Payload: novel}
	return params
}
