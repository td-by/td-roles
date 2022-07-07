package utils

import (
	"gorm.io/gorm"
	e "td_roles.go/entities"
)

func MigrationTables(database *gorm.DB) error {
	err := database.AutoMigrate(&e.Role{})
	if err != nil {
		return err
	}

	err = database.AutoMigrate(&e.Authorization{})
	if err != nil {
		return err
	}

	err = database.AutoMigrate(&e.UserRole{})
	if err != nil {
		return err
	}

	err = database.AutoMigrate(&e.UserAuthorization{})
	if err != nil {
		return err
	}

	return nil
}
