package data

import (
	"context"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type RedisRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.RedisRepo = (*RedisRepo)(nil)

func NewRedisRepo(data *Data, logger log.Logger) biz.RedisRepo {
	return &RedisRepo{
		data: data,
		log: log.NewHelper(log.With(
			logger,
			"module", "admin/data",
		)),
	}
}

func (r *RedisRepo) SetHashKey(ctx context.Context, key string, field string, value interface{}) error {
	return r.data.rdb.HSet(ctx, key, field, value).Err()
}

func (r *RedisRepo) IncrHashKey(ctx context.Context, key string, field string, value int64) error {
	return r.data.rdb.HIncrBy(ctx, key, field, value).Err()
}

func (r *RedisRepo) GetHashKey(ctx context.Context, key string, field string) (string, error) {
	return r.data.rdb.HGet(ctx, key, field).Result()
}

func (r *RedisRepo) DelHashKey(ctx context.Context, key string, field string) error {
	return r.data.rdb.HDel(ctx, key, field).Err()
}

func (r *RedisRepo) GetHashLen(ctx context.Context, key string) error {
	return r.data.rdb.HLen(ctx, key).Err()
}

func (r *RedisRepo) Lock(ctx context.Context, key string, value interface{}, duration time.Duration) (bool, error) {
	res, err := r.data.rdb.SetNX(ctx, key, value, duration).Result()
	return res, err
}

func (r *RedisRepo) GetHashAllKeyAndVal(ctx context.Context, key string) (map[string]string, error) {
	res, err := r.data.rdb.HGetAll(ctx, key).Result()
	mm := make(map[string]string)
	for k, v := range res {
		mm[k] = v
	}
	return mm, err
}

func (r RedisRepo) Set(ctx context.Context, key string, value string, expire time.Duration) error {
	return r.data.rdb.Set(ctx, key, value, expire).Err()
}

func (r RedisRepo) Get(ctx context.Context, key string) string {
	return r.data.rdb.Get(ctx, key).Val()
}

func (r RedisRepo) SRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return r.data.rdb.SRem(ctx, key, members).Result()
}
