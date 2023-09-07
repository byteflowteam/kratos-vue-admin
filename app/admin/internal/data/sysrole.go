package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
)

type sysRoleRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysRoleRepo(data *Data, logger log.Logger) biz.SysRoleRepo {
	return &sysRoleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *sysRoleRepo) Create(ctx context.Context, role *model.SysRole) error {
	q := r.data.Query(ctx).SysRole
	return q.WithContext(ctx).Create(role)
}

func (r *sysRoleRepo) Save(ctx context.Context, role *model.SysRole) error {
	q := r.data.Query(ctx).SysRole
	return q.WithContext(ctx).Save(role)
}

func (r *sysRoleRepo) Delete(ctx context.Context, ids ...int64) error {
	q := r.data.Query(ctx).SysRole
	_, err := q.WithContext(ctx).Where(q.ID.In(ids...)).Delete()
	return err
}

func (r *sysRoleRepo) Update(ctx context.Context, role *model.SysRole) error {
	q := r.data.Query(ctx).SysRole
	_, err := q.WithContext(ctx).Select(q.UpdatedAt, q.Sort, q.DefaultRouter, q.RoleName, q.RoleKey, q.Status, q.DataScope, q.Remark).Where(q.ID.Eq(role.ID)).Updates(role)
	return err
}

func (r *sysRoleRepo) FindByID(ctx context.Context, id int64) (*model.SysRole, error) {
	q := r.data.Query(ctx).SysRole
	return q.WithContext(ctx).Where(q.ID.Eq(id)).First()

}

func (r *sysRoleRepo) ListPage(ctx context.Context, name, key string, status int32, page, size int32) ([]*model.SysRole, error) {
	q := r.data.Query(ctx).SysRole
	db := q.WithContext(ctx)
	if name != "" {
		db = db.Where(q.RoleName.Like(buildLikeValue(name)))
	}
	if key != "" {
		db = db.Where(q.RoleKey.Eq(key))
	}
	if status != 0 {
		db = db.Where(q.Status.Eq(status))
	}
	limit, offset := convertPageSize(page, size)
	return db.Limit(limit).Offset(offset).Find()
}

func (r *sysRoleRepo) Count(ctx context.Context, name, key string, status int32) (int32, error) {
	q := r.data.Query(ctx).SysRole
	db := q.WithContext(ctx)
	if name != "" {
		db = db.Where(q.RoleName.Like(buildLikeValue(name)))
	}
	if key != "" {
		db = db.Where(q.RoleKey.Eq(key))
	}
	if status != 0 {
		db = db.Where(q.Status.Eq(status))
	}
	counts, err := db.Count()
	return int32(counts), err
}

func (r *sysRoleRepo) FindByIDList(ctx context.Context, ids ...int64) ([]*model.SysRole, error) {
	q := r.data.Query(ctx).SysRole
	return q.WithContext(ctx).Where(q.ID.In(ids...)).Find()
}

func (r *sysRoleRepo) FindAll(ctx context.Context) ([]*model.SysRole, error) {
	q := r.data.Query(ctx).SysRole
	return q.WithContext(ctx).Find()
}
