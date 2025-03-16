package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	UUID     int64  `gorm:"column:uuid; type:int; not null; primaryKey:autoIncrement; comment: 'primary key is'"`
	RoleName string `gorm:"column:role_name"`
	RoleNote string `gorm:"column:role_note; type:text;"`
}

func (u *Role) TableName() string {
	return "go_db_role"
}
