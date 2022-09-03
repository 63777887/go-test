package handler

import (
	"context"
	"jwk/test/pkg/dao"
	"jwk/test/pkg/models"
	"jwk/test/pkg/service"

	//"fmt"
	log_ "log"
)

// 服务主结构体，命名规则为：服务名+Service
type TestService struct{}

// 创建服务结构体，命名规则为：New+服务名+Service
func NewTestService() service.Service {
	return TestService{}
}

func (s TestService) GetSysUser(_ context.Context, id int64) (models.SysUserStruct, error) {
	log_.Println("GetSysUser start")
	var resp models.SysUserStruct
	resp, err := dao.GetSysUserSql(id)
	if err != nil {
		return models.SysUserStruct{}, err
	}
	return resp, nil
}

func (s TestService) UpdateSysUser(ctx context.Context, id int64) error {

	err := dao.UpdateSysUserSql(id)

	if err != nil {
		return err
	}
	return err
}
