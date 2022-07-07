package methods

import (
	"gorm.io/gorm"
	e "td_roles.go/entities"
)

type UserM struct {
	Database *gorm.DB
}

type UserMethods interface {
	AddUserRole(userID string, roles []e.Role) error
	DeleteUserRole(userID string, roles []e.Role) (err error)
	//DeleteUserRoles(roles []e.Role) (err error)
}

func (methods *UserM) AddUserRole(userID string, roles []e.Role) error {
	var userRoles []e.UserRole
	for _, role := range roles {
		userRoles = append(userRoles, e.UserRole{
			UserID: userID,
			RoleID: role.ID,
		})
	}
	return methods.Database.Create(&userRoles).Error
}

func (methods *UserM) DeleteUserRole(userID string, roles []e.Role) (err error) {
	var userRoles []e.UserRole
	var userRole e.UserRole
	for _, role := range roles {
		err = methods.Database.Where("user_id=? AND role_id=?", userID, role.ID).First(&userRole).Error
		if err != nil && err.Error() != "record not found" {
			return
		} else if err == nil {
			userRoles = append(userRoles, userRole)
		}
		userRole = e.UserRole{}
	}
	err = methods.Database.Delete(&userRoles).Error
	return
}

//func (methods *UserM) DeleteUserRoles(roles []e.Role) (err error) {
//	for _, role := range roles {
//		err = methods.Database.Where("user_roles.role_id IN (?)", role.ID).Delete(&[]e.UserRole{}).Error
//		if err != nil {
//			return
//		}
//	}
//	return
//}
