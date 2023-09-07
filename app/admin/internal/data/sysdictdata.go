package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
)

type sysDictDataRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysDictDataRepo(data *Data, logger log.Logger) biz.SysDictDatumRepo {
	return &sysDictDataRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p *sysDictDataRepo) Create(ctx context.Context, post *model.SysDictDatum) error {
	q := p.data.Query(ctx).SysDictDatum
	return q.WithContext(ctx).Create(post)
}

func (p *sysDictDataRepo) Save(ctx context.Context, post *model.SysDictDatum) error {
	q := p.data.Query(ctx).SysDictDatum
	return q.WithContext(ctx).Save(post)
}

func (p *sysDictDataRepo) Delete(ctx context.Context, ids []int64) error {
	q := p.data.Query(ctx).SysDictDatum
	_, err := q.WithContext(ctx).Where(q.DictCode.In(ids...)).Delete()
	return err
}

func (p *sysDictDataRepo) FindByID(ctx context.Context, id int64) (*model.SysDictDatum, error) {
	q := p.data.Query(ctx).SysDictDatum
	return q.WithContext(ctx).Where(q.DictCode.Eq(id)).First()
}

func (p *sysDictDataRepo) ListPage(ctx context.Context, dictLabel, dictType string, status int32, page, size int32) ([]*model.SysDictDatum, error) {
	q := p.data.Query(ctx).SysDictDatum
	db := q.WithContext(ctx)
	if dictLabel != "" {
		db = db.Where(q.DictLabel.Like(buildLikeValue(dictLabel)))
	}
	if dictType != "" {
		db = db.Where(q.DictType.Eq(dictType))
	}
	if status > 0 {
		db = db.Where(q.Status.Eq(status))
	}
	limit, offset := convertPageSize(page, size)
	return db.Limit(limit).Offset(offset).Find()
}

func (p *sysDictDataRepo) ListPageCount(ctx context.Context, dictLabel, dictType string, status int32) (int32, error) {
	q := p.data.Query(ctx).SysDictDatum
	db := q.WithContext(ctx)
	if dictLabel != "" {
		db = db.Where(q.DictLabel.Like(buildLikeValue(dictLabel)))
	}
	if dictType != "" {
		db = db.Where(q.DictType.Eq(dictType))
	}
	if status > 0 {
		db = db.Where(q.Status.Eq(status))
	}
	count, err := db.Count()
	return int32(count), err
}

func (p *sysDictDataRepo) FindByIDList(ctx context.Context, ids ...int64) ([]*model.SysDictDatum, error) {
	q := p.data.Query(ctx).SysDictDatum
	return q.WithContext(ctx).Where(q.DictCode.In(ids...)).Find()
}

func (p *sysDictDataRepo) FindAll(ctx context.Context) ([]*model.SysDictDatum, error) {
	q := p.data.Query(ctx).SysDictDatum
	return q.WithContext(ctx).Find()
}
