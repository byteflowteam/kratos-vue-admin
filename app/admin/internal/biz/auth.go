package biz

import (
	"context"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/conf"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/authz"
	"github.com/byteflowteam/kratos-vue-admin/pkg/util"
)

type AuthUseCase struct {
	key      string
	expire   time.Duration
	userRepo SysUserRepo
	roleRepo SysRoleRepo
	log      *log.Helper
}

func NewAuthUseCase(conf *conf.Auth, userRepo SysUserRepo, roleRepo SysRoleRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{
		key:      conf.JwtKey,
		expire:   conf.Expires.AsDuration(),
		userRepo: userRepo,
		roleRepo: roleRepo,
		log:      log.NewHelper(logger),
	}
}

func (receiver *AuthUseCase) Login(ctx context.Context, req *pb.LoginRequest) (token string, expireAt int64, pErr error) {
	// get user
	user, err := receiver.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		pErr = pb.ErrorUserNotFound("用户名或密码错误")
		return
	}

	gAuth := util.NewGoogleAuth()
	code, err := gAuth.GetCode(user.Secret)

	if err != nil {
		pErr = pb.ErrorInternalErr(err.Error())
		return
	}

	if req.Code != code {
		pErr = pb.ErrorCodeNotMatch(pkg.ErrGoogleCode)
		return
	}

	if !util.BcryptCheck(req.Password, user.Password) {
		pErr = pb.ErrorLoginFail(pkg.ErrPassword)
		return
	}

	role, err := receiver.roleRepo.FindByID(ctx, user.RoleID)
	if err != nil {
		pErr = err
		return
	}

	// generate token
	expire := time.Now().Add(receiver.expire)
	token, err = authz.NewToken(receiver.key, expire, user.ID, user.RoleID, role.RoleKey, user.NickName)
	if err != nil {
		pErr = pb.ErrorLoginFail("generate token failed: %s", err.Error())
		return
	}
	expireAt = expire.Unix()
	return
}
