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
	"go.uber.org/zap"
	"time"
)

// CatalogHomeJob
// 1. load the home page of one catalog
// 2. make sure how many pages to analyse and generate each page link into stream for next step
type CatalogHomeJob struct {
}

func (ch *CatalogHomeJob) Validate(ctx context.Context) error {
	return nil
}

func (ch *CatalogHomeJob) Run(ctx context.Context) {
	pageCountParser := ctx.Value(common.ParserCtx).(parser.Parser)

	//construct source
	redisClient, source, err := CreateSource(ctx,
		message.SiteCatalogHomeUrlStream, message.SiteCatalogHomeUrlStreamConsumer)
	if err != nil {
		return
	}

	// fetch url from stream
	paramsConvertFlow := flow.NewMap(ch.ConvertParams, 1)

	sink := stream.NewRedisStreamSink(ctx, redisClient, message.CatalogPageUrlStream)

	source.Via(paramsConvertFlow).
		Via(FlatMany[*message.SiteCatalogPageMessage](pageCountParser)).
		To(sink)

}

func (ch *CatalogHomeJob) ConvertParams(jsonData string) *parser.ParseParams {

	//convert the message from map to struct
	var scMessage = &message.SiteCatalogHomeMessage{}
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

	siteCatalog := &model.SiteCatalog{Name: scMessage.Name, ParentId: site.Id,
		CreatedDate: time.Now()}

	//pass params to next step
	var params = &parser.ParseParams{Url: scMessage.CatalogLink, SiteKey: scMessage.SiteKey, Payload: siteCatalog}
	return params
}

func (ch *CatalogHomeJob) parseLinks(pageCountParser parser.Parser, j *parser.ParseParams) any {
	result := pageCountParser.Parse(j)
	return result.Payload
}
