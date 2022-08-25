package models

import (
	"database/sql"
	"time"
)

type Activities struct {
	ID        uint         `json:"id"`
	Email     string       `json:"email"`
	Title     string       `json:"title"`
	CreatedAt time.Time    `json:"created_at"`
	UpdateAt  time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type ToDo struct {
	ID        uint         `json:"id"`
	GroupID   uint         `json:"activity_group_id"`
	Title     string       `json:"title"`
	IsActive  bool         `json:"is_active"`
	Priority  string       `json:"priority"`
	CreatedAt time.Time    `json:"created_at"`
	UpdateAt  time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
