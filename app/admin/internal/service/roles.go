package service

import (
	"context"

	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

type RolesService struct {
	pb.UnimplementedRolesServer
	rc     *biz.SysRoleUseCase
	casbin *biz.CasbinRuleUseCase
	log    *log.Helper
}

func NewRolesService(rc *biz.SysRoleUseCase, logger log.Logger, casbin *biz.CasbinRuleUseCase) *RolesService {
	return &RolesService{
		rc:     rc,
		casbin: casbin,
		log:    log.NewHelper(log.With(logger, "module", "service/roles")),
	}
}

func (r *RolesService) ListRoles(ctx context.Context, req *pb.ListRolesRequest) (*pb.ListRolesReply, error) {
	roleList, total, err := r.rc.ListPage(ctx, req.RoleName, req.RoleKey, req.Status, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	data := make([]*pb.RoleData, len(roleList))
	for i, d := range roleList {
		data[i] = &pb.RoleData{
			RoleId:     d.ID,
			RoleName:   d.RoleName,
			Status:     d.Status,
			RoleKey:    d.RoleKey,
			RoleSort:   d.RoleSort,
			DataScope:  int64(d.DataScope),
			CreateBy:   d.CreateBy,
			UpdateBy:   d.UpdateBy,
			Remark:     d.Remark,
			CreateTime: util.NewTimestamp(d.CreatedAt),
			UpdateTime: util.NewTimestamp(d.UpdatedAt),
		}

	}
	return &pb.ListRolesReply{
		Total:    total,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Data:     data,
	}, nil
}

func (r *RolesService) GetRoles(ctx context.Context, req *pb.GetRolesRequest) (*pb.GetRolesReply, error) {
	role, err := r.rc.GetRole(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	menuIds, err := r.rc.GetRoleMenuId(ctx, role.ID)
	if err != nil {
		return nil, err
	}
	return &pb.GetRolesReply{
		Role: &pb.RoleData{
			RoleId:     role.ID,
			RoleName:   role.RoleName,
			Status:     role.Status,
			RoleKey:    role.RoleKey,
			RoleSort:   role.RoleSort,
			DataScope:  int64(role.DataScope),
			MenuIds:    menuIds,
			CreateBy:   role.CreateBy,
			UpdateBy:   role.UpdateBy,
			Remark:     role.Remark,
			CreateTime: util.NewTimestamp(role.CreatedAt),
			UpdateTime: util.NewTimestamp(role.UpdatedAt),
		},
	}, nil
}

func (r *RolesService) CreateRoles(ctx context.Context, req *pb.CreateRolesRequest) (*pb.CreateRolesReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := r.rc.CreateRole(ctx, &model.SysRole{
		ParentID:      req.ParentId,
		RoleName:      req.RoleName,
		Status:        req.Status,
		RoleKey:       req.RoleKey,
		DataScope:     req.DataScope,
		RoleSort:      req.Sort,
		DefaultRouter: req.DefaultRouter,
		Remark:        req.Remark,
	}, req.MenuIds, req.ApiIds)

	return &pb.CreateRolesReply{}, err
}

func (r *RolesService) UpdateRoles(ctx context.Context, req *pb.UpdateRolesRequest) (*pb.UpdateRolesReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	_, err := r.rc.UpdateRole(ctx, &model.SysRole{
		ID:            req.RoleId,
		ParentID:      req.ParentId,
		RoleName:      req.RoleName,
		Status:        req.Status,
		RoleKey:       req.RoleKey,
		DataScope:     req.DataScope,
		RoleSort:      req.Sort,
		DefaultRouter: req.DefaultRouter,
		Remark:        req.Remark,
	}, req.MenuIds, req.ApiIds)
	return &pb.UpdateRolesReply{}, err
}

func (r *RolesService) ChangeRoleStatus(ctx context.Context, req *pb.ChangeRoleStatusRequest) (*pb.ChangeRoleStatusReply, error) {
	err := r.rc.ChangeRoleStatus(ctx, req.RoleId, req.Status)
	return &pb.ChangeRoleStatusReply{}, err
}

func (r *RolesService) DataScope(ctx context.Context, req *pb.DataScopeRequest) (*pb.DataScopeReply, error) {
	err := r.rc.ChangeDataScope(ctx, req.RoleId, req.DataScope)
	return &pb.DataScopeReply{}, err
}

func (r *RolesService) DeleteRoles(ctx context.Context, req *pb.DeleteRolesRequest) (*pb.DeleteRolesReply, error) {
	ids := util.Split2Int64Slice(req.Id)
	err := r.rc.DeleteRole(ctx, ids)
	return &pb.DeleteRolesReply{}, err
}
