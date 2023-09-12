package service

import (
	"context"
	"github.com/byteflowteam/kratos-vue-admin/pkg/common/constant"
	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/conf"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/authz"
	"github.com/byteflowteam/kratos-vue-admin/pkg/util"
)

type SysuserService struct {
	pb.UnimplementedSysuserServer
	serverConf   *conf.Server
	userCase     *biz.SysUserUseCase
	authCase     *biz.AuthUseCase
	roleCase     *biz.SysRoleUseCase
	roleMenuCase *biz.SysRoleMenuUseCase
	postCase     *biz.SysPostUseCase
	deptCase     *biz.SysDeptUseCase
	log          *log.Helper
}

func NewSysuserService(serverConf *conf.Server, userCase *biz.SysUserUseCase, authCase *biz.AuthUseCase, roleCase *biz.SysRoleUseCase, roleMenuCase *biz.SysRoleMenuUseCase, postCase *biz.SysPostUseCase, deptCase *biz.SysDeptUseCase, logger log.Logger) *SysuserService {
	return &SysuserService{
		serverConf:   serverConf,
		userCase:     userCase,
		authCase:     authCase,
		roleCase:     roleCase,
		roleMenuCase: roleMenuCase,
		postCase:     postCase,
		deptCase:     deptCase,
		log:          log.NewHelper(log.With(logger, "module", "service/sysuser")),
	}
}

