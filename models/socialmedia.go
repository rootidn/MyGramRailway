package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required,maxstringlength(50)~Name maximum characters is 50"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Social media URL is required"`
	UserID         uint   `gorm:"not null" json:"user_id" form:"user_id"`
	User           *User  `json:"user" form:"user"`
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
