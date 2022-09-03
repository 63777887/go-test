package endpoint

import (
	"context"
	"github.com/go-kit/kit/log"
	"jwk/test/pkg/models"
	"jwk/test/pkg/service"
)

type LogMiddleware func(service.Service) service.Service

func AddLogMiddleware(logger log.Logger) LogMiddleware {
	return func(next service.Service) service.Service {
		return logMiddleware{logger, next}
	}
}

type logMiddleware struct {
	logger log.Logger
	next   service.Service
}

func (mw logMiddleware) GetSysUser(ctx context.Context, id int64) (models.SysUserStruct, error) {
	defer func() {
		//mw.logger.Log("method", "Sum", "a", req.A, "b", req.B, "v", resp.Value, "err", err)
	}()
	return mw.next.GetSysUser(ctx, id)
}
