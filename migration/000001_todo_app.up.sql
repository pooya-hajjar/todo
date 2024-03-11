CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(120) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    status INTEGER CHECK(status BETWEEN -1 AND 1) DEFAULT -1,
    email VARCHAR(255),
    avatar VARCHAR(255)
);

CREATE TABLE tasks(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    title VARCHAR(50) NOT NULL,
    priority INTEGER DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    status INTEGER CHECK(status BETWEEN -2 AND 1) DEFAULT -1,
    start_time TIMESTAMP WITH TIME ZONE ,
    end_time TIMESTAMP WITH TIME ZONE ,
    CONSTRAINT check_times CHECK((start_time IS NOT NULL AND end_time IS NOT NULL) OR (start_time IS NULL AND end_time IS NULL))
);