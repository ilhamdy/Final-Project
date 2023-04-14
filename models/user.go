package models

import (
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string         `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required"`
	Email       string         `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required, email~Invalid email formal"`
	Password    string         `gorm:"not null" json:"password,omitempty" form:"password" valid:"required,minstringlength(6)~Password has to have minimum length of 6 characters"`
	Age         uint           `gorm:"not null" json:"age,omitempty" form:"age" valid:"required~Age is required,range(8|70)~Minimum age is 8 years old"`
	Photos      *[]Photo       `json:"-"`
	SocialMedia *[]SocialMedia `json:"-"`
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

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
