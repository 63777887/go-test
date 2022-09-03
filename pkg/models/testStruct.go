package models

import (
	"gitee.com/chunanyong/zorm"
)

var SysUserTableName string = "sys_user"

type SysUserStruct struct {
	zorm.EntityStruct
	ID         int64  `column:"id"`
	OrgId      int64  `column:"org_id"`
	Username   string `column:"username"`
	Password   string `column:"password"`
	Phone      string `column:"phone"`
	Email      string `column:"email"`
	Icon       string `column:"icon"`
	Status     int8   `column:"status"`
	CreateTime int64  `column:"create_time"`
	UpdateTime int64  `column:"update_time"`
}

func (entity *SysUserStruct) GetTableName() string {
	return SysUserTableName
}

func (entity *SysUserStruct) GetPKColumnName() string {
	return "id"
}
