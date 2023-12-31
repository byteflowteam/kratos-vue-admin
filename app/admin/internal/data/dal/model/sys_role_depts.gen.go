// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysRoleDept = "sys_role_depts"

// SysRoleDept mapped from table <sys_role_depts>
type SysRoleDept struct {
	ID     int64 `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true;comment:主键id" json:"id"`
	RoleID int64 `gorm:"column:role_id;type:bigint(20);not null;comment:角色id" json:"role_id"`
	DeptID int64 `gorm:"column:dept_id;type:bigint(20);not null;comment:部门id" json:"dept_id"`
}

// TableName SysRoleDept's table name
func (*SysRoleDept) TableName() string {
	return TableNameSysRoleDept
}
