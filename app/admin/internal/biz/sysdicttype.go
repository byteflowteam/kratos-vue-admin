package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/authz"
)

type SysDictTypeRepo interface {
	Create(ctx context.Context, post *model.SysDictType) error
	Save(ctx context.Context, post *model.SysDictType) error
	Delete(ctx context.Context, ids []int64) error
	FindByID(ctx context.Context, id int64) (*model.SysDictType, error)
	FindByIDList(ctx context.Context, ids ...int64) ([]*model.SysDictType, error)
	FindAll(ctx context.Context) ([]*model.SysDictType, error)

	ListPage(ctx context.Context, dictName, dictType string, status int32, page, size int32) ([]*model.SysDictType, error)
	ListPageCount(ctx context.Context, dictName, dictType string, status int32) (int32, error)
}

type SysDictTypeUseCase struct {
	repo SysDictTypeRepo
	log  *log.Helper
}

func NewSysDictTypeUseCase(repo SysDictTypeRepo, logger log.Logger) *SysDictTypeUseCase {
	return &SysDictTypeUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (p *SysDictTypeUseCase) ListDictType(ctx context.Context, dictName, dictType string, status int32, page, size int32) ([]*model.SysDictType, int32, error) {
	total, err := p.repo.ListPageCount(ctx, dictName, dictType, status)
	if err != nil {
		return nil, 0, err
	}
	dictLists, err := p.repo.ListPage(ctx, dictName, dictType, status, page, size)
	return dictLists, total, err
}

func (p *SysDictTypeUseCase) CreateDictType(ctx context.Context, post *model.SysDictType) (*model.SysDictType, error) {
	claims := authz.MustFromContext(ctx)
	post.CreateBy = claims.Nickname

	err := p.repo.Create(ctx, post)
	return post, err
}

func (p *SysDictTypeUseCase) UpdateDictType(ctx context.Context, post *model.SysDictType) (*model.SysDictType, error) {
	claims := authz.MustFromContext(ctx)
	post.UpdateBy = claims.Nickname
	err := p.repo.Save(ctx, post)
	return post, err
}

func (p *SysDictTypeUseCase) DeleteDictType(ctx context.Context, id []int64) error {
	return p.repo.Delete(ctx, id)
}

func (p *SysDictTypeUseCase) FindDictTypeByIDList(ctx context.Context, ids []int64) ([]*model.SysDictType, error) {
	if len(ids) == 0 {
		return []*model.SysDictType{}, nil
	}
	return p.repo.FindByIDList(ctx, ids...)
}

func (p *SysDictTypeUseCase) FindDictTypeByID(ctx context.Context, id int64) (*model.SysDictType, error) {
	return p.repo.FindByID(ctx, id)
}

func (p *SysDictTypeUseCase) FindDictTypeAll(ctx context.Context) ([]*model.SysDictType, error) {
	return p.repo.FindAll(ctx)
}
