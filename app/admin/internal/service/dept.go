package service

import (
	"context"
	"fmt"

	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

type DeptService struct {
	pb.UnimplementedDeptServer
	deptUseCase *biz.SysDeptUseCase
	log         *log.Helper
}

func NewDeptService(bc *biz.SysDeptUseCase, logger log.Logger) *DeptService {
	return &DeptService{
		deptUseCase: bc,
		log:         log.NewHelper(log.With(logger, "module", "service/dept")),
	}
}

func (s *DeptService) ListDept(ctx context.Context, req *pb.ListDeptRequest) (*pb.ListDeptReply, error) {
	deptList, err := s.deptUseCase.ListByNameStatusId(ctx, req.DeptName, req.Status, req.DeptId)
	if err != nil {
		return nil, err
	}
	var data []*pb.DeptTree
	if req.DeptName == "" {
		data = biz.ConvertToDeptTree(deptList)
	} else {
		data = make([]*pb.DeptTree, len(deptList))
		for i, d := range deptList {
			data[i] = &pb.DeptTree{
				DeptId:     d.ID,
				ParentId:   d.ParentID,
				DeptPath:   d.DeptPath,
				DeptName:   d.DeptName,
				Sort:       d.Sort,
				Leader:     d.Leader,
				Phone:      d.Phone,
				Email:      d.Email,
				Status:     d.Status,
				CreateBy:   d.CreateBy,
				UpdateBy:   d.UpdateBy,
				CreateTime: util.NewTimestamp(d.CreatedAt),
				UpdateTime: util.NewTimestamp(d.UpdatedAt),
			}
		}
	}

	return &pb.ListDeptReply{
		Data: data,
	}, nil
}
func (s *DeptService) GetDeptTree(ctx context.Context, req *pb.GetDeptTreeRequest) (*pb.GetDeptTreeReply, error) {
	deptList, err := s.deptUseCase.ListByNameStatusId(ctx, req.DeptName, req.Status, req.DeptId)
	if err != nil {
		return nil, err
	}
	data := biz.ConvertToDeptTree(deptList)
	return &pb.GetDeptTreeReply{
		Data: data,
	}, nil
}
func (s *DeptService) CreateDept(ctx context.Context, req *pb.CreateDeptRequest) (*pb.CreateDeptReply, error) {
	_, err := s.deptUseCase.CreateDept(ctx, &model.SysDept{
		ParentID: req.ParentId,
		DeptName: req.DeptName,
		Sort:     req.Sort,
		Leader:   req.Leader,
		Phone:    req.Phone,
		Email:    req.Email,
		Status:   req.Status,
	})
	return &pb.CreateDeptReply{}, err
}
func (s *DeptService) UpdateDept(ctx context.Context, req *pb.UpdateDeptRequest) (*pb.UpdateDeptReply, error) {
	fmt.Print("status:", req.Status)
	_, err := s.deptUseCase.UpdateDept(ctx, &model.SysDept{
		ID:       req.DeptId,
		ParentID: req.ParentId,
		DeptName: req.DeptName,
		Sort:     req.Sort,
		Leader:   req.Leader,
		Phone:    req.Phone,
		Email:    req.Email,
		Status:   int32(req.Status),
	})
	return &pb.UpdateDeptReply{}, err
}
func (s *DeptService) DeleteDept(ctx context.Context, req *pb.DeleteDeptRequest) (*pb.DeleteDeptReply, error) {
	err := s.deptUseCase.DeleteDept(ctx, req.Id)
	return &pb.DeleteDeptReply{}, err
}

func (s *DeptService) RoleDeptTreeSelect(ctx context.Context, req *pb.RoleDeptTreeSelectRequest) (*pb.RoleDeptTreeSelectReply, error) {
	return s.deptUseCase.RoleDeptTreeSelect(ctx, req.RoleId)
}
