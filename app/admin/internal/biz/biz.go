package biz

import (
	"context"

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
