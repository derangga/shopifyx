package entity

import "time"

type User struct {
	ID          int
	Username    string
	Name        string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	AccessToken string
}
