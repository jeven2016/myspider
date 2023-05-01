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
	"strings"
)

// HomeJob
// 1. parse the home page and get catalog links
// 2. send the catalog link messages into redis stream for further process
type HomeJob struct {
}

func (hj *HomeJob) Validate(ctx context.Context) error {
	site := ctx.Value(common.SiteCtx).(*model.Site)
	if strings.TrimSpace(site.HomeUrl) == "" {
		return common.InvalidHomeUrlErr
	}

	return nil
}

func (hj *HomeJob) Run(ctx context.Context) {
	pageParser := ctx.Value(common.ParserCtx).(parser.Parser)

	//construct source
	redisClient, source, err := CreateSource(ctx, message.HomeUrlStream, message.HomeUrlStreamConsumer)
	if err != nil {
		return
	}

	// fetch url from stream
	paramsConvertFlow := flow.NewMap(hj.ConvertParams, common.JobParallelism)

	//convert to messages slice
	flatMap := FlatMany[*message.SiteCatalogHomeMessage](pageParser)

	sink := stream.NewRedisStreamSink(ctx, redisClient, message.SiteCatalogHomeUrlStream)

	//chain all functions
	source.
		Via(paramsConvertFlow).
		Via(flatMap).
		To(sink)

}

func (hj *HomeJob) ConvertParams(jsonData string) *parser.ParseParams {
	var homeUrl string

	//convert the message type from map to struct
	var homeMsg = &message.HomeMessage{}
	crtErr := json.Unmarshal([]byte(jsonData), homeMsg)
	if crtErr != nil {
		log.Logger().Error("unknown HomeMessage", zap.String("homeMessage", jsonData))
		return &parser.ParseParams{Url: homeUrl}
	}

	site := service.GetSiteByKey(homeMsg.SiteKey)

	if site == nil {
		log.Logger().Error("no site with this key exists", zap.String("siteKey", homeMsg.SiteKey))
	} else {
		homeUrl = site.HomeUrl
	}

	//pass params to next step
	var params = &parser.ParseParams{Url: homeUrl, SiteKey: homeMsg.SiteKey}
	return params
}
