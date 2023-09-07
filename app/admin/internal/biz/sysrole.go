package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/authz"
)

type SysRoleRepo interface {
	Create(ctx context.Context, role *model.SysRole) error
	Save(ctx context.Context, role *model.SysRole) error
	Delete(ctx context.Context, id ...int64) error
	FindByID(ctx context.Context, id int64) (*model.SysRole, error)
	FindByIDList(ctx context.Context, ids ...int64) ([]*model.SysRole, error)
	FindAll(ctx context.Context) ([]*model.SysRole, error)
	ListPage(ctx context.Context, name, key string, status int32, page, size int32) ([]*model.SysRole, error)
	Count(ctx context.Context, name, key string, status int32) (int32, error)
	Update(ctx context.Context, role *model.SysRole) error
}

type SysRoleUseCase struct {
	repo         SysRoleRepo
	menuRepo     SysMenuRepo
	log          *log.Helper
	roleMenuCase *SysRoleMenuUseCase
	casbinCase   *CasbinRuleUseCase
	userUseCase  *SysUserUseCase
	tx           Transaction
}

func NewSysRoleUseCase(repo SysRoleRepo, logger log.Logger, rmc *SysRoleMenuUseCase, casbin *CasbinRuleUseCase, userUseCase *SysUserUseCase, tx Transaction, menuRepo SysMenuRepo) *SysRoleUseCase {
	return &SysRoleUseCase{
		repo:         repo,
		log:          log.NewHelper(logger),
		roleMenuCase: rmc,
		casbinCase:   casbin,
		userUseCase:  userUseCase,
		tx:           tx,
		menuRepo:     menuRepo,
	}
}

func (r *SysRoleUseCase) ListPage(ctx context.Context, roleName, roleKey string, status int32, page, size int32) ([]*model.SysRole, int32, error) {
	total, err := r.repo.Count(ctx, roleName, roleKey, status)
	if err != nil {
		return nil, 0, err
	}
	roleList, err := r.repo.ListPage(ctx, roleName, roleKey, status, page, size)
	return roleList, total, err
}

func (r *SysRoleUseCase) GetRole(ctx context.Context, id int64) (*model.SysRole, error) {
	return r.repo.FindByID(ctx, id)
}

func (r *SysRoleUseCase) CreateRole(ctx context.Context, role *model.SysRole, menuIds []int64, apis []*pb.ApiBase) (*model.SysRole, error) {
	claim := authz.MustFromContext(ctx)
	role.CreateBy = claim.Nickname
	role.CreatedAt = time.Now()

	err := r.tx.Transaction(ctx, func(ctx context.Context) error {
		if err := r.repo.Create(ctx, role); err != nil {
			return err
		}
		// 添加菜单
		if err := r.roleMenuCase.CreateRoleMenus(ctx, role, menuIds); err != nil {
			return err
		}
		// 添加权限
		if err := r.casbinCase.UpdateCasbin(ctx, role.RoleKey, apis); err != nil {
			return err
		}
		return nil
	})

	return role, err
}

func (r *SysRoleUseCase) UpdateRole(ctx context.Context, role *model.SysRole, menuIds []int64, apis []*pb.ApiBase) (*model.SysRole, error) {
	//claims := authz.MustFromContext(ctx)
	oldRole, err := r.repo.FindByID(ctx, role.ID)
	if err != nil {
		return nil, err
	}
	//oldRole.UpdateBy = claims.Nickname
	oldRole.UpdatedAt = time.Now()

	oldRole.ParentID = role.ParentID
	oldRole.RoleName = role.RoleName
	oldRole.Status = role.Status
	oldRole.RoleKey = role.RoleKey
	oldRole.DataScope = role.DataScope
	oldRole.RoleSort = role.RoleSort
	oldRole.DefaultRouter = role.DefaultRouter
	oldRole.Remark = role.Remark

	err = r.tx.Transaction(ctx, func(ctx context.Context) error {
		// 更新Role
		if err = r.repo.Update(ctx, role); err != nil {
			return err
		}
		// 更新菜单
		if err = r.roleMenuCase.DeleteByRoleId(ctx, role.ID); err != nil {
			return err
		}
		if err = r.roleMenuCase.CreateRoleMenus(ctx, role, menuIds); err != nil {
			return err
		}
		// 更新权限
		if err = r.casbinCase.UpdateCasbin(ctx, role.RoleKey, apis); err != nil {
			return err
		}
		return nil
	})
	return role, err
}

func (r *SysRoleUseCase) ChangeRoleStatus(ctx context.Context, id int64, status int32) error {
	claims := authz.MustFromContext(ctx)
	role, err := r.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	role.UpdateBy = claims.Nickname
	role.UpdatedAt = time.Now()
	role.Status = status

	return r.repo.Save(ctx, role)
}

func (r *SysRoleUseCase) ChangeDataScope(ctx context.Context, id int64, scope int32) error {
	claims := authz.MustFromContext(ctx)
	role, err := r.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	role.UpdateBy = claims.Nickname
	role.UpdatedAt = time.Now()
	role.DataScope = scope
	return r.repo.Save(ctx, role)
}

func (r *SysRoleUseCase) DeleteRole(ctx context.Context, ids []int64) error {
	if len(ids) == 0 {
		return nil
	}
	delList := make([]int64, 0)
	roleM := make(map[int64]*model.SysRole)

	for _, rid := range ids {
		//  查询角色下有没有用户
		userCount, err := r.userUseCase.CountByRoleId(ctx, rid)
		if err != nil {
			return err
		}
		if userCount > 0 {
			r.log.Errorf("role: %d 存在用户无法删除", rid)
			continue
		}
		delList = append(delList, rid)

		role, err := r.repo.FindByID(ctx, rid)
		if err != nil {
			return err
		}

		roleM[rid] = role
	}

	err := r.tx.Transaction(ctx, func(ctx context.Context) error {
		// 删除角色
		if err := r.repo.Delete(ctx, delList...); err != nil {
			return err
		}
		// 删除菜单
		if err := r.roleMenuCase.DeleteByRoleId(ctx, delList...); err != nil {
			return err
		}
		// 删除角色绑定api
		for _, roleID := range delList {
			if err := r.casbinCase.ClearCasbin(roleM[roleID].RoleKey); err != nil {
				log.Errorf("删除role:(%d)权限错误; %s", roleID, err.Error())
				continue
			}
		}
		return nil
	})

	return err
}

func (r *SysRoleUseCase) FindRoleByIDList(ctx context.Context, ids []int64) ([]*model.SysRole, error) {
	if len(ids) == 0 {
		return []*model.SysRole{}, nil
	}
	return r.repo.FindByIDList(ctx, ids...)
}

func (r *SysRoleUseCase) FindRoleAll(ctx context.Context) ([]*model.SysRole, error) {
	return r.repo.FindAll(ctx)
}

func (r *SysRoleUseCase) GetRoleMenuId(ctx context.Context, roleId int64) ([]int32, error) {
	return r.menuRepo.GetRoleMenuId(ctx, roleId)
}
