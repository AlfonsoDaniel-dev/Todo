package repositories

const saveUserOnDB = "INSERT INTO users (id, user_name, email, password, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING created_at"

const getUserOnDB = "SELECT (id, user_name, email, password, created_at, updated_at) FROM users WHERE id=$1"

const getUserIdByEmailOnDB = "SELECT id FROM users WHERE email=$1"

const getUserEmailByIdOnDB = "SELECT email FROM users WHERE id=$1"

const userGetItsTasks = "SELECT id, owner_id, title, body, created_at, updated_at FROM tasks WHERE owner_id=$1"
