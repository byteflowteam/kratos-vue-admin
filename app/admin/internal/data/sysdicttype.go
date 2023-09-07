package data

import (
	"context"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/go-kratos/kratos/v2/log"
)

type sysDictTypeRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysDictTypeRepo(data *Data, logger log.Logger) biz.SysDictTypeRepo {
	return &sysDictTypeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p *sysDictTypeRepo) Create(ctx context.Context, post *model.SysDictType) error {
	q := p.data.Query(ctx).SysDictType
	return q.WithContext(ctx).Create(post)
}

func (p *sysDictTypeRepo) Save(ctx context.Context, post *model.SysDictType) error {
	q := p.data.Query(ctx).SysDictType
	return q.WithContext(ctx).Save(post)
}

func (p *sysDictTypeRepo) Delete(ctx context.Context, ids []int64) error {
	q := p.data.Query(ctx).SysDictType
	_, err := q.WithContext(ctx).Where(q.DictID.In(ids...)).Delete()
	return err
}

func (p *sysDictTypeRepo) FindByID(ctx context.Context, id int64) (*model.SysDictType, error) {
	q := p.data.Query(ctx).SysDictType
	return q.WithContext(ctx).Where(q.DictID.Eq(id)).First()
}

func (p *sysDictTypeRepo) ListPage(ctx context.Context, dictName, dictType string, status int32, page, size int32) ([]*model.SysDictType, error) {
	q := p.data.Query(ctx).SysDictType
	db := q.WithContext(ctx)
	if dictName != "" {
		db = db.Where(q.DictName.Like(buildLikeValue(dictName)))
	}
	if dictType != "" {
		db = db.Where(q.DictType.Like(buildLikeValue(dictType)))
	}
	if status > 0 {
		db = db.Where(q.Status.Eq(status))
	}
	limit, offset := convertPageSize(page, size)
	return db.Limit(limit).Offset(offset).Find()
}

func (p *sysDictTypeRepo) ListPageCount(ctx context.Context, dictName, dictType string, status int32) (int32, error) {
	q := p.data.Query(ctx).SysDictType
	db := q.WithContext(ctx)
	if dictName != "" {
		db = db.Where(q.DictName.Like(buildLikeValue(dictName)))
	}
	if dictType != "" {
		db = db.Where(q.DictType.Like(buildLikeValue(dictType)))
	}
	if status > 0 {
		db = db.Where(q.Status.Eq(status))
	}
	count, err := db.Count()
	return int32(count), err
}

func (p *sysDictTypeRepo) FindByIDList(ctx context.Context, ids ...int64) ([]*model.SysDictType, error) {
	q := p.data.Query(ctx).SysDictType
	return q.WithContext(ctx).Where(q.DictID.In(ids...)).Find()
}

func (p *sysDictTypeRepo) FindAll(ctx context.Context) ([]*model.SysDictType, error) {
	q := p.data.Query(ctx).SysDictType
	return q.WithContext(ctx).Find()
}
