package biz

import (
	"context"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
)

type CasbinRuleRepo interface {
	GetModel() model.Model
	GetAdapter() persist.Adapter

	UpdateCasbin(ctx context.Context, roleKey string, p [][]string) error
	ClearCasbin(v int, p ...string) error
	UpdateCasbinApi(ctx context.Context, oldPath string, newPath string, oldMethod string, newMethod string) error
	GetPolicyPathByRoleId(roleKey string) [][]string
}

type CasbinRuleUseCase struct {
	repo CasbinRuleRepo
	log  *log.Helper
}

func NewCasbinRuleUseCase(repo CasbinRuleRepo, logger log.Logger) *CasbinRuleUseCase {
	return &CasbinRuleUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c *CasbinRuleUseCase) UpdateCasbin(ctx context.Context, roleKey string, apis []*pb.ApiBase) error {
	var rules [][]string
	for _, api := range apis {
		rules = append(rules, []string{roleKey, api.Path, api.Method})
	}
	return c.repo.UpdateCasbin(ctx, roleKey, rules)
}

func (c *CasbinRuleUseCase) GetPolicyPathByRoleId(roleKey string) [][]string {
	return c.repo.GetPolicyPathByRoleId(roleKey)
}

func (c *CasbinRuleUseCase) ClearCasbin(roleKey string) error {
	return c.repo.ClearCasbin(0, roleKey)
}
