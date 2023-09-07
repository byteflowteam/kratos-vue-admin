package biz

import (
	"context"
	"path/filepath"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/kakuilan/kgo"
	"gorm.io/gorm"

	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/conf"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/authz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/common"
	"github.com/byteflowteam/kratos-vue-admin/pkg/util"
)

var (
	ErrUserNotFound    = errors.New(400, "user not found", "user not found")
	ErrPasswordInvalid = errors.New(400, "password invalid", "user not found")
)

// SysUserRepo is a Greater repo.

type UserListCondition struct {
	UserName string
	Phone    string
	Status   int32
}

type SysUserRepo interface {
	Save(ctx context.Context, user *model.SysUser) (*model.SysUser, error)
	Delete(ctx context.Context, id int64) error
	UpdateByID(ctx context.Context, id int64, user *model.SysUser) error

	FindByID(ctx context.Context, id int64) (*model.SysUser, error)
	FindByUsername(ctx context.Context, username string) (*model.SysUser, error)
	FindByPostId(ctx context.Context, postId int64) ([]*model.SysUser, error)
	ListPage(ctx context.Context, page, size int32, condition UserListCondition) ([]*model.SysUser, error)
	Count(ctx context.Context, condition UserListCondition) (int32, error)
	CountByRoleId(ctx context.Context, roleId int64) (int64, error)
	FindAll(ctx context.Context) ([]*model.SysUser, error)
}

// SysUserUseCase is a SysUser use case.
type SysUserUseCase struct {
	userRepo   SysUserRepo
	uploadRepo OssRepo

	severConfig *conf.Server
	log         *log.Helper
}

// NewSysUserUseCase new a SysUser use case.
func NewSysUserUseCase(userRepo SysUserRepo, uploadRepo OssRepo, severConfig *conf.Server, logger log.Logger) *SysUserUseCase {
	return &SysUserUseCase{
		userRepo:    userRepo,
		uploadRepo:  uploadRepo,
		severConfig: severConfig,
		log:         log.NewHelper(logger),
	}
}

// CreateSysUser creates a SysUser, and returns the new SysUser.
func (uc *SysUserUseCase) CreateSysUser(ctx context.Context, u *model.SysUser) (*model.SysUser, error) {
	uc.log.WithContext(ctx).Infof("CreateSysUser: %v", u)
	u.Password = util.BcryptHash(u.Password)
	u.UUID = uuid.NewString()

	claims := authz.MustFromContext(ctx)
	u.CreateBy = claims.Nickname
	u.UpdateBy = claims.Nickname
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return uc.userRepo.Save(ctx, u)
}

func (uc *SysUserUseCase) UpdateSysUser(ctx context.Context, u *model.SysUser) error {
	uc.log.WithContext(ctx).Infof("UpdateSysUser: %v", u)
	oldUser, err := uc.userRepo.FindByID(ctx, u.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}
	claims := authz.MustFromContext(ctx)
	u.UUID = oldUser.UUID
	u.Salt = oldUser.Salt
	u.Password = oldUser.Password
	u.UpdateBy = claims.Nickname
	u.CreateBy = oldUser.CreateBy
	return uc.userRepo.UpdateByID(ctx, u.ID, u)
}

func (uc *SysUserUseCase) DeleteSysUser(ctx context.Context, id int64) error {
	return uc.userRepo.Delete(ctx, id)
}

func (uc *SysUserUseCase) FindSysUserById(ctx context.Context, id int64) (*model.SysUser, error) {
	return uc.userRepo.FindByID(ctx, id)
}

func (uc *SysUserUseCase) ListPage(ctx context.Context, req *pb.ListSysuserRequest) (users []*model.SysUser, total int32, err error) {
	var condition = UserListCondition{
		UserName: req.Username,
		Phone:    req.Phone,
		Status:   req.Status,
	}
	total, err = uc.userRepo.Count(ctx, condition)
	if err != nil {
		return
	}
	users, err = uc.userRepo.ListPage(ctx, req.PageNum, req.PageSize, condition)
	return
}

func (uc *SysUserUseCase) ChangeStatus(ctx context.Context, id int64, status int32) error {
	user, err := uc.FindSysUserById(ctx, id)
	if err != nil {
		return err
	}
	user.Status = status
	_, err = uc.userRepo.Save(ctx, user)
	return err
}

func (uc *SysUserUseCase) UpdatePassword(ctx context.Context, id int64, newPwd, oldPwd string) error {
	user, err := uc.FindSysUserById(ctx, id)
	if err != nil {
		return err
	}
	if !util.BcryptCheck(oldPwd, user.Password) {
		return ErrPasswordInvalid
	}

	user.Password = util.BcryptHash(newPwd)
	_, err = uc.userRepo.Save(ctx, user)
	return err
}

func (uc *SysUserUseCase) FindByPostId(ctx context.Context, postId int64) ([]*model.SysUser, error) {
	return uc.userRepo.FindByPostId(ctx, postId)
}

func (uc *SysUserUseCase) CountByRoleId(ctx context.Context, roleId int64) (int64, error) {
	return uc.userRepo.CountByRoleId(ctx, roleId)
}

func (uc *SysUserUseCase) FindSysUserByUsername(ctx context.Context, username string) (*model.SysUser, error) {
	return uc.userRepo.FindByUsername(ctx, username)
}

func (uc *SysUserUseCase) UpdateAvatar(ctx context.Context) error {
	claims := authz.MustFromContext(ctx)

	file, fileHeader, err := common.FormFile(ctx)
	if err != nil {
		return err
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	if !common.IsAllowedFileExt(ext) {
		return errors.New(401, "file type not allowed", "file type not allowed")
	}

	guid, _ := kgo.KStr.UuidV4()
	filePath := uc.severConfig.GetEnv().String() + "/userAvatar/" + guid + ext
	domain, err := uc.uploadRepo.UploadFile(file, filePath)
	if err != nil {
		return err
	}

	return uc.userRepo.UpdateByID(ctx, claims.UserID, &model.SysUser{
		UpdateBy:  claims.Nickname,
		UpdatedAt: time.Now(),
		Avatar:    domain + filePath,
	})
}

func (uc *SysUserUseCase) UploadFile(ctx context.Context) (string, error) {
	file, fileHeader, err := common.FormFile(ctx)
	if err != nil {
		return "", err
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	if !common.IsAllowedFileExt(ext) {
		return "", errors.New(401, "file type not allowed", "file type not allowed")
	}

	guid, _ := kgo.KStr.UuidV4()
	filePath := "files/" + guid + ext
	domain, err := uc.uploadRepo.UploadFile(file, filePath)
	if err != nil {
		return "", err
	}
	//return filePath, nil
	return domain + "/" + filePath, nil
}
