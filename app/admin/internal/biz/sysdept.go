package biz

import (
	"context"
	"errors"
	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/byteflowteam/kratos-vue-admin/pkg/util"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/authz"
)

type SysDeptRepo interface {
	Create(ctx context.Context, dept *model.SysDept) error
	Save(ctx context.Context, dept *model.SysDept) error

	UpdateByID(ctx context.Context, id int64, dept *model.SysDept) error

	Delete(ctx context.Context, id int64) error

	FindByIDList(ctx context.Context, ids ...int64) ([]*model.SysDept, error)

	FindByID(ctx context.Context, id int64) (*model.SysDept, error)
	ListByNameStatusId(ctx context.Context, deptName string, status int32, id int64) ([]*model.SysDept, error)
	GetRoleDeptId(ctx context.Context, roleId int32) ([]int32, error)
	SelectDept(ctx context.Context) ([]*pb.DeptTree, error)
	SelectDeptLabel(ctx context.Context) ([]*pb.DeptLabel, error)
}

type SysDeptUseCase struct {
	repo SysDeptRepo
	log  *log.Helper
}

func NewSysDeptUseCase(repo SysDeptRepo, logger log.Logger) *SysDeptUseCase {
	return &SysDeptUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (d *SysDeptUseCase) ListByNameStatusId(ctx context.Context, deptName string, status int32, id int64) ([]*model.SysDept, error) {
	return d.repo.ListByNameStatusId(ctx, deptName, status, id)
}

func (d *SysDeptUseCase) CreateDept(ctx context.Context, sysDept *model.SysDept) (*model.SysDept, error) {
	claims := authz.MustFromContext(ctx)
	sysDept.CreateBy = claims.Nickname

	if err := d.repo.Create(ctx, sysDept); err != nil {
		return nil, err
	}

	// 更新deptPath
	deptPath, err := d.buildDeptPath(ctx, sysDept)
	if err != nil {
		return nil, err
	}
	sysDept.DeptPath = deptPath

	if err := d.repo.Save(ctx, sysDept); err != nil {
		d.log.WithContext(ctx).Errorf("save SysDept.deptPath error: %v", err)
	}
	return sysDept, nil
}

func (d *SysDeptUseCase) UpdateDept(ctx context.Context, sysDept *model.SysDept) (*model.SysDept, error) {
	oldDept, err := d.repo.FindByID(ctx, sysDept.ID)
	if err != nil {
		return nil, err
	}
	if oldDept.ParentID != sysDept.ParentID {
		return nil, errors.New("上级部门不允许修改")
	}

	claims := authz.MustFromContext(ctx)
	sysDept.UpdateBy = claims.Nickname

	deptPath, err := d.buildDeptPath(ctx, sysDept)
	if err != nil {
		return nil, err
	}
	sysDept.DeptPath = deptPath
	err = d.repo.UpdateByID(ctx, sysDept.ID, sysDept)
	return sysDept, err
}

func (d *SysDeptUseCase) DeleteDept(ctx context.Context, id int64) error {
	return d.repo.Delete(ctx, id)
}

func (d *SysDeptUseCase) buildDeptPath(ctx context.Context, dept *model.SysDept) (string, error) {
	deptPath := "/" + strconv.FormatInt(dept.ID, 10)
	if dept.ParentID == 0 {
		deptPath = "/0" + deptPath
	} else {
		parent, err := d.repo.FindByID(ctx, dept.ParentID)
		if err != nil {
			return "", err
		}
		deptPath = parent.DeptPath + deptPath
	}
	return deptPath, nil
}

func (d *SysDeptUseCase) GetDept(ctx context.Context, id int64) (*model.SysDept, error) {
	return d.repo.FindByID(ctx, id)
}

func (d *SysDeptUseCase) GetDeptList(ctx context.Context) ([]*pb.DeptTree, error) {
	return d.repo.SelectDept(ctx)
}

func (d *SysDeptUseCase) RoleDeptTreeSelect(ctx context.Context, roleId int32) (*pb.RoleDeptTreeSelectReply, error) {
	var err error
	result, err := d.repo.SelectDeptLabel(ctx)
	if err != nil {
		return nil, err
	}
	deptIds := make([]int32, 0)

	if roleId != 0 {
		deptIds, err = d.repo.GetRoleDeptId(ctx, roleId)
		if err != nil {
			return nil, err
		}
	}
	reply := &pb.RoleDeptTreeSelectReply{
		CheckedKeys: deptIds,
		Depts:       result,
	}
	return reply, nil
}

func (d *SysDeptUseCase) FindDeptByIDList(ctx context.Context, ids []int64) ([]*model.SysDept, error) {
	if len(ids) == 0 {
		return []*model.SysDept{}, nil
	}
	return d.repo.FindByIDList(ctx, ids...)
}

func ConvertToDeptTree(deptList []*model.SysDept) []*pb.DeptTree {
	menuMap := make(map[int64]*pb.DeptTree)

	for _, dept := range deptList {
		menuTree := &pb.DeptTree{
			DeptId:     dept.ID,
			ParentId:   dept.ParentID,
			DeptPath:   dept.DeptPath,
			DeptName:   dept.DeptName,
			Sort:       dept.Sort,
			Leader:     dept.Leader,
			Phone:      dept.Phone,
			Email:      dept.Email,
			Status:     dept.Status,
			CreateBy:   dept.CreateBy,
			UpdateBy:   dept.UpdateBy,
			CreateTime: util.NewTimestamp(dept.CreatedAt),
			UpdateTime: util.NewTimestamp(dept.UpdatedAt),
			Children:   []*pb.DeptTree{},
		}
		menuMap[dept.ID] = menuTree
	}

	var root []*pb.DeptTree
	for _, menuTree := range menuMap {
		parentID := menuTree.ParentId
		if parentID == 0 {
			root = append(root, menuTree)
		} else {
			parent, ok := menuMap[parentID]
			if ok {
				parent.Children = append(parent.Children, menuTree)
			}
		}
	}

	return root
}

func ConvertToDeptTreeChildren(deptList []*pb.DeptTree) []*pb.DeptTree {
	menuMap := make(map[int64]*pb.DeptTree)

	for _, dept := range deptList {
		menuMap[dept.DeptId] = dept
	}

	var root []*pb.DeptTree
	for _, menuTree := range menuMap {
		parentID := menuTree.ParentId
		if parentID == 0 {
			root = append(root, menuTree)
		} else {
			parent, ok := menuMap[parentID]
			if ok {
				parent.Children = append(parent.Children, menuTree)
			}
		}
	}

	return root
}
