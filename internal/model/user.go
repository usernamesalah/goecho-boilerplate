package model

import (
	"strings"

	"gorm.io/gorm"
)

// User model.
type User struct {
	ID       int    `gorm:"primaryKey" json:"-"`
	UID      string `json:"uid" valid:"required"`
	Name     string `json:"name" valid:"required"`
	Password string `json:"password" valid:"required,length(8|50)"`
	Email    string `json:"email" valid:"email,required"`

	CreatedAt int `gorm:"autoCreateTime;<-:create" json:"created_at"`
	UpdatedAt int `gorm:"autoUpdateTime" json:"updated_at"`

	Profile *Profile `gorm:"foreignKey:ProfileID" json:"profile"`
	Token   string   `gorm:"-" json:"token,omitempty"`
}

func (User) TableName() string {
	return "users"
}

func (m *User) setData(tx *gorm.DB) {
	m.Email = strings.ToLower(m.Email)
}

func (m *User) BeforeCreate(tx *gorm.DB) (err error) {
	m.setData(tx)
	return
}

func (m *User) BeforeUpdate(tx *gorm.DB) (err error) {
	m.setData(tx)
	return
}

func (m *User) ForPublic() {
	m.Password = ""
	m.UpdatedAt = 0
}
