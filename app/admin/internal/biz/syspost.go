package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/authz"
)

type SysPostRepo interface {
	Create(ctx context.Context, post *model.SysPost) error
	Save(ctx context.Context, post *model.SysPost) error
	Delete(ctx context.Context, ids []int64) error
	FindByID(ctx context.Context, id int64) (*model.SysPost, error)
	FindByIDList(ctx context.Context, ids ...int64) ([]*model.SysPost, error)
	FindAll(ctx context.Context) ([]*model.SysPost, error)

	ListPage(ctx context.Context, postName, postCode string, status int32, page, size int32) ([]*model.SysPost, error)
	ListPageCount(ctx context.Context, postName, postCode string, status int32) (int32, error)
}

type SysPostUseCase struct {
	repo SysPostRepo
	log  *log.Helper
	uc   *SysUserUseCase
}

func NewSysPostUseCase(repo SysPostRepo, logger log.Logger, uc *SysUserUseCase) *SysPostUseCase {
	return &SysPostUseCase{repo: repo, log: log.NewHelper(logger), uc: uc}
}

func (p *SysPostUseCase) ListPost(ctx context.Context, postName, postCode string, status int32, page, size int32) ([]*model.SysPost, int32, error) {
	total, err := p.repo.ListPageCount(ctx, postName, postCode, status)
	if err != nil {
		return nil, 0, err
	}
	posts, err := p.repo.ListPage(ctx, postName, postCode, status, page, size)
	return posts, total, err
}

func (p *SysPostUseCase) CreatePost(ctx context.Context, post *model.SysPost) (*model.SysPost, error) {
	claims := authz.MustFromContext(ctx)
	post.CreateBy = claims.Nickname

	err := p.repo.Create(ctx, post)
	return post, err
}

func (p *SysPostUseCase) UpdatePost(ctx context.Context, post *model.SysPost) (*model.SysPost, error) {
	claims := authz.MustFromContext(ctx)
	post.UpdateBy = claims.Nickname
	err := p.repo.Save(ctx, post)
	return post, err
}

func (p *SysPostUseCase) DeletePost(ctx context.Context, ids []int64) error {
	deList := make([]int64, 0)
	for _, postId := range ids {
		posts, err := p.uc.FindByPostId(ctx, postId)
		if err != nil {
			return err
		}
		if len(posts) == 0 {
			deList = append(deList, postId)
		} else {
			return errors.New("岗位已绑定用户, 无法删除")
		}
	}
	if len(deList) == 0 {
		return nil
	}

	return p.repo.Delete(ctx, deList)
}

func (p *SysPostUseCase) FindPostByIDList(ctx context.Context, ids []int64) ([]*model.SysPost, error) {
	if len(ids) == 0 {
		return []*model.SysPost{}, nil
	}
	return p.repo.FindByIDList(ctx, ids...)
}

func (p *SysPostUseCase) FindPostByID(ctx context.Context, id int64) (*model.SysPost, error) {
	return p.repo.FindByID(ctx, id)
}

func (p *SysPostUseCase) FindPostAll(ctx context.Context) ([]*model.SysPost, error) {
	return p.repo.FindAll(ctx)
}
