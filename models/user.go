package models

import (
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username        string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required,maxstringlength(50)~Username maximum characters is 50"`
	Email           string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Invalid email format,maxstringlength(150)~Email maximum characters is 150"`
	Password        string        `gorm:"not null" json:"password,omitempty" form:"password" valid:"required~Your password is required,minstringlength(6)~Password minimum characters is 6"`
	Age             int           `gorm:"not null" json:"age,omitempty" form:"age" valid:"required~Age is required"`
	ProfileImageURL string        `json:"profile_image_url,omitempty" form:"profile_image_url"`
	Photo           []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos,omitempty"`
	Comment         []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments,omitempty"`
	SocialMedia     []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"social_medias,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
