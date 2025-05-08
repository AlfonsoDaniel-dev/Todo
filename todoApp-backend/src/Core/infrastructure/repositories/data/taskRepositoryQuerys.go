package data

const saveTaskOnDB = "INSERT INTO tasks (id, owner_id, title, body, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
