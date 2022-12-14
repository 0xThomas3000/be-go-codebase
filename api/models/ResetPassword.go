package models

import (
	"html"

	"strings"

	"gorm.io/gorm"
)

/*
  - A user might forget his/her password. When this happens, they can request to change to a new one.
    A notification will be sent to their email address with instructions to create a new password.
*/
type ResetPassword struct {
	gorm.Model
	Email string `gorm:"size:100;not null;" json:"email"`
	Token string `gorm:"size:255;not null;" json:"token"`
}

func (resetPassword *ResetPassword) Prepare() {
	resetPassword.Token = html.EscapeString(strings.TrimSpace(resetPassword.Token))
	resetPassword.Email = html.EscapeString(strings.TrimSpace(resetPassword.Email))
}

func (resetPassword *ResetPassword) SaveDatails(db *gorm.DB) (*ResetPassword, error) {
	var err error
	if err = db.Debug().Create(&resetPassword).Error; err != nil {
		return &ResetPassword{}, err
	}
	return resetPassword, nil
}

func (resetPassword *ResetPassword) DeleteDatails(db *gorm.DB) (int64, error) {

	db = db.Debug().Model(&ResetPassword{}).Where("id = ?", resetPassword.ID).Take(&ResetPassword{}).Delete(&ResetPassword{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
