package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Title is required,maxstringlength(200)~Title maximum characters is 200"`
	UserID  uint   `gorm:"not null" json:"user_id" form:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id" form:"photo_id"`
	User    *User  `json:"user" form:"user"`
	Photo   *Photo `json:"photo" form:"photo"`
}

func (sm *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
