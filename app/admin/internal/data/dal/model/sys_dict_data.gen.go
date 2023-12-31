// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysDictDatum = "sys_dict_data"

// SysDictDatum mapped from table <sys_dict_data>
type SysDictDatum struct {
	DictCode   int64     `gorm:"column:dict_code;type:bigint(20);primaryKey;autoIncrement:true" json:"dict_code"`
	DictSort   int32     `gorm:"column:dict_sort;type:int(11);comment:排序" json:"dict_sort"`
	DictLabel  string    `gorm:"column:dict_label;type:varchar(64);comment:标签" json:"dict_label"`
	DictValue  string    `gorm:"column:dict_value;type:varchar(64);comment:值" json:"dict_value"`
	DictType   string    `gorm:"column:dict_type;type:varchar(64);comment:字典类型" json:"dict_type"`
	Status     int32     `gorm:"column:status;type:tinyint(2);comment:状态（0正常 1停用）" json:"status"`
	CSSClass   string    `gorm:"column:css_class;type:varchar(128);comment:CssClass" json:"css_class"`
	ListClass  string    `gorm:"column:list_class;type:varchar(128);comment:ListClass" json:"list_class"`
	IsDefault  string    `gorm:"column:is_default;type:varchar(8);comment:IsDefault" json:"is_default"`
	CreateBy   string    `gorm:"column:create_by;type:varchar(191)" json:"create_by"`
	UpdateBy   string    `gorm:"column:update_by;type:varchar(191)" json:"update_by"`
	Remark     string    `gorm:"column:remark;type:varchar(256);comment:备注" json:"remark"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
	DeleteTime time.Time `gorm:"column:delete_time;type:datetime" json:"delete_time"`
}

// TableName SysDictDatum's table name
func (*SysDictDatum) TableName() string {
	return TableNameSysDictDatum
}
