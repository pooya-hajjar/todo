CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(120) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    status INTEGER CHECK(status BETWEEN -1 AND 1) DEFAULT -1,
    email VARCHAR(255) UNIQUE,
    avatar VARCHAR(255)
);

CREATE TABLE tasks(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    title VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    priority INTEGER,
    status INTEGER CHECK(status BETWEEN -2 AND 1) DEFAULT -1,
    start_time TIMESTAMP WITH TIME ZONE ,
    end_time TIMESTAMP WITH TIME ZONE ,
    CONSTRAINT check_times CHECK((start_time IS NOT NULL AND end_time IS NOT NULL) OR (start_time IS NULL AND end_time IS NULL))
);


CREATE INDEX tasks_user_id_idx ON tasks (user_id);

CREATE OR REPLACE FUNCTION get_task_count(user_id INT)
RETURNS INT AS $$
DECLARE
task_count INT;
BEGIN
SELECT COUNT(*)
INTO task_count
FROM tasks
WHERE tasks.user_id = get_task_count.user_id AND tasks.deleted_at IS NULL;

RETURN task_count;
END;
$$ LANGUAGE plpgsql;