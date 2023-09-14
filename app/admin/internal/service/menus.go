package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"strings"

	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/authz"
	"github.com/byteflowteam/kratos-vue-admin/pkg/util"
)

type MenusService struct {
	pb.UnimplementedMenusServer
	menuUseCase     *biz.SysMenuUseCase
	roleMenuUseCase *biz.SysRoleMenuUseCase
	log             *log.Helper
}

func NewMenusService(menuUseCase *biz.SysMenuUseCase, roleMenuUseCase *biz.SysRoleMenuUseCase, logger log.Logger) *MenusService {
	return &MenusService{
		menuUseCase:     menuUseCase,
		roleMenuUseCase: roleMenuUseCase,
		log:             log.NewHelper(log.With(logger, "module", "service/menus")),
	}
}

func (s *MenusService) CreateMenus(ctx context.Context, req *pb.CreateMenusRequest) (*pb.CreateMenusReply, error) {

	_, err := s.menuUseCase.CreateMenus(ctx, &model.SysMenu{
		MenuName:   req.MenuName,
		Title:      req.Title,
		ParentID:   int64(req.ParentId),
		Sort:       req.Sort,
		Icon:       req.Icon,
		Path:       req.Path,
		Component:  req.Component,
		IsIframe:   req.IsIframe,
		Link:       req.IsLink,
		MenuType:   req.MenuType,
		Hidden:     int32(req.IsHide),
		KeepAlive:  req.IsKeepAlive,
		IsAffix:    req.IsAffix,
		Permission: req.Permission,
		Status:     req.Status,
		Remark:     req.Remark,
	})
	if err != nil {
		return nil, err
	}
	claims := authz.MustFromContext(ctx)

	permissions, err := s.roleMenuUseCase.GetPermission(ctx, claims.RoleID)
	if err != nil {
		return nil, err
	}
	menus, err := s.roleMenuUseCase.FindMenuByRoleId(ctx, claims.RoleID)
	if err != nil {
		return nil, err
	}

	return &pb.CreateMenusReply{
		Menus:       convertToMenuTree(menus),
		Permissions: permissions,
	}, nil
}
func (s *MenusService) UpdateMenus(ctx context.Context, req *pb.UpdateMenusRequest) (*pb.UpdateMenusReply, error) {

	_, err := s.menuUseCase.UpdateMenus(ctx, &model.SysMenu{
		ID:         req.MenuId,
		MenuName:   req.MenuName,
		Title:      req.Title,
		ParentID:   int64(req.ParentId),
		Sort:       req.Sort,
		Icon:       req.Icon,
		Path:       req.Path,
		Component:  req.Component,
		IsIframe:   req.IsIframe,
		Link:       req.IsLink,
		MenuType:   req.MenuType,
		Hidden:     int32(req.IsHide),
		KeepAlive:  req.IsKeepAlive,
		IsAffix:    req.IsAffix,
		Permission: req.Permission,
		Status:     req.Status,
		Remark:     req.Remark,
	})
	if err != nil {
		return nil, err
	}
	claims := authz.MustFromContext(ctx)

	permissions, err := s.roleMenuUseCase.GetPermission(ctx, claims.RoleID)
	if err != nil {
		return nil, err
	}
	menus, err := s.roleMenuUseCase.FindMenuByRoleId(ctx, claims.RoleID)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateMenusReply{
		Menus:       convertToMenuTree(menus),
		Permissions: permissions,
	}, nil
}
func (s *MenusService) DeleteMenus(ctx context.Context, req *pb.DeleteMenusRequest) (*pb.DeleteMenusReply, error) {
	if err := s.menuUseCase.DeleteMenus(ctx, req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteMenusReply{}, nil
}
func (s *MenusService) GetMenus(ctx context.Context, req *pb.GetMenusRequest) (*pb.GetMenusReply, error) {
	menu, err := s.menuUseCase.GetMenus(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetMenusReply{
		MenuId:      int32(menu.ID),
		MenuName:    menu.MenuName,
		Title:       menu.Title,
		ParentId:    int32(menu.ParentID),
		Sort:        menu.Sort,
		Icon:        menu.Icon,
		Path:        menu.Path,
		Component:   menu.Component,
		IsIframe:    menu.IsIframe,
		IsLink:      menu.Link,
		MenuType:    menu.MenuType,
		IsHide:      int64(menu.Hidden),
		IsKeepAlive: menu.KeepAlive,
		IsAffix:     menu.IsAffix,
		Permission:  menu.Permission,
		Status:      menu.Status,
		CreateBy:    menu.CreateBy,
		UpdateBy:    menu.UpdateBy,
		Remark:      menu.Remark,
		CreatedAt:   util.NewTimestamp(menu.CreatedAt),
		UpdatedAt:   util.NewTimestamp(menu.UpdatedAt),
	}, nil
}
func (s *MenusService) GetMenusTree(ctx context.Context, req *pb.GetMenusTreeRequest) (*pb.GetMenusTreeReply, error) {
	menuList, err := s.menuUseCase.ListByNameStatus(ctx, "", 0)
	if err != nil {
		return nil, err
	}

	return &pb.GetMenusTreeReply{
		List: convertToSimpleMenu(menuList),
	}, nil
}

func convertToSimpleMenu(sysMenus []*model.SysMenu) []*pb.SimpleMenu {
	menuMap := make(map[int64]*pb.SimpleMenu)
	var rootMenus []*pb.SimpleMenu

	// 构建菜单映射
	for _, sysMenu := range sysMenus {
		menu := &pb.SimpleMenu{
			MenuId:   sysMenu.ID,
			MenuName: sysMenu.MenuName,
		}
		menuMap[sysMenu.ID] = menu
	}

	// 构建菜单树
	for _, sysMenu := range sysMenus {
		menu := menuMap[sysMenu.ID]
		if sysMenu.ParentID == 0 {
			rootMenus = append(rootMenus, menu)
		} else {
			parentMenu := menuMap[sysMenu.ParentID]
			parentMenu.Children = append(parentMenu.Children, menu)
		}
	}

	return rootMenus
}

func (s *MenusService) ListMenus(ctx context.Context, req *pb.ListMenusRequest) (*pb.ListMenusReply, error) {
	menuList, err := s.menuUseCase.ListByNameStatus(ctx, req.MenuName, req.Status)
	if err != nil {
		return nil, err
	}
	var reqData []*pb.MenuTree
	if req.MenuName == "" {
		reqData = convertToMenuTree(menuList)
	} else {
		reqData = make([]*pb.MenuTree, len(menuList))
		for i, menu := range menuList {
			reqData[i] = &pb.MenuTree{
				MenuId:     menu.ID,
				MenuName:   menu.MenuName,
				Title:      menu.Title,
				ParentId:   menu.ParentID,
				Sort:       menu.Sort,
				Icon:       menu.Icon,
				Path:       menu.Path,
				Component:  menu.Component,
				IsIframe:   menu.IsIframe,
				Link:       menu.Link,
				MenuType:   menu.MenuType,
				Hidden:     menu.Hidden,
				KeepAlive:  menu.KeepAlive,
				IsAffix:    menu.IsAffix,
				Permission: menu.Permission,
				Status:     menu.Status,
				CreateBy:   menu.CreateBy,
				UpdateBy:   menu.UpdateBy,
				Remark:     menu.Remark,
				CreateTime: util.NewTimestamp(menu.CreatedAt),
				UpdateTime: util.NewTimestamp(menu.UpdatedAt),
			}
		}
	}

	return &pb.ListMenusReply{
		List: reqData,
	}, nil
}

func (s *MenusService) RoleMenuTreeSelect(ctx context.Context, req *pb.RoleMenuTreeSelectRequest) (*pb.RoleMenuTreeSelectReply, error) {
	return s.menuUseCase.RoleMenuTreeSelect(ctx, req)
}

func convertToMenuTree(sysMenus []*model.SysMenu) []*pb.MenuTree {
	menuMap := make(map[int64]*pb.MenuTree)

	// 创建 MenuTree 节点并将其存储在 menuMap 中
	for _, menu := range sysMenus {
		menuTree := &pb.MenuTree{
			MenuId:     menu.ID,
			MenuName:   menu.MenuName,
			Title:      menu.Title,
			ParentId:   menu.ParentID,
			Sort:       menu.Sort,
			Icon:       menu.Icon,
			Path:       menu.Path,
			Component:  menu.Component,
			IsIframe:   menu.IsIframe,
			Link:       menu.Link,
			MenuType:   menu.MenuType,
			Hidden:     menu.Hidden,
			KeepAlive:  menu.KeepAlive,
			IsAffix:    menu.IsAffix,
			Permission: menu.Permission,
			Status:     menu.Status,
			CreateBy:   menu.CreateBy,
			UpdateBy:   menu.UpdateBy,
			Remark:     menu.Remark,
			CreateTime: util.NewTimestamp(menu.CreatedAt),
			UpdateTime: util.NewTimestamp(menu.UpdatedAt),
			Children:   []*pb.MenuTree{},
		}
		menuMap[menu.ID] = menuTree
	}

	// 构建 MenuTree 的层级关系
	var rootMenuTrees []*pb.MenuTree
	for _, menu := range sysMenus {
		if menu.ParentID == 0 {
			rootMenuTrees = append(rootMenuTrees, menuMap[menu.ID])
		} else {
			parentMenuTree, ok := menuMap[menu.ParentID]
			if ok {
				parentMenuTree.Children = append(parentMenuTree.Children, menuMap[menu.ID])
			}
		}
	}

	return rootMenuTrees
}

// Build 构建前端路由
func Build(menus []*pb.MenuTree) []*pb.MenuTreeAuth {
	equals := func(a string, b string) bool {
		if a == b {
			return true
		}
		return false
	}
	rvs := make([]*pb.MenuTreeAuth, 0)
	for _, ms := range menus {
		rv := &pb.MenuTreeAuth{
			Name:      ms.MenuName,
			Path:      ms.Path,
			Component: ms.Component,
			Meta: &pb.MenuTreeMeta{
				Title:       ms.MenuName,
				IsLink:      ms.Link,
				IsHide:      equals("1", string(ms.Hidden)),
				IsKeepAlive: equals("0", string(ms.KeepAlive)),
				IsAffix:     equals("0", string(ms.IsAffix)),
				IsIframe:    equals("0", string(ms.IsIframe)),
				Auth:        make([]string, 0),
				Icon:        ms.Icon,
			},
		}

		if ms.Permission != "" {
			rv.Meta.Auth = strings.Split(ms.Permission, ",")
		}
		rv.Children = Build(ms.Children)
		rvs = append(rvs, rv)
	}
	return rvs
}
