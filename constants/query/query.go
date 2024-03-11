package query

const (
	AddNewUser = "INSERT INTO users (username, password, avatar) VALUES ($1 , $2 , $3) RETURNING id "
	GetUser    = "SELECT id , username , password  FROM users WHERE username = $1 "
)
