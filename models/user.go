package models

import (
	"challenge294/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string    `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Role     string    `json:"role" form:"role"`
	Products []Product `gorm:constraint;OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
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
