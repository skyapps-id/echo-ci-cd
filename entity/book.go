package entity

import "time"

type Book struct {
	UUID        string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
