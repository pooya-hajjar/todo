package models

import "database/sql"

type Users struct {
	ID        int            `json:"id"`
	UserName  string         `json:"username"`
	Password  string         `json:"password"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	DeletedAt sql.NullTime   `json:"deleted_at"`
	Status    int            `json:"status"`
	Email     sql.NullString `json:"email"`
	Avatar    sql.NullString `json:"avatar"`
}
