package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
)

type SysApiRepo interface {
	FindByID(ctx context.Context, id int64) (*model.SysAPI, error)
	Create(ctx context.Context, api *model.SysAPI) error
	Save(ctx context.Context, api *model.SysAPI) error
	Delete(ctx context.Context, id int64) error
	FindAll(ctx context.Context) ([]*model.SysAPI, error)
	ListPage(ctx context.Context, page, size int32) ([]*model.SysAPI, error)
	ListPageCount(ctx context.Context) (int32, error)
}

type SysApiUseCase struct {
	apiRepo    SysApiRepo
	casbinRepo CasbinRuleRepo
	log        *log.Helper
}

func NewSysApiUseCase(repo SysApiRepo, casbinRepo CasbinRuleRepo, logger log.Logger) *SysApiUseCase {
	return &SysApiUseCase{
		apiRepo:    repo,
		casbinRepo: casbinRepo,
		log:        log.NewHelper(logger),
	}
}

func (a *SysApiUseCase) ListPage(ctx context.Context, page, size int32) ([]*model.SysAPI, int32, error) {
	total, err := a.apiRepo.ListPageCount(ctx)
	if err != nil || total == 0 {
		return nil, 0, err
	}
	apis, err := a.apiRepo.ListPage(ctx, page, size)
	return apis, total, err
}

func (a *SysApiUseCase) AllApi(ctx context.Context) ([]*model.SysAPI, error) {
	return a.apiRepo.FindAll(ctx)
}

func (a *SysApiUseCase) CreateApi(ctx context.Context, api *model.SysAPI) (*model.SysAPI, error) {
	err := a.apiRepo.Create(ctx, api)
	return api, err
}

func (a *SysApiUseCase) UpdateApi(ctx context.Context, api *model.SysAPI) (*model.SysAPI, error) {
	err := a.apiRepo.Save(ctx, api)
	return api, err
}

func (a *SysApiUseCase) DeleteApi(ctx context.Context, id int64) error {
	return a.apiRepo.Delete(ctx, id)
}

func (a *SysApiUseCase) GetPolicyPathByRoleKey(ctx context.Context, roleKey string) ([][]string, error) {
	return a.casbinRepo.GetPolicyPathByRoleId(roleKey), nil
}

func (a *SysApiUseCase) GetApiByID(ctx context.Context, id int64) (*model.SysAPI, error) {
	return a.apiRepo.FindByID(ctx, id)
}
