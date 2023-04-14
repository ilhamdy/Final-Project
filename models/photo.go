package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string   `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption  string   `json:"caption" form:"caption"`
	PhotoURL string   `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo_url is required"`
	UserID   uint     `gorm:"not null" json:"user_id"`
	User     *User    `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Comments *Comment ` json:"-"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
