package methods

import (
	"gorm.io/gorm"
	e "td_roles.go/entities"
	u "td_roles.go/utils"
)

type AuthorizationM struct {
	Database *gorm.DB
}

type AuthorizationMethods interface {
	GetAuthorizationByID(id uint) (authorization e.Authorization, err error)
	GetAuthorizationBySlugName(slugName string) (authorization e.Authorization, err error)
	GetAuthorizationsByID(ids []uint) (authorizations e.Authorizations, err error)
	GetAuthorizationsBySlugName(slugName []string) (authorizations e.Authorizations, err error)
	CreateAuthorization(authorization *e.Authorization) error
	UpdateAuthorization(authorization e.Authorization, name string, description string) error
	DeleteAuthorizations(authorizations []e.Authorization) error
}

func (methods *AuthorizationM) GetAuthorizationByID(id uint) (authorization e.Authorization, err error) {
	err = methods.Database.Where("id=?", id).First(&authorization).Error
	return
}

func (methods *AuthorizationM) GetAuthorizationBySlugName(slugName string) (authorization e.Authorization, err error) {
	err = methods.Database.Where("slug_name=?", slugName).First(&authorization).Error
	return
}

func (methods *AuthorizationM) GetAuthorizationsByID(ids []uint) (authorizations e.Authorizations, err error) {
	err = methods.Database.Where("authorization.id IN (?)", ids).Find(&authorizations).Error
	return
}

func (methods *AuthorizationM) GetAuthorizationsBySlugName(slugName []string) (authorizations e.Authorizations, err error) {
	err = methods.Database.Where("authorization.slug_name IN (?)", slugName).Find(&authorizations).Error
	return
}

func (methods *AuthorizationM) CreateAuthorization(authorization *e.Authorization) error {
	return methods.Database.Create(&authorization).Error
}

func (methods *AuthorizationM) UpdateAuthorization(authorization e.Authorization, name string, description string) error {
	authorization.Name = name
	authorization.SlugName = u.SlugString(name)
	authorization.Description = description
	return methods.Database.Save(&authorization).Error
}

func (methods *AuthorizationM) DeleteAuthorizations(authorizations []e.Authorization) error {
	return methods.Database.Delete(&authorizations).Error
}
