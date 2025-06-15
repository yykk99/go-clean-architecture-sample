package model

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreateAt  time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
