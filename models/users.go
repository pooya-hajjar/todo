package models

type Users struct {
	ID        uint
	UserName  string
	Password  string
	CreatedAt uint
	UpdatedAt uint
	DeletedAt uint
	Status    int
	Email     string
	Avatar    string
}
