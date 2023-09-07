package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/pkg/util"
)

type PostService struct {
	pb.UnimplementedSysPostServer
	pc  *biz.SysPostUseCase
	log *log.Helper
}

func NewPostService(pc *biz.SysPostUseCase, logger log.Logger) *PostService {
	return &PostService{
		pc:  pc,
		log: log.NewHelper(log.With(logger, "module", "service/post")),
	}
}

func (s *PostService) ListPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostReply, error) {
	postList, total, err := s.pc.ListPost(ctx, req.PostName, req.PostCode, req.Status, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	data := make([]*pb.PostData, len(postList))
	for i, d := range postList {
		data[i] = &pb.PostData{
			PostId:     d.ID,
			PostName:   d.PostName,
			PostCode:   d.PostCode,
			Sort:       d.Sort,
			Status:     d.Status,
			Remark:     d.Remark,
			CreateBy:   d.CreateBy,
			UpdateBy:   d.UpdateBy,
			CreateTime: util.NewTimestamp(d.CreatedAt),
			UpdateTime: util.NewTimestamp(d.UpdatedAt),
		}
	}

	return &pb.ListPostReply{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Total:    total,
		Data:     data,
	}, nil
}
func (s *PostService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := s.pc.CreatePost(ctx, &model.SysPost{
		PostName: req.PostName,
		PostCode: req.PostCode,
		Sort:     req.Sort,
		Status:   req.Status,
		Remark:   req.Remark,
	})
	return &pb.CreatePostReply{}, err
}
func (s *PostService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := s.pc.UpdatePost(ctx, &model.SysPost{
		ID:       req.PostId,
		PostName: req.PostName,
		PostCode: req.PostCode,
		Sort:     req.Sort,
		Status:   req.Status,
		Remark:   req.Remark,
	})
	return &pb.UpdatePostReply{}, err
}
func (s *PostService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostReply, error) {
	ids := util.Split2Int64Slice(req.Id)
	err := s.pc.DeletePost(ctx, ids)
	return &pb.DeletePostReply{}, err
}
