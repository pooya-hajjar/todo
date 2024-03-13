package query

const (
	AddNewUser = "INSERT INTO users (username, password, avatar) VALUES ($1 , $2 , $3) RETURNING id "
	GetUser    = "SELECT id , username , password  FROM users WHERE username = $1 "
	AddTask    = "INSERT INTO tasks (user_id , title , priority , start_time ,end_time ) VALUES ($1 , $2 , $3 , $4 , $5 )"
	GetTasks   = "SELECT title , created_at , updated_at , priority , status , start_time , end_time FROM tasks WHERE user_id = $1"
	GetTask    = "SELECT title , created_at , updated_at , priority , status , start_time , end_time FROM tasks WHERE id = $1 AND user_id = $2"
)
