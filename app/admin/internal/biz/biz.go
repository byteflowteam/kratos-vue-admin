package biz

import (
	"context"
	"time"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewSysUserUseCase,
	NewAuthUseCase,
	NewSysMenusUseCase,
	NewSysDeptUseCase,
	NewSysPostUseCase,
	NewSysApiUseCase,
	NewSysRoleUseCase,
	NewSysRoleMenuUseCase,
	NewCasbinRuleUseCase,
	NewSysDictDatumUseCase,
	NewSysDictTypeUseCase,
)

type Transaction interface {
	Transaction(context.Context, func(ctx context.Context) error) error
}

type RedisRepo interface {
	SetHashKey(context.Context, string, string, interface{}) error
	GetHashKey(context.Context, string, string) (string, error)
	DelHashKey(ctx context.Context, key string, field string) error
	GetHashLen(ctx context.Context, key string) error
	Lock(context.Context, string, interface{}, time.Duration) (bool, error)
	IncrHashKey(context.Context, string, string, int64) error
	GetHashAllKeyAndVal(context.Context, string) (map[string]string, error)
	Set(ctx context.Context, key string, value string, expire time.Duration) error
	Get(ctx context.Context, key string) string
	SRem(ctx context.Context, key string, members ...interface{}) (int64, error)
}
