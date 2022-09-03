package service

import (
	"context"
	"jwk/test/pkg/models"
)

type Service interface {
	GetSysUser(ctx context.Context, id int64) (models.SysUserStruct, error)
	UpdateSysUser(ctx context.Context, id int64) error
}
