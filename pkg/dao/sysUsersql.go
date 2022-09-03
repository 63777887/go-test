package dao

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"go.uber.org/zap"
	"jwk/test/pkg/driver"
	"jwk/test/pkg/models"
	"jwk/test/utils"
)

var ctx = context.Background()

func GetContext() (ctx context.Context) {
	ctx1, _ := driver.Db.BindContextDBConnection(context.Background())
	return ctx1
}
func GetSysUserSql(id int64) (models.SysUserStruct, error) {
	resp := &models.SysUserStruct{}
	finder := zorm.NewSelectFinder(models.SysUserTableName)
	finder.Append(" where id = ?", id)
	_, err := zorm.QueryRow(ctx, finder, resp)
	if err != nil {
		utils.Logger.Error("GetSysUserSql Error, Err: ", zap.Error(err))
		return models.SysUserStruct{}, err
	}
	return *resp, nil
}

func UpdateSysUserSql(id int64) error {
	getContext := GetContext()
	zorm.Transaction(getContext, func(ctx context.Context) (interface{}, error) {
		sysUser := &models.SysUserStruct{}
		sysUser.ID = id
		sysUser.Status = -1
		zorm.UpdateNotZeroValue(getContext, sysUser)
		return nil, nil
	})
	return nil
}
