package models

type Tasks struct {
	ID        uint
	UserID    uint
	Title     string
	CreatedAt uint
	UpdatedAt uint
	DeletedAt uint
	Priority  int
	Status    int
	StartTime int
	EndTime   int
}
