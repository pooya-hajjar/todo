package query

const (
	AddNewUser     = "INSERT INTO users (username, password, avatar) VALUES ($1 , $2 , $3) RETURNING id "
	GetUser        = "SELECT id , username , password  FROM users WHERE username = $1 "
	AddTask        = "INSERT INTO tasks (user_id , title , priority , start_time ,end_time ) VALUES ($1 , $2 , $3 , $4 , $5 )"
	GetTasks       = "SELECT title , created_at , updated_at , priority , status , start_time , end_time FROM tasks WHERE user_id = $1 AND deleted_at IS NULL"
	GetTask        = "SELECT title , created_at , updated_at , priority , status , start_time , end_time FROM tasks WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL"
	DeleteTask     = "UPDATE tasks SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL"
	RenameTask     = "UPDATE tasks SET title = $1 , updated_at = CURRENT_TIMESTAMP WHERE id = $2 AND user_id = $3 AND deleted_at IS NULL"
	UpdateTask     = "UPDATE tasks SET status = $1 , updated_at = CURRENT_TIMESTAMP WHERE id = $2 AND user_id = $3 AND deleted_at IS NULL"
	UserInfo       = "SELECT username , email ,status , avatar , (SELECT COUNT(*) FROM tasks WHERE user_id = $1) AS total_tasks  ,  (SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND DATE(created_at) = CURRENT_DATE) AS today_tasks  FROM users WHERE id = $1"
	UpdateUser     = "UPDATE users SET username = $2 , email = $3 ,status = $4, avatar = $5 , updated_at = CURRENT_TIMESTAMP WHERE id = $1"
	GetTopTen      = "SELECT username , COUNT(*) as completed_tasks FROM tasks JOIN users ON users.id = tasks.user_id WHERE tasks.deleted_at IS NULL AND users.deleted_at IS NULL GROUP BY username ORDER BY completed_tasks DESC LIMIT 10"
	CheckUserExist = "SELECT id FROM users WHERE username = $1 "
)
