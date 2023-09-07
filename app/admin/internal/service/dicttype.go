package service

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/pkg/util"
)

type DictTypeService struct {
	pb.UnimplementedDictTypeServer
	pc  *biz.SysDictTypeUseCase
	log *log.Helper
}

func NewDictTypeService(pc *biz.SysDictTypeUseCase, logger log.Logger) *DictTypeService {
	return &DictTypeService{
		pc:  pc,
		log: log.NewHelper(log.With(logger, "module", "service/dict_type")),
	}
}

func (s *DictTypeService) ListDictType(ctx context.Context, req *pb.ListDictTypeRequest) (*pb.ListDictTypeReply, error) {
	List, total, err := s.pc.ListDictType(ctx, req.DictType, req.DictName, req.Status, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	data := make([]*pb.DictTypeContent, len(List))
	for i, d := range List {
		data[i] = &pb.DictTypeContent{
			DictId:     uint32(d.DictID),
			DictName:   d.DictName,
			DictType:   d.DictType,
			Status:     d.Status,
			Remark:     d.Remark,
			CreateBy:   d.CreateBy,
			UpdateBy:   d.UpdateBy,
			CreateTime: util.NewTimestamp(d.CreateTime),
			UpdateTime: util.NewTimestamp(d.UpdateTime),
		}
	}

	return &pb.ListDictTypeReply{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Total:    total,
		Data:     data,
	}, nil
}
func (s *DictTypeService) CreateDictType(ctx context.Context, req *pb.CreateDictTypeRequest) (*pb.CreateDictTypeReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := s.pc.CreateDictType(ctx, &model.SysDictType{
		DictType: req.DictType,
		DictName: req.DictName,
		Status:   req.Status,
		Remark:   req.Remark,
	})
	return &pb.CreateDictTypeReply{}, err
}
func (s *DictTypeService) UpdateDictType(ctx context.Context, req *pb.UpdateDictTypeRequest) (*pb.UpdateDictTypeReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := s.pc.UpdateDictType(ctx, &model.SysDictType{
		DictID:   int64(req.DictId),
		DictName: req.DictName,
		DictType: req.DictType,
		Status:   req.Status,
		Remark:   req.Remark,
		CreateBy: req.CreateBy,
		UpdateBy: req.UpdateBy,
		//CreateTime: util.NewTimestamp(req.CreateTime),
		//UpdateTime: util.NewTimestamp(req.UpdateTime),
	})
	return &pb.UpdateDictTypeReply{}, err
}
func (s *DictTypeService) DeleteDictType(ctx context.Context, req *pb.DeleteDictTypeRequest) (*pb.DeleteDictTypeReply, error) {
	ids := util.Split2Int64Slice(strconv.FormatInt(req.DictId, 10))
	err := s.pc.DeleteDictType(ctx, ids)
	return &pb.DeleteDictTypeReply{}, err
}
