package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  uint   `gorm:"not null" json:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id" form:"photo_id"`
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~message is required"`
	User    *User  `gorm:"foreignKey:UserID;constraint:opUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Photo   *Photo `gorm:"foreignKey:PhotoID;constraint:opUpdate:CASCADE,onDelete:CASCADE" json:"photo"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
