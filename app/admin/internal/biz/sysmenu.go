package biz

import (
	"context"
	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/authz"
)

type SysMenuRepo interface {
	Create(ctx context.Context, menu *model.SysMenu) error
	Save(ctx context.Context, menu *model.SysMenu) error
	Delete(ctx context.Context, id int64) error

	FindById(ctx context.Context, id int64) (*model.SysMenu, error)
	ListAll(ctx context.Context) ([]*model.SysMenu, error)
	FindByNameStatus(ctx context.Context, name string, status int32) ([]*model.SysMenu, error)
	GetRoleMenuId(ctx context.Context, roleId int64) ([]int32, error)
	SelectMenuLabel(data model.SysMenu) ([]*pb.MenuLabel, error)
}

type SysMenuUseCase struct {
	repo SysMenuRepo
	log  *log.Helper
}

func NewSysMenusUseCase(repo SysMenuRepo, logger log.Logger) *SysMenuUseCase {
	return &SysMenuUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (m *SysMenuUseCase) CreateMenus(ctx context.Context, menu *model.SysMenu) (*model.SysMenu, error) {
	claims := authz.MustFromContext(ctx)
	menu.CreateBy = claims.Nickname

	err := m.repo.Create(ctx, menu)
	return menu, err
}

func (m *SysMenuUseCase) UpdateMenus(ctx context.Context, menu *model.SysMenu) (*model.SysMenu, error) {
	claims := authz.MustFromContext(ctx)
	menu.UpdateBy = claims.Nickname

	err := m.repo.Save(ctx, menu)
	return menu, err
}

func (m *SysMenuUseCase) DeleteMenus(ctx context.Context, id int64) error {
	return m.repo.Delete(ctx, id)
}

func (m *SysMenuUseCase) GetMenus(ctx context.Context, id int64) (*model.SysMenu, error) {
	return m.repo.FindById(ctx, id)
}

type MenuSimpleTree struct {
	MenuId   int64             `json:"menuId"`
	MenuName string            `json:"menuName"`
	Children []*MenuSimpleTree `json:"children,omitempty"`
}

type MenuTree struct {
	model.SysMenu
	Children []*MenuTree `json:"children,omitempty"`
}

func (m *SysMenuUseCase) ListByNameStatus(ctx context.Context, menuName string, status int32) ([]*model.SysMenu, error) {
	return m.repo.FindByNameStatus(ctx, menuName, status)
}

func (m *SysMenuUseCase) RoleMenuTreeSelect(ctx context.Context, req *pb.RoleMenuTreeSelectRequest) (*pb.RoleMenuTreeSelectReply, error) {
	var err error
	result, err := m.repo.SelectMenuLabel(model.SysMenu{})
	if err != nil {
		return nil, err
	}
	menuIds := make([]int32, 0)
	if req.RoleId != 0 {
		menuIds, err = m.repo.GetRoleMenuId(ctx, req.RoleId)
		if err != nil {
			return nil, err
		}
	}
	reply := &pb.RoleMenuTreeSelectReply{
		Menus:       result,
		CheckedKeys: menuIds,
	}
	return reply, err
}
