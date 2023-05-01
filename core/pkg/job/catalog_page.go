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
)

// CatalogPageJob
type CatalogPageJob struct {
}

func (hj *CatalogPageJob) Validate(ctx context.Context) error {
	return nil
}

func (hj *CatalogPageJob) Run(ctx context.Context) {
	catlogPageParser := ctx.Value(common.ParserCtx).(parser.Parser)

	//construct source
	redisClient, source, err := CreateSource(ctx,
		message.CatalogPageUrlStream, message.CatalogPageUrlStreamConsumer)
	if err != nil {
		return
	}

	// fetch url from stream
	paramsConvertFlow := flow.NewMap(hj.ConvertParams, 1)

	sink := stream.NewRedisStreamSink(ctx, redisClient, message.BookUrlStream)

	source.Via(paramsConvertFlow).
		Via(FlatMany[*message.BookMessage](catlogPageParser)).
		To(sink)

}

func (hj *CatalogPageJob) ConvertParams(jsonData string) *parser.ParseParams {
	//convert the message from map to struct
	var scMessage = &message.SiteCatalogPageMessage{}
	crtErr := json.Unmarshal([]byte(jsonData), scMessage)
	if crtErr != nil {
		log.Logger().Error("unknown SiteCatalogPageMessage", zap.Any("SiteCatalogPageMessage", jsonData))
		return &parser.ParseParams{Url: ""}
	}

	site := service.GetSiteByKey(scMessage.SiteKey)
	if site == nil {
		log.Logger().Error("no site with this key exists", zap.String("siteKey", scMessage.SiteKey))
		return nil
	}

	//only need to pass the url of catalog page to analyse the books
	// TODO: the catalog id is required to associate with the books analysed, basically it should be loaded from cache/storage
	var params = &parser.ParseParams{Url: scMessage.CatalogPageLink, SiteKey: scMessage.SiteKey,
		Payload: &model.SiteCatalog{}}

	return params
}
