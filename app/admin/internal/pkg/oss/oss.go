package oss

import (
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/conf"
)

var ProviderSet = wire.NewSet(
	NewOssRepo,
)

func NewOssRepo(c *conf.Oss, logger log.Logger) biz.OssRepo {
	logs := log.NewHelper(log.With(logger, "module", "app/admin/internal/pkg/oss"))
	var err error
	var repo biz.OssRepo
	switch c.Use {
	case conf.OssUseMode_aliyun:
		repo, err = newAliyunClient(logs, c.Aliyun)
	case conf.OssUseMode_local:
		repo, err = newLocalClient(logs, c.Local)
	default:
		err = errors.New("invalid oss use mod")
	}
	if err != nil {
		panic(err)
	}
	return repo
}
