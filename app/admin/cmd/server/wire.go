//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/conf"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/oss"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/server"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Auth, *conf.Casbin, *conf.Oss, log.Logger, *conf.Data_Redis) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, oss.ProviderSet, newApp))
}
