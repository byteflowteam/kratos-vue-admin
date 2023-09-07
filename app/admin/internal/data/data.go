package data

import (
	"context"
	"fmt"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/conf"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/query"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	go_redis "github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDB,
	NewSysUserRepo,
	NewSysMenuRepo,
	NewSysDeptRepo,
	NewSysPostRepo,
	NewSysApiRepo,
	NewSysRoleRepo,
	NewSysRoleMenuRepo,
	NewTransaction,
	NewCasbinRuleRepo,
	NewRedis,
	NewSysDictDataRepo,
	NewSysDictTypeRepo,
	NewRedisRepo,
)

// Data .
type Data struct {
	log   *log.Helper
	query *query.Query
	db    *gorm.DB
	rdb   go_redis.UniversalClient
}

type contextTxKey struct{}

func toGormLogLevel(d conf.GormLogLevel) gormLogger.LogLevel {
	switch d {
	case conf.GormLogLevel_silent:
		return gormLogger.Silent
	case conf.GormLogLevel_error:
		return gormLogger.Error
	case conf.GormLogLevel_warn:
		return gormLogger.Warn
	case conf.GormLogLevel_info:
		return gormLogger.Info
	default:
		return gormLogger.Warn
	}
}

func NewDB(config *conf.Data, logger log.Logger) *gorm.DB {
	logs := log.NewHelper(log.With(logger, "module", "receive-service/data/gorm"))

	db, err := gorm.Open(mysql.Open(config.Database.Source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   gormLogger.Default.LogMode(toGormLogLevel(config.Database.LogLevel)),
	})
	if err != nil {
		logs.Fatalf("failed opening connection to mysql: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logs.Fatalf("failed get sql.DB : %v", err)
	}
	sqlDB.SetMaxIdleConns(int(config.Database.MaxIdleConns))
	sqlDB.SetMaxOpenConns(int(config.Database.MaxOpenConns))

	return db
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger, rdb go_redis.UniversalClient) (*Data, func(), error) {
	logs := log.NewHelper(log.With(logger, "module", "receive-service/data"))
	d := &Data{
		log:   logs,
		query: query.Use(db),
		db:    db,
		rdb:   rdb,
	}
	return d, func() {}, nil
}

func NewTransaction(d *Data) biz.Transaction {
	return d
}

func (d *Data) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.query.Transaction(func(tx *query.Query) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) Query(ctx context.Context) *query.Query {
	tx, ok := ctx.Value(contextTxKey{}).(*query.Query)
	if ok {
		return tx
	}
	return d.query
}

func convertPageSize(page, size int32) (limit, offset int) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	limit = int(size)
	offset = int((page - 1) * size)
	return
}

func buildLikeValue(key string) string {
	return fmt.Sprintf("%%%s%%", key)
}

func NewRedis(conf *conf.Data) go_redis.UniversalClient {
	var redisClient go_redis.UniversalClient
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	redisClient = go_redis.NewClient(&go_redis.Options{
		Addr:     conf.Redis.Addr,
		Username: conf.Redis.Username,
		Password: conf.Redis.Password,      // no password set
		DB:       int(conf.Redis.Database), // use default DB
		PoolSize: 100,                      // 连接池大小
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return redisClient
}
