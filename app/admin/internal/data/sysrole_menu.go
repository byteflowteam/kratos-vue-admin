package data

import (
	"context"
	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/pkg/util"
	"github.com/go-kratos/kratos/v2/log"
)

/**
  菜单类型（M目录 C菜单 F按钮）
  菜单类型 (1-目录 2-菜单 3-按钮）
*/

type sysRoleMenuRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysRoleMenuRepo(data *Data, logger log.Logger) biz.SysRoleMenuRepo {
	return &sysRoleMenuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *sysRoleMenuRepo) Create(ctx context.Context, roleMenus ...*model.SysRoleMenu) error {
	q := s.data.Query(ctx).SysRoleMenu
	return q.WithContext(ctx).Create(roleMenus...)
}

func (s *sysRoleMenuRepo) DeleteByRoleId(ctx context.Context, roleIDs ...int64) error {
	q := s.data.Query(ctx).SysRoleMenu
	_, err := q.WithContext(ctx).Where(q.RoleID.In(roleIDs...)).Delete()
	return err
}

// GetPermission 查询权限标识
func (s *sysRoleMenuRepo) GetPermission(ctx context.Context, roleID int64) ([]string, error) {
	query := s.data.Query(ctx)
	roleMenu := query.SysRoleMenu
	menu := query.SysMenu

	var result []string
	err := menu.WithContext(ctx).
		Select(menu.Permission).
		LeftJoin(roleMenu, menu.ID.EqCol(roleMenu.MenuID)).
		Where(roleMenu.RoleID.Eq(roleID)).
		Where(menu.MenuType.In("C", "F")).Scan(&result)
	return result, err
}

// FindMenuByRoleId 查询菜单路径
func (s *sysRoleMenuRepo) FindMenuByRoleId(ctx context.Context, roleID int64) ([]*model.SysMenu, error) {
	query := s.data.Query(ctx)
	roleMenu := query.SysRoleMenu
	menu := query.SysMenu

	return menu.WithContext(ctx).
		LeftJoin(roleMenu, menu.ID.EqCol(roleMenu.MenuID)).
		Where(roleMenu.RoleID.Eq(roleID)).
		Where(menu.MenuType.In("M", "C")).
		Order(menu.Sort).
		Find()
}

func (s *sysRoleMenuRepo) SelectMenuRole(ctx context.Context, roleKey string) ([]*pb.MenuTree, error) {
	redData := make([]*pb.MenuTree, 0)

	menuList := s.GetMenuByRoleKey(roleKey)

	for i := 0; i < len(menuList); i++ {
		if menuList[i].ParentID != 0 {
			continue
		}
		menuTree := &pb.MenuTree{}
		_ = util.CopyStructFields(menuTree, menuList[i])
		menuTree.MenuId = menuList[i].ID

		menusInfo := DiguiMenu(menuList, menuTree)

		redData = append(redData, menusInfo)
	}
	return redData, nil
}

func (s *sysRoleMenuRepo) GetMenuByRoleKey(roleKey string) []*model.SysMenu {
	menus := make([]*model.SysMenu, 0)

	db := s.data.db.Table("sys_menus")
	db = db.Select("sys_menus.*").Joins("left join sys_role_menus on sys_role_menus.menu_id=sys_menus.id")
	db = db.Where("sys_role_menus.role_name=? and sys_menus.menu_type in ('M','C')", roleKey)
	err := db.Debug().Order("sys_menus.sort").Find(&menus).Error
	if err != nil {
		return nil
	}
	return menus
}
