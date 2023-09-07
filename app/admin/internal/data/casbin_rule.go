package data

import (
	"context"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/conf"
)

type casbinRuleRepo struct {
	data           *Data
	log            *log.Helper
	syncedEnforcer *casbin.SyncedEnforcer

	ModelPath string
}

func NewCasbinRuleRepo(data *Data, db *gorm.DB, conf *conf.Casbin, logger log.Logger) biz.CasbinRuleRepo {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic("新建权限适配器失败1:\n" + err.Error())
	}
	syncedEnforcer, err := casbin.NewSyncedEnforcer(conf.Path, adapter)
	if err != nil {
		panic("新建权限适配器失败2:\n" + err.Error())
	}
	if err := syncedEnforcer.LoadPolicy(); err != nil {
		panic("新建权限适配器失败3:\n" + err.Error())
	}
	return &casbinRuleRepo{
		data:           data,
		log:            log.NewHelper(logger),
		syncedEnforcer: syncedEnforcer,
		ModelPath:      conf.Path,
	}
}

// UpdateCasbin
// RoleKey = v0
// Path = v1
// Method = v2
func (c *casbinRuleRepo) UpdateCasbin(ctx context.Context, roleKey string, rules [][]string) error {
	if err := c.ClearCasbin(0, roleKey); err != nil {
		return err
	}
	_ = c.syncedEnforcer.LoadPolicy()
	success, err := c.syncedEnforcer.AddPolicies(rules)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	_ = c.syncedEnforcer.SavePolicy()
	return nil
}

func (c *casbinRuleRepo) UpdateCasbinApi(ctx context.Context, oldPath string, newPath string, oldMethod string, newMethod string) error {
	q := c.data.Query(ctx).CasbinRule
	_, err := q.WithContext(ctx).Where(q.V1.Eq(oldPath), q.V2.Eq(oldMethod)).UpdateColumns(map[string]any{
		"v1": newPath,
		"v2": newMethod,
	})
	return err
}

func (c *casbinRuleRepo) GetPolicyPathByRoleId(roleKey string) [][]string {
	return c.syncedEnforcer.GetFilteredPolicy(0, roleKey)
}

func (c *casbinRuleRepo) ClearCasbin(v int, p ...string) error {
	_, err := c.syncedEnforcer.RemoveFilteredPolicy(v, p...)
	return err
}

func (c *casbinRuleRepo) GetModel() model.Model {
	return c.syncedEnforcer.GetModel()
}

func (c *casbinRuleRepo) GetAdapter() persist.Adapter {
	return c.syncedEnforcer.GetAdapter()
}
