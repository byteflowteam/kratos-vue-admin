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

type DictDataService struct {
	pb.UnimplementedDictDataServer
	pc  *biz.SysDictDatumUseCase
	log *log.Helper
}

func NewDictDataService(pc *biz.SysDictDatumUseCase, logger log.Logger) *DictDataService {
	return &DictDataService{
		pc:  pc,
		log: log.NewHelper(log.With(logger, "module", "service/dict_data")),
	}
}

func (s *DictDataService) ListDictData(ctx context.Context, req *pb.ListDictDataRequest) (*pb.ListDictDataReply, error) {
	postList, total, err := s.pc.ListDictData(ctx, req.DictLabel, req.DictType, req.Status, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	data := make([]*pb.DictDataContent, len(postList))
	for i, d := range postList {
		data[i] = &pb.DictDataContent{
			DictCode:   uint32(d.DictCode),
			DictType:   d.DictType,
			DictSort:   uint32(d.DictSort),
			DictLabel:  d.DictLabel,
			DictValue:  d.DictValue,
			Status:     d.Status,
			CssClass:   d.CSSClass,
			ListClass:  d.ListClass,
			IsDefault:  d.IsDefault,
			Remark:     d.Remark,
			CreateBy:   d.CreateBy,
			UpdateBy:   d.UpdateBy,
			CreateTime: util.NewTimestamp(d.CreateTime),
			UpdateTime: util.NewTimestamp(d.UpdateTime),
		}
	}

	return &pb.ListDictDataReply{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Total:    total,
		List:     data,
	}, nil
}
func (s *DictDataService) CreateDictData(ctx context.Context, req *pb.CreateDictDataRequest) (*pb.CreateDictDataReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := s.pc.CreateDictData(ctx, &model.SysDictDatum{
		DictType:  req.DictType,
		DictLabel: req.DictLabel,
		DictValue: req.DictValue,
		DictSort:  int32(req.DictSort),
		Status:    req.Status,
		Remark:    req.Remark,
	})
	return &pb.CreateDictDataReply{}, err
}
func (s *DictDataService) UpdateDictData(ctx context.Context, req *pb.UpdateDictDataRequest) (*pb.UpdateDictDataReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := s.pc.UpdateDictData(ctx, &model.SysDictDatum{
		DictCode:  req.DictCode,
		DictType:  req.DictType,
		DictLabel: req.DictLabel,
		DictValue: req.DictValue,
		DictSort:  req.DictSort,
		Status:    req.Status,
		Remark:    req.Remark,
	})
	return &pb.UpdateDictDataReply{}, err
}
func (s *DictDataService) DeleteDictData(ctx context.Context, req *pb.DeleteDictDataRequest) (*pb.DeleteDictDataReply, error) {
	ids := util.Split2Int64Slice(strconv.FormatInt(req.DictCode, 10))
	err := s.pc.DeleteDictData(ctx, ids)
	return &pb.DeleteDictDataReply{}, err
}
