package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/authz"
)

type SysDictDatumRepo interface {
	Create(ctx context.Context, post *model.SysDictDatum) error
	Save(ctx context.Context, post *model.SysDictDatum) error
	Delete(ctx context.Context, ids []int64) error
	FindByID(ctx context.Context, id int64) (*model.SysDictDatum, error)
	FindByIDList(ctx context.Context, ids ...int64) ([]*model.SysDictDatum, error)
	FindAll(ctx context.Context) ([]*model.SysDictDatum, error)

	ListPage(ctx context.Context, dictLabel, dictType string, status int32, page, size int32) ([]*model.SysDictDatum, error)
	ListPageCount(ctx context.Context, dictLabel, dictType string, status int32) (int32, error)
}

type SysDictDatumUseCase struct {
	repo SysDictDatumRepo
	log  *log.Helper
}

func NewSysDictDatumUseCase(repo SysDictDatumRepo, logger log.Logger) *SysDictDatumUseCase {
	return &SysDictDatumUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (p *SysDictDatumUseCase) ListDictData(ctx context.Context, dictLabel, dictType string, status int32, page, size int32) ([]*model.SysDictDatum, int32, error) {
	total, err := p.repo.ListPageCount(ctx, dictLabel, dictType, status)
	if err != nil {
		return nil, 0, err
	}
	posts, err := p.repo.ListPage(ctx, dictLabel, dictType, status, page, size)
	return posts, total, err
}

func (p *SysDictDatumUseCase) CreateDictData(ctx context.Context, post *model.SysDictDatum) (*model.SysDictDatum, error) {
	claims := authz.MustFromContext(ctx)
	post.CreateBy = claims.Nickname

	err := p.repo.Create(ctx, post)
	return post, err
}

func (p *SysDictDatumUseCase) UpdateDictData(ctx context.Context, post *model.SysDictDatum) (*model.SysDictDatum, error) {
	claims := authz.MustFromContext(ctx)
	post.UpdateBy = claims.Nickname
	err := p.repo.Save(ctx, post)
	return post, err
}

func (p *SysDictDatumUseCase) DeleteDictData(ctx context.Context, id []int64) error {
	return p.repo.Delete(ctx, id)
}

func (p *SysDictDatumUseCase) FindDictDataByIDList(ctx context.Context, ids []int64) ([]*model.SysDictDatum, error) {
	if len(ids) == 0 {
		return []*model.SysDictDatum{}, nil
	}
	return p.repo.FindByIDList(ctx, ids...)
}

func (p *SysDictDatumUseCase) FindDictDataByID(ctx context.Context, id int64) (*model.SysDictDatum, error) {
	return p.repo.FindByID(ctx, id)
}

func (p *SysDictDatumUseCase) FindDictDataAll(ctx context.Context) ([]*model.SysDictDatum, error) {
	return p.repo.FindAll(ctx)
}
