package biz

import (
	"context"
	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
)

type SysRoleMenuRepo interface {
	Create(ctx context.Context, roleMenus ...*model.SysRoleMenu) error
	DeleteByRoleId(ctx context.Context, roleIDs ...int64) error
	GetPermission(ctx context.Context, roleID int64) ([]string, error)
	FindMenuByRoleId(ctx context.Context, roleID int64) ([]*model.SysMenu, error)
	SelectMenuRole(ctx context.Context, roleName string) ([]*pb.MenuTree, error)
}

type SysRoleMenuUseCase struct {
	repo SysRoleMenuRepo
	log  *log.Helper
}

func NewSysRoleMenuUseCase(repo SysRoleMenuRepo, logger log.Logger) *SysRoleMenuUseCase {
	return &SysRoleMenuUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *SysRoleMenuUseCase) CreateRoleMenus(ctx context.Context, role *model.SysRole, menuIDs []int64) error {
	roleMenus := make([]*model.SysRoleMenu, len(menuIDs))
	for i, menuID := range menuIDs {
		roleMenus[i] = &model.SysRoleMenu{
			MenuID:   menuID,
			RoleID:   role.ID,
			RoleName: role.RoleName,
		}
	}
	return r.repo.Create(ctx, roleMenus...)
}

func (r *SysRoleMenuUseCase) DeleteByRoleId(ctx context.Context, roleIDs ...int64) error {
	return r.repo.DeleteByRoleId(ctx, roleIDs...)
}

func (r *SysRoleMenuUseCase) GetPermission(ctx context.Context, roleID int64) ([]string, error) {
	return r.repo.GetPermission(ctx, roleID)
}

func (r *SysRoleMenuUseCase) FindMenuByRoleId(ctx context.Context, roleID int64) ([]*model.SysMenu, error) {
	return r.repo.FindMenuByRoleId(ctx, roleID)
}

func (r *SysRoleMenuUseCase) SelectMenuRole(ctx context.Context, roleName string) ([]*pb.MenuTree, error) {
	return r.repo.SelectMenuRole(ctx, roleName)
}
