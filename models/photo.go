package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string    `gorm:"not null" json:"title" form:"title" valid:"required~Title is required,maxstringlength(100)~Title maximum characters is 100"`
	Caption  string    `json:"caption" form:"caption" valid:"maxstringlength(200)~Username maximum characters is 200"`
	PhotoUrl string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Title is required"`
	UserID   uint      `gorm:"not null" json:"user_id" form:"user_id"`
	User     *User     `json:"user,omitempty" form:"user"`
	Comment  []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments,omitempty"`
}

func (sm *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
