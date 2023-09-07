package authz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/tx7do/kratos-casbin/authz"
)

var (
	ErrTokenMiss  = errors.New(500, "token miss", "token miss")
	ErrClaimsMiss = errors.New(500, "claims miss", "claims miss")
)

type TokenClaims struct {
	UserID   int64  `json:"user_id"`
	RoleID   int64  `json:"role_id"`
	RoleKey  string `json:"role_key"`
	Nickname string `json:"nickname"`
	jwtV4.RegisteredClaims
}

type securityUser struct {
	Path        string
	Method      string
	AuthorityId string
	Domain      string
}

func NewSecurityUser() authz.SecurityUser {
	return &securityUser{}
}

func (su *securityUser) ParseFromContext(ctx context.Context) error {
	claims, err := FromContext(ctx)
	if err != nil {
		return err
	}
	su.AuthorityId = claims.RoleKey
	ts, ok := transport.FromServerContext(ctx)
	if !ok {
		return ErrClaimsMiss
	}
	su.Path = ts.Operation()
	ht, ok := ts.(http.Transporter)
	if !ok {
		return ErrClaimsMiss
	}
	su.Method = ht.Request().Method
	return nil
}

func (su *securityUser) GetSubject() string {
	return su.AuthorityId
}

func (su *securityUser) GetObject() string {
	return su.Path
}

func (su *securityUser) GetAction() string {
	return su.Method
}

func (su *securityUser) GetDomain() string {
	return su.Domain
}

func FromContext(ctx context.Context) (*TokenClaims, error) {
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, ErrTokenMiss
	}
	return claims.(*TokenClaims), nil
}

func MustFromContext(ctx context.Context) *TokenClaims {
	claims, err := FromContext(ctx)
	if err != nil {
		panic(err)
	}
	return claims
}

func NewToken(key string, expireAt time.Time, userID, roleID int64, roleKey, nickname string) (string, error) {
	claims := jwtV4.NewWithClaims(jwtV4.SigningMethodHS256, &TokenClaims{
		UserID:   userID,
		RoleID:   roleID,
		Nickname: nickname,
		RoleKey:  roleKey,
		RegisteredClaims: jwtV4.RegisteredClaims{
			Issuer:    "admin",
			ExpiresAt: jwtV4.NewNumericDate(expireAt),
		},
	})
	return claims.SignedString([]byte(key))
}
