package methods

import (
	"gorm.io/gorm"
	e "td_roles.go/entities"
	u "td_roles.go/utils"
)

type RoleM struct {
	Database *gorm.DB
}

type RoleMethods interface {
	GetRoleByID(id uint, withAuthorization bool) (role e.Role, err error)
	GetRoleBySlugName(slugName string, withAuthorization bool) (role e.Role, err error)
	GetRolesByID(ids []uint, withAuthorization bool) (roles []e.Role, err error)
	GetRolesBySlugName(slugName []string, withAuthorization bool) (roles []e.Role, err error)
	CreateRole(role *e.Role) error
	UpdateRole(role e.Role, name string, description string) error
	DeleteRoles(roles []e.Role) error

	AddAuthorizations(role *e.Role, authorizations e.Authorizations) error
	ReplaceAuthorizations(role *e.Role, authorizations e.Authorizations) error
	DelAuthorizations(role *e.Role, authorizations e.Authorizations) error
}

func (methods *RoleM) GetRoleByID(id uint, withAuthorization bool) (role e.Role, err error) {
	if withAuthorization {
		err = methods.Database.Preload("Authorizations").Where("id=?)", id).First(&role).Error
	} else {
		err = methods.Database.Where("id=?", id).First(&role).Error
	}
	return
}

func (methods *RoleM) GetRoleBySlugName(slugName string, withAuthorization bool) (role e.Role, err error) {
	if withAuthorization {
		err = methods.Database.Preload("Authorizations").Where("slug_name=?", slugName).First(&role).Error
	} else {
		err = methods.Database.Where("slug_name=?", slugName).First(&role).Error
	}
	return
}

func (methods *RoleM) GetRolesByID(ids []uint, withAuthorization bool) (roles []e.Role, err error) {
	if withAuthorization {
		err = methods.Database.Preload("Authorizations").Where("roles.id IN (?)", ids).Find(&roles).Error
	} else {
		err = methods.Database.Where("roles.id IN (?)", ids).Find(&roles).Error
	}
	return
}

func (methods *RoleM) GetRolesBySlugName(slugName []string, withAuthorization bool) (roles []e.Role, err error) {
	if withAuthorization {
		err = methods.Database.Preload("Authorizations").Where("roles.slug_name IN (?)", slugName).Find(&roles).Error
	} else {
		err = methods.Database.Where("roles.slug_name IN (?)", slugName).Find(&roles).Error
	}
	return
}

func (methods *RoleM) CreateRole(role *e.Role) error {
	return methods.Database.Create(&role).Error
}

func (methods *RoleM) UpdateRole(role e.Role, name string, description string) error {
	role.Name = name
	role.SlugName = u.SlugString(name)
	role.Description = description
	return methods.Database.Save(&role).Error
}

func (methods *RoleM) DeleteRoles(roles []e.Role) error {
	return methods.Database.Delete(&roles).Error
}

func (methods *RoleM) AddAuthorizations(role *e.Role, authorizations e.Authorizations) error {
	return methods.Database.Model(role).Association("Authorizations").Append(authorizations.Origin())
}

func (methods *RoleM) ReplaceAuthorizations(role *e.Role, authorizations e.Authorizations) error {
	return methods.Database.Model(role).Association("Authorizations").Replace(authorizations.Origin())
}

func (methods *RoleM) DelAuthorizations(role *e.Role, authorizations e.Authorizations) error {
	return methods.Database.Model(role).Association("Authorizations").Delete(authorizations.Origin())
}
