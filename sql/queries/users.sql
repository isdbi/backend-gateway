-- db/queries/users.sql

-- Create a new user
-- name: CreateUser :exec
INSERT INTO users (id, email, password, name, family_name, age, sex) 
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- Get a user by ID
-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- Get a user by email
-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1;

-- Get all users
-- name: ListUsers :many
SELECT * FROM users;

-- Update user information
-- name: UpdateUser :exec
UPDATE users
SET email = $2, password = $3, name = $4, age = $5, family_name = $6, sex = $7
WHERE id = $1;

-- Activate user
-- name: ActivateUser :exec
UPDATE users
SET is_active = TRUE
WHERE id = $1;

-- Delete a user
-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;

