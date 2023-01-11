package user

import "time"

type User struct {
	ID           int
	Name         string
	Ocuppation   string
	Email        string
	PasswordHash string
	AvatarFile   string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
