package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	e "td_roles.go/entities"
	m "td_roles.go/methods"
	u "td_roles.go/utils"
)

type TDRole struct {
	RoleMethods          m.RoleMethods
	AuthorizationMethods m.AuthorizationMethods
	UserMethods          m.UserMethods
}

func NewTDRole(database *gorm.DB) (t *TDRole, err error) {
	err = u.MigrationTables(database)
	if err != nil {
		return t, err
	}

	roleM := &m.RoleM{Database: database}
	authorizationM := &m.AuthorizationM{Database: database}
	userM := &m.UserM{Database: database}

	t = &TDRole{
		RoleMethods:          roleM,
		AuthorizationMethods: authorizationM,
		UserMethods:          userM,
	}
	return t, err
}

// Role Function

func (s *TDRole) GetRole(r interface{}, withAuthorization bool) (role e.Role, err error) {
	if u.IsString(r) {
		role, err = s.RoleMethods.GetRoleBySlugName(u.SlugString(r.(string)), withAuthorization)
		return
	} else if u.IsUInt(r) {
		role, err = s.RoleMethods.GetRoleByID(r.(uint), withAuthorization)
	}
	return e.Role{}, errors.New("err unsupported value type")
}

func (s *TDRole) GetRoles(r interface{}, withAuthorization bool) (roles []e.Role, err error) {
	if !u.IsArray(r) {
		var role e.Role
		role, err = s.GetRole(r, withAuthorization)
		if err != nil {
			return []e.Role{}, err
		}
		roles = []e.Role{role}
		return
	} else if u.IsStringArray(r) {
		roles, err = s.RoleMethods.GetRolesBySlugName(u.SlugArray(r.([]string)), withAuthorization)
		return
	} else if u.IsUIntArray(r) {
		roles, err = s.RoleMethods.GetRolesByID(r.([]uint), withAuthorization)
		return
	}
	return []e.Role{}, errors.New("err unsupported value type")
}

func (s *TDRole) CreateRole(name string, description string) error {
	return s.RoleMethods.CreateRole(&e.Role{Name: name, SlugName: u.SlugString(name), Description: description})
}

func (s *TDRole) UpdateRole(r interface{}, newName string, newDescription string) (err error) {
	var role e.Role
	role, err = s.GetRole(r, false)
	if err != nil {
		return
	}
	return s.RoleMethods.UpdateRole(role, newName, newDescription)
}

func (s *TDRole) DeleteRole(r interface{}) (err error) {
	var roles []e.Role
	roles, err = s.GetRoles(r, false)
	if err != nil {
		return
	}
	err = s.RoleMethods.DeleteRoles(roles)
	return
}

func (s *TDRole) AddAuthorizationToRole(r interface{}, a interface{}) (err error) {
	var role e.Role
	var authorizations e.Authorizations
	role, err = s.GetRole(r, false)
	if err != nil {
		return
	}
	authorizations, err = s.GetAuthorizations(a)
	if err != nil {
		return
	}
	if len(authorizations) > 0 {
		err = s.RoleMethods.AddAuthorizations(&role, authorizations)
	}
	return
}

func (s *TDRole) UpdateAuthorizationToRole(r interface{}, a interface{}) (err error) {
	var role e.Role
	var authorizations e.Authorizations
	role, err = s.GetRole(r, false)
	if err != nil {
		return
	}
	authorizations, err = s.GetAuthorizations(a)
	if err != nil {
		return
	}
	if len(authorizations) > 0 {
		err = s.RoleMethods.ReplaceAuthorizations(&role, authorizations)
	}
	return
}

func (s *TDRole) DeleteAuthorizationToRole(r interface{}, a interface{}) (err error) {
	var role e.Role
	var authorizations e.Authorizations
	role, err = s.GetRole(r, false)
	if err != nil {
		return
	}
	authorizations, err = s.GetAuthorizations(a)
	if err != nil {
		return
	}
	if len(authorizations) > 0 {
		err = s.RoleMethods.DelAuthorizations(&role, authorizations)
	}
	return
}

// Authorization Function

