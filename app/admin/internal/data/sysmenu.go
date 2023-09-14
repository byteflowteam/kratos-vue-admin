package data

import (
	"context"
	"fmt"
	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/pkg/util"
	"github.com/go-kratos/kratos/v2/log"
)

type MenuIdList struct {
	ID int64 `json:"id"`
}

type sysMenuRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysMenuRepo(data *Data, logger log.Logger) biz.SysMenuRepo {
	return &sysMenuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m *sysMenuRepo) Create(ctx context.Context, menu *model.SysMenu) error {
	q := m.data.Query(ctx).SysMenu
	return q.WithContext(ctx).Create(menu)
}

func (m *sysMenuRepo) Save(ctx context.Context, menu *model.SysMenu) error {
	q := m.data.Query(ctx).SysMenu
	return q.WithContext(ctx).Save(menu)
}

func (m *sysMenuRepo) Delete(ctx context.Context, id int64) error {
	q := m.data.Query(ctx).SysMenu
	_, err := q.WithContext(ctx).Where(q.ID.Eq(id)).Delete()
	return err
}

func (m *sysMenuRepo) FindById(ctx context.Context, id int64) (*model.SysMenu, error) {
	q := m.data.Query(ctx).SysMenu
	return q.WithContext(ctx).Where(q.ID.Eq(id)).First()
}

func (m *sysMenuRepo) ListAll(ctx context.Context) ([]*model.SysMenu, error) {
	q := m.data.Query(ctx).SysMenu
	return q.WithContext(ctx).Find()
}

func (m *sysMenuRepo) FindByNameStatus(ctx context.Context, name string, status int32) ([]*model.SysMenu, error) {
	q := m.data.Query(ctx).SysMenu
	db := q.WithContext(ctx)
	if name != "" {
		db = db.Where(q.MenuName.Like(fmt.Sprintf("%%%s%%", name)))
	}
	if status != 0 {
		db = db.Where(q.Status.Eq(status))
	}
	return db.Order(q.Sort).Find()
}

// GetRoleMenuId 获取角色对应的菜单ids
func (m *sysMenuRepo) GetRoleMenuId(ctx context.Context, roleId int64) ([]int32, error) {
	menuIds := make([]int32, 0)
	menuList := make([]MenuIdList, 0)

	err := m.data.db.WithContext(ctx).
		Table("sys_menus").
		Select("sys_menus.id").
		Joins("LEFT JOIN sys_role_menus on sys_role_menus.menu_id=sys_menus.id").Where("sys_role_menus.role_id = ? ", roleId).
		Where("sys_menus.id not in (select sys_menus.parent_id from sys_menus LEFT JOIN sys_role_menus on sys_menus.id=sys_role_menus.menu_id where sys_role_menus.role_id =? )", roleId).
		Find(&menuList).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(menuList); i++ {
		menuIds = append(menuIds, int32(menuList[i].ID))
	}
	return menuIds, nil
}

func (m *sysMenuRepo) SelectMenuLabel(data model.SysMenu) ([]*pb.MenuLabel, error) {
	menuList, err := m.FindList(data)
	if err != nil {
		return nil, err
	}

	redData := make([]*pb.MenuLabel, 0)
	ml := *menuList
	for i := 0; i < len(ml); i++ {
		if ml[i].ParentID != 0 {
			continue
		}
		e := &pb.MenuLabel{}
		e.MenuId = int32(ml[i].ID)
		e.MenuName = ml[i].MenuName
		menusInfo := DiguiMenuLabel(menuList, e)

		redData = append(redData, menusInfo)
	}
	return redData, err
}

func (m *sysMenuRepo) FindList(data model.SysMenu) (*[]model.SysMenu, error) {
	list := make([]model.SysMenu, 0)

	db := m.data.db.Table("sys_menus")
	// 此处填写 where参数判断
	if data.MenuName != "" {
		db = db.Where("menu_name like ?", "%"+data.MenuName+"%")
	}
	if data.Path != "" {
		db = db.Where("path = ?", data.Path)
	}
	if data.MenuType != "" {
		db = db.Where("menu_type = ?", data.MenuType)
	}
	if data.Title != "" {
		db = db.Where("title like ?", "%"+data.Title+"%")
	}
	if data.Status != 0 {
		db = db.Where("status = ?", data.Status)
	}
	db.Where("deleted_at IS NULL")
	err := db.Order("sort").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func DiguiMenu(menulist []*model.SysMenu, menu *pb.MenuTree) *pb.MenuTree {
	list := menulist

	min := make([]*pb.MenuTree, 0)
	for j := 0; j < len(list); j++ {

		if menu.MenuId != list[j].ParentID {
			continue
		}
		mi := &pb.MenuTree{}
		mi.MenuId = list[j].ID
		mi.MenuName = list[j].MenuName
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.MenuType = list[j].MenuType
		mi.KeepAlive = list[j].KeepAlive
		mi.Permission = list[j].Permission
		mi.ParentId = list[j].ParentID
		mi.IsAffix = list[j].IsAffix
		mi.IsIframe = list[j].IsIframe
		mi.Link = list[j].Link
		mi.Component = list[j].Component
		mi.Sort = list[j].Sort
		mi.Status = list[j].Status
		mi.Hidden = list[j].Hidden
		mi.CreateTime = util.NewTimestamp(list[j].CreatedAt)
		mi.UpdateTime = util.NewTimestamp(list[j].UpdatedAt)
		mi.Children = []*pb.MenuTree{}

		if mi.MenuType != "F" {
			ms := DiguiMenu(menulist, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	menu.Children = min
	return menu
}

func DiguiMenuLabel(menulist *[]model.SysMenu, menu *pb.MenuLabel) *pb.MenuLabel {
	list := *menulist

	min := make([]*pb.MenuLabel, 0)
	for j := 0; j < len(list); j++ {

		if menu.MenuId != int32(list[j].ParentID) {
			continue
		}
		mi := pb.MenuLabel{}
		mi.MenuId = int32(list[j].ID)
		mi.MenuName = list[j].MenuName
		mi.Children = []*pb.MenuLabel{}
		if list[j].MenuType != "F" {
			ms := DiguiMenuLabel(menulist, &mi)
			min = append(min, ms)
		} else {
			min = append(min, &mi)
		}

	}
	menu.Children = min
	return menu
}
