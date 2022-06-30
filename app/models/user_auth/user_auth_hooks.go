package user_auth

import (
	"gohub/pkg/hash"

	"gorm.io/gorm"
)

func (userAuth *UserAuth) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(userAuth.Credential) {
		userAuth.Credential = hash.BcryptHash(userAuth.Credential)
	}
	return
}

// func (userAuth *UserAuth) BeforeCreate(tx *gorm.DB) (err error) {}
// func (userAuth *UserAuth) AfterCreate(tx *gorm.DB) (err error) {}
// func (userAuth *UserAuth) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (userAuth *UserAuth) AfterUpdate(tx *gorm.DB) (err error) {}
// func (userAuth *UserAuth) AfterSave(tx *gorm.DB) (err error) {}
// func (userAuth *UserAuth) BeforeDelete(tx *gorm.DB) (err error) {}
// func (userAuth *UserAuth) AfterDelete(tx *gorm.DB) (err error) {}
// func (userAuth *UserAuth) AfterFind(tx *gorm.DB) (err error) {}
