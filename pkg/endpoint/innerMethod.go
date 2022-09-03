package endpoint

import (
	"context"
	"github.com/go-kit/log"
	"jwk/test/pkg/models"
	"jwk/test/pkg/service"
	log_ "log"
)

// 初始化服务
// 1、创建服务结构体
// 2、加载日志中间件
func ServerNew(logger log.Logger) service.Service {
	var svc service.Service

	{
		svc = NewSysUserService()
		svc = AddLogMiddleware(logger)(svc)
	}

	return svc
}

// 服务主结构体，命名规则为：服务名+Service
type InnerService struct{}

// 创建服务结构体，命名规则为：New+服务名+Service
func NewSysUserService() service.Service {
	return InnerService{}
}

func (s InnerService) GetSysUser(_ context.Context, id int64) (models.SysUserStruct, error) {
	log_.Println("GetSysUser start")
	var resp models.SysUserStruct
	return resp, nil
}
