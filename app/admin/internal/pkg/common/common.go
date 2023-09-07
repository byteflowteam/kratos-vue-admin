package common

import (
	"context"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/authz"
)

// GetAdminNickname 获取当前登录管理员昵称
func GetAdminNickname(ctx context.Context) string {
	claims := authz.MustFromContext(ctx)
	return claims.Nickname
}
