package models

type RoleBaseModel struct {
	Id    uint            `gorm:"primary_key"`
	Name  string          `gorm:"column:name"`
	Users []UserBaseModel `gorm:"many2many:user_roles;"`
}
