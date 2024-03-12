package models

import "database/sql"

type Tasks struct {
	ID        uint          `json:"id"`
	UserID    uint          `json:"user_id"`
	Title     string        `json:"title"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
	DeletedAt sql.NullTime  `json:"deleted_at"`
	Priority  sql.NullInt32 `json:"priority"`
	Status    int           `json:"status"`
	StartTime sql.NullTime  `json:"start_time"`
	EndTime   sql.NullTime  `json:"end_time"`
}
