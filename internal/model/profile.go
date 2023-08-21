package model

// Profile model.
type Profile struct {
	ID      int    `gorm:"primaryKey" json:"-"`
	UserID  int    `gorm:"<-:create" json:"-"`
	Address string `json:"name" valid:"required"`

	CreatedAt int `gorm:"autoCreateTime;<-:create" json:"created_at"`
	UpdatedAt int `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Profile) TableName() string {
	return "profile"
}
