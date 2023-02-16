package user

import "time"

type User struct {
	ID           int
	Name         string
	Ocuppation   string
	Email        string
	PasswordHash string
	Role         string
	Token        string
	AvatarFile   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
