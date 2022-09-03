package endpoint

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/tracing/zipkin"
	"github.com/go-kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	service "jwk/test/pkg/service"
)

type Endpoints struct {
	// endpoint.Endpoint 是一个函数变量，入参为context和request，出参为response
	// endpoint 做的事情是解析request，调service中的方法，再组织成response
	GetSysUserEndpoint endpoint.Endpoint
}

func New(service service.Service, logger log.Logger, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer) Endpoints {

	var getSysUserEndpoint_ endpoint.Endpoint

	{

		getSysUserEndpoint_ = MakeGetSysUserEndpoint(service)
		//getSysUserEndpoint_ = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(getSysUserEndpoint_)
		getSysUserEndpoint_ = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getSysUserEndpoint_)
		getSysUserEndpoint_ = opentracing.TraceServer(otTracer, "GetSysUser")(getSysUserEndpoint_)
		if zipkinTracer != nil {
			getSysUserEndpoint_ = zipkin.TraceEndpoint(zipkinTracer, "GetSysUser")(getSysUserEndpoint_)
		}
		getSysUserEndpoint_ = LoggingMiddleware(log.With(logger, "method", "GetSysUser"))(getSysUserEndpoint_)
	}

	return Endpoints{
		GetSysUserEndpoint: getSysUserEndpoint_,
	}

}

func MakeGetSysUserEndpoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(int64)
		resp, err := service.GetSysUser(ctx, req)
		return resp, err
	}
}