func (s *SysuserService) CreateSysuser(ctx context.Context, req *pb.CreateSysuserRequest) (*pb.CreateSysuserReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	_, err := s.userCase.CreateSysUser(ctx, &model.SysUser{
		NickName: req.NickName,
		Phone:    req.Phone,
		RoleID:   req.RoleId,
		Avatar:   req.Avatar,
		Sex:      req.Sex,
		Email:    req.Email,
		DeptID:   req.DeptId,
		PostID:   req.PostId,
		Remark:   req.Remark,
		Status:   req.Status,
		Username: req.Username,
		Password: req.Password,
		RoleIds:  req.RoleIds,
		PostIds:  req.PostIds,
		Secret:   req.Secret,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateSysuserReply{}, nil
}

func (s *SysuserService) UpdateSysuser(ctx context.Context, req *pb.UpdateSysuserRequest) (*pb.UpdateSysuserReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	err := s.userCase.UpdateSysUser(ctx, &model.SysUser{
		ID:       req.UserId,
		NickName: req.NickName,
		Phone:    req.Phone,
		RoleID:   req.RoleId,
		Avatar:   req.Avatar,
		Sex:      req.Sex,
		Email:    req.Email,
		DeptID:   req.DeptId,
		PostID:   req.PostId,
		Remark:   req.Remark,
		Status:   req.Status,
		Username: req.Username,
		Password: req.Password,
		RoleIds:  req.RoleIds,
		PostIds:  req.PostIds,
		Secret:   req.Secret,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateSysuserReply{}, nil
}

func (s *SysuserService) DeleteSysuser(ctx context.Context, req *pb.DeleteSysuserRequest) (*pb.DeleteSysuserReply, error) {
	err := s.userCase.DeleteSysUser(ctx, req.Id)
	return &pb.DeleteSysuserReply{}, err
}

func (s *SysuserService) GetSysuser(ctx context.Context, req *pb.GetSysuserRequest) (*pb.GetSysuserReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, err := s.userCase.FindSysUserById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	role, err := s.roleCase.GetRole(ctx, user.RoleID)
	if err != nil {
		return nil, err
	}

	roles, err := s.roleCase.FindRoleByIDList(ctx, util.Split2Int64Slice(user.RoleIds))
	if err != nil {
		return nil, err
	}
	posts, err := s.postCase.FindPostByIDList(ctx, util.Split2Int64Slice(user.PostIds))
	if err != nil {
		return nil, err
	}

	deptList, err := s.deptCase.GetDeptList(ctx)
	if err != nil {
		return nil, err
	}

	replyRole := make([]*pb.RoleData, len(roles))
	for i, d := range roles {
		replyRole[i] = &pb.RoleData{
			RoleId:     d.ID,
			RoleName:   d.RoleName,
			Status:     int64(d.Status),
			RoleKey:    d.RoleKey,
			RoleSort:   d.RoleSort,
			DataScope:  int64(d.DataScope),
			CreateBy:   d.CreateBy,
			UpdateBy:   d.UpdateBy,
			Remark:     d.Remark,
			CreateTime: util.NewTimestamp(d.CreatedAt),
			UpdateTime: util.NewTimestamp(d.UpdatedAt),
		}
	}

	replyPost := make([]*pb.PostData, len(posts))
	for i, d := range posts {
		replyPost[i] = &pb.PostData{
			PostId:     d.ID,
			PostName:   d.PostName,
			PostCode:   d.PostCode,
			Sort:       d.Sort,
			Status:     d.Status,
			Remark:     d.Remark,
			CreateBy:   d.CreateBy,
			UpdateBy:   d.UpdateBy,
			CreateTime: util.NewTimestamp(d.CreatedAt),
			UpdateTime: util.NewTimestamp(d.UpdatedAt),
		}
	}

	replyUser := &pb.UserData{
		UserId:     user.ID,
		NickName:   user.NickName,
		Phone:      user.Phone,
		RoleId:     int32(user.RoleID),
		Avatar:     user.Avatar,
		Sex:        int64(user.Sex),
		Email:      user.Email,
		DeptId:     int32(user.DeptID),
		PostId:     int32(user.PostID),
		RoleIds:    user.RoleIds,
		PostIds:    user.PostIds,
		CreateBy:   user.CreateBy,
		UpdateBy:   user.UpdateBy,
		Remark:     user.Remark,
		Status:     user.Status,
		Username:   user.Username,
		RoleName:   role.RoleName,
		CreateTime: util.NewTimestamp(user.CreatedAt),
		UpdateTime: util.NewTimestamp(user.UpdatedAt),
		Secret:     user.Secret,
		Qrcode:     util.NewGoogleAuth().GetQrcode(user.Secret),
	}

	replyDepts := biz.ConvertToDeptTreeChildren(deptList)
	reply := &pb.GetSysuserReply{
		User:    replyUser,
		Roles:   replyRole,
		Posts:   replyPost,
		Depts:   replyDepts,
		PostIds: replyUser.PostIds,
		RoleIds: replyUser.RoleIds,
	}
	return reply, nil
}

func (s *SysuserService) ListSysuser(ctx context.Context, req *pb.ListSysuserRequest) (*pb.ListSysuserReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	users, total, err := s.userCase.ListPage(ctx, req)
	if err != nil {
		return nil, err
	}

	deptCache := util.NewCache(func(id int64) (*model.SysDept, error) {
		d, err := s.deptCase.GetDept(ctx, id)
		if d == nil {
			d = &model.SysDept{}
		}
		return d, err
	})
	roleCache := util.NewCache(func(id int64) (*model.SysRole, error) {
		d, err := s.roleCase.GetRole(ctx, id)
		if d == nil {
			d = &model.SysRole{}
		}
		return d, err
	})

	gAuth := util.NewGoogleAuth()
	replyData := make([]*pb.UserData, len(users))
	for i, user := range users {
		role, _ := roleCache.Get(user.RoleID)
		dept, _ := deptCache.Get(user.DeptID)
		replyData[i] = &pb.UserData{
			UserId:     user.ID,
			NickName:   user.NickName,
			Phone:      user.Phone,
			RoleId:     int32(user.RoleID),
			Avatar:     user.Avatar,
			Sex:        int64(user.Sex),
			Email:      user.Email,
			DeptId:     int32(user.DeptID),
			PostId:     int32(user.PostID),
			RoleIds:    user.RoleIds,
			PostIds:    user.PostIds,
			CreateBy:   user.CreateBy,
			UpdateBy:   user.UpdateBy,
			Remark:     user.Remark,
			Status:     user.Status,
			CreateTime: util.NewTimestamp(user.CreatedAt),
			UpdateTime: util.NewTimestamp(user.UpdatedAt),
			Username:   user.Username,
			RoleName:   role.RoleName,
			DeptName:   dept.DeptName,
			Secret:     user.Secret,
			Qrcode:     gAuth.GetQrcode(user.Secret),
		}
	}

	return &pb.ListSysuserReply{
		Total:    total,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Data:     replyData,
	}, nil
}

func (s *SysuserService) GetCaptcha(context.Context, *pb.GetCaptchaRequest) (*pb.GetCaptchaReply, error) {
	id, content, image := util.Generate()
	if s.serverConf.GetEnv() != conf.Env_dev {
		content = ""
	}
	return &pb.GetCaptchaReply{
		Base64Captcha: image,
		CaptchaId:     id,
		Content:       content,
	}, nil
}

func (s *SysuserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	token, expireAt, err := s.authCase.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		Token:  token,
		Expire: expireAt,
	}, nil
}

func (s *SysuserService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}

// Auth 用户权限信息
func (s *SysuserService) Auth(ctx context.Context, req *pb.AuthRequest) (*pb.AuthReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, err := s.userCase.FindSysUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	role, err := s.roleCase.GetRole(ctx, user.RoleID)
	if err != nil {
		return nil, err
	}

	permits, err := s.roleMenuCase.GetPermission(ctx, role.ID)
	if err != nil {
		return nil, err
	}

	var menus []*pb.MenuTree
	// 被禁用了，菜单显示空
	if role.Status == constant.StatusMenusForbidden {
		menus = make([]*pb.MenuTree, 0)
	} else {
		menus, err = s.roleMenuCase.SelectMenuRole(ctx, role.RoleName)
		if err != nil {
			return nil, err
		}
	}

	pbUser := &pb.AuthReply_User{
		UserId:    user.ID,
		NickName:  user.NickName,
		Phone:     user.Phone,
		RoleId:    user.RoleID,
		Avatar:    user.Avatar,
		Sex:       user.Sex,
		Email:     user.Email,
		DeptId:    user.DeptID,
		PostId:    user.PostID,
		RoleIds:   user.RoleIds,
		PostIds:   user.PostIds,
		CreateBy:  user.CreateBy,
		UpdateBy:  user.UpdateBy,
		Remark:    user.Remark,
		Status:    user.Status,
		CreatedAt: util.NewTimestamp(user.CreatedAt),
		UpdatedAt: util.NewTimestamp(user.UpdatedAt),
		Username:  user.Username,
		RoleName:  role.RoleName,
	}

	pbRole := &pb.AuthReply_Role{
		RoleId:    role.ID,
		RoleName:  role.RoleName,
		Status:    role.Status,
		RoleKey:   role.RoleKey,
		RoleSort:  role.RoleSort,
		DataScope: role.DataScope,
		CreateBy:  role.CreateBy,
		UpdateBy:  role.UpdateBy,
		Remark:    role.Remark,
		ApiIds:    nil,
		MenuIds:   nil,
		DeptIds:   nil,
		CreatedAt: util.NewTimestamp(user.CreatedAt),
		UpdatedAt: util.NewTimestamp(user.UpdatedAt),
	}

	return &pb.AuthReply{
		User:        pbUser,
		Role:        pbRole,
		Permissions: permits,
		Menus:       Build(menus),
	}, nil
}

func (s *SysuserService) ChangeStatus(ctx context.Context, req *pb.ChangeStatusRequest) (*pb.ChangeStatusReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	err := s.userCase.ChangeStatus(ctx, req.UserId, req.Status)
	return &pb.ChangeStatusReply{}, err
}

func (s *SysuserService) UpdateAvatar(ctx context.Context) error {
	return s.userCase.UpdateAvatar(ctx)
}

func (s *SysuserService) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	claims := authz.MustFromContext(ctx)
	err := s.userCase.UpdatePassword(ctx, claims.UserID, req.NewPassword, req.OldPassword)
	return &pb.UpdatePasswordReply{}, err
}

// GetPostInit 获取初始化角色岗位信息
func (s *SysuserService) GetPostInit(ctx context.Context, req *pb.GetPostInitRequest) (*pb.GetPostInitReply, error) {
	// 获取所有角色
	roleList, err := s.roleCase.FindRoleAll(ctx)
	if err != nil {
		return nil, err
	}
	// 获取所有岗位
	postList, err := s.postCase.FindPostAll(ctx)
	if err != nil {
		return nil, err
	}

	replyRoles := make([]*pb.RoleData, len(roleList))
	for i, d := range roleList {
		replyRoles[i] = &pb.RoleData{
			RoleId:     d.ID,
			RoleName:   d.RoleName,
			Status:     int64(d.Status),
			RoleKey:    d.RoleKey,
			RoleSort:   d.RoleSort,
			DataScope:  int64(d.DataScope),
			CreateBy:   d.CreateBy,
			UpdateBy:   d.UpdateBy,
			Remark:     d.Remark,
			CreateTime: util.NewTimestamp(d.CreatedAt),
			UpdateTime: util.NewTimestamp(d.UpdatedAt),
		}
	}

	replyPosts := make([]*pb.PostData, len(postList))
	for i, d := range postList {
		replyPosts[i] = &pb.PostData{
			PostId:     d.ID,
			PostName:   d.PostName,
			PostCode:   d.PostCode,
			Sort:       d.Sort,
			Status:     d.Status,
			Remark:     d.Remark,
			CreateBy:   d.CreateBy,
			UpdateBy:   d.UpdateBy,
			CreateTime: util.NewTimestamp(d.CreatedAt),
			UpdateTime: util.NewTimestamp(d.UpdatedAt),
		}
	}

	return &pb.GetPostInitReply{
		Roles: replyRoles,
		Posts: replyPosts,
	}, nil
}

// GetUserRolePost 获取用户角色岗位信息
func (s *SysuserService) GetUserRolePost(ctx context.Context, req *pb.GetUserRolePostRequest) (*pb.GetUserRolePostReply, error) {
	claims := authz.MustFromContext(ctx)
	user, err := s.userCase.FindSysUserById(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}
	roleIds := util.Split2Int64Slice(user.RoleIds)
	postIds := util.Split2Int64Slice(user.PostIds)

	roleList, err := s.roleCase.FindRoleByIDList(ctx, roleIds)
	if err != nil {
		return nil, err
	}
	postList, err := s.postCase.FindPostByIDList(ctx, postIds)
	if err != nil {
		return nil, err
	}

	replyRoles := make([]*pb.RoleData, len(roleList))
	for i, d := range roleList {
		replyRoles[i] = &pb.RoleData{
			RoleId:     d.ID,
			RoleName:   d.RoleName,
			Status:     int64(d.Status),
			RoleKey:    d.RoleKey,
			RoleSort:   d.RoleSort,
			DataScope:  int64(d.DataScope),
			CreateBy:   d.CreateBy,
			UpdateBy:   d.UpdateBy,
			Remark:     d.Remark,
			CreateTime: util.NewTimestamp(d.CreatedAt),
			UpdateTime: util.NewTimestamp(d.UpdatedAt),
		}
	}

	replyPosts := make([]*pb.PostData, len(postList))
	for i, d := range postList {
		replyPosts[i] = &pb.PostData{
			PostId:     d.ID,
			PostName:   d.PostName,
			PostCode:   d.PostCode,
			Sort:       d.Sort,
			Status:     d.Status,
			Remark:     d.Remark,
			CreateBy:   d.CreateBy,
			UpdateBy:   d.UpdateBy,
			CreateTime: util.NewTimestamp(d.CreatedAt),
			UpdateTime: util.NewTimestamp(d.UpdatedAt),
		}
	}

	return &pb.GetUserRolePostReply{
		Roles: replyRoles,
		Posts: replyPosts,
	}, err
}

func (s *SysuserService) GetUserGoogleSecret(ctx context.Context, req *pb.GetUserGoogleSecretRequest) (*pb.GetUserGoogleSecretReply, error) {
	gAuth := util.NewGoogleAuth()
	secret := gAuth.GetSecret()
	qrcode := gAuth.GetQrcode(secret)
	var rep = &pb.GetUserGoogleSecretReply{}
	rep.Secret = secret
	rep.Qrcode = qrcode
	return rep, nil
}

func (s *SysuserService) UploadFile(ctx context.Context) (string, error) {
	return s.userCase.UploadFile(ctx)
}
