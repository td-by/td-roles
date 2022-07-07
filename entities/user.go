package entities

type UserRole struct {
	UserID string `gorm:"primary_key"`
	RoleID uint   `gorm:"primary_key"`
}

type UserAuthorization struct {
	UserID          string `gorm:"primary_key"`
	AuthorizationID uint   `gorm:"primary_key"`
}