func (s *TDRole) GetAuthorization(a interface{}) (authorization e.Authorization, err error) {
	if u.IsString(a) {
		authorization, err = s.AuthorizationMethods.GetAuthorizationBySlugName(u.SlugString(a.(string)))
		return
	} else if u.IsUInt(a) {
		authorization, err = s.AuthorizationMethods.GetAuthorizationByID(a.(uint))
		return
	}
	return e.Authorization{}, errors.New("err unsupported value type")
}

func (s *TDRole) GetAuthorizations(a interface{}) (authorizations e.Authorizations, err error) {
	if !u.IsArray(a) {
		var authorization e.Authorization
		authorization, err = s.GetAuthorization(a)
		if err != nil {
			return e.Authorizations{}, err
		}
		authorizations = e.Authorizations{authorization}
		return
	} else if u.IsStringArray(a) {
		authorizations, err = s.AuthorizationMethods.GetAuthorizationsBySlugName(u.SlugArray(a.([]string)))
		return
	} else if u.IsUIntArray(a) {
		authorizations, err = s.AuthorizationMethods.GetAuthorizationsByID(a.([]uint))
		return
	}
	return e.Authorizations{}, errors.New("err unsupported value type")
}

func (s *TDRole) CreateAuthorization(name string, description string) error {
	return s.AuthorizationMethods.CreateAuthorization(&e.Authorization{Name: name, SlugName: u.SlugString(name), Description: description})
}

func (s *TDRole) UpdateAuthorization(a interface{}, newName string, newDescription string) (err error) {
	var authorization e.Authorization
	authorization, err = s.GetAuthorization(a)
	if err != nil {
		return
	}
	return s.AuthorizationMethods.UpdateAuthorization(authorization, newName, newDescription)
}

// User Function

func (s *TDRole) AddRolesToUser(userID string, r interface{}) (err error) {
	var roles []e.Role
	roles, err = s.GetRoles(r, false)
	if err != nil {
		return err
	}
	if len(roles) > 0 {
		err = s.UserMethods.AddUserRole(userID, roles)
	}
	return
}

func (s *TDRole) DeleteRolesToUser(userID string, r interface{}) (err error) {
	var roles []e.Role
	roles, err = s.GetRoles(r, false)
	if err != nil {
		return err
	}
	if len(roles) > 0 {
		err = s.UserMethods.DeleteUserRole(userID, roles)
	}
	return
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return
	}

	test, tErr := NewTDRole(db)
	if tErr != nil {
		fmt.Println(tErr.Error())
		return
	}

	//err = test.CreateRole("admin des admin", "un super admin")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	//err = test.CreateRole("nul des nuls", "un super nul")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	//test.CreateAuthorization("test", "faire un test")

	//err = test.AddAuthorizationToRole("nul des nuls", "test")
	//fmt.Println(err.Error())

	//test.AddRolesToUser("terry", "admin des admin")
	//test.AddRolesToUser("terry", "nul des nuls")

	//var te []string
	//te = append(te, "admin des admin")
	//te = append(te, "nul des nul")
	//test.DeleteRolesToUser("terry", te)

	err = test.DeleteRole("nul des nuls")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	//var r *e.Role
	//db.Where("id=?", 1).First(&r)

	//db.Create(&e.User{ID: "1"})

	//var u *[]e.User
	//db.Where("users.id IN (?)", "1").Find(&u)
	//db.Where("id=?", "1").First(&u)
	//fmt.Println(u)

	//err = db.Model(r).Association("Users").Append(&U)
	//fmt.Println("err", err.Error())
	//db.Delete(u)

	//totalCount := int64(0)
	//var userIDs []uint
	//
	//var u []e.User
	//db.Find(&u)

	//err = repository.Database.Table("role_permissions").Distinct("role_permissions.permission_id").Where("role_permissions.role_id IN (?)", roleIDs).Count(&totalCount).Scopes(repository.paginate(pagination)).Pluck("role_permissions.permission_id", &permissionIDs).Error
	//err = db.Table("role_users").Distinct("role_users.user_id").Where("role_users.role_id =?", 1).Count(&totalCount).Pluck("role_users.permission_id", &userIDs).Error
	return
}
