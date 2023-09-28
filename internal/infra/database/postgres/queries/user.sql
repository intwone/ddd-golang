-- name: GetUserByEmail :one
select * from "users" where email = $1 limit 1; 

-- name: CreateUser :exec
insert into "users" (user_id, name, email, password, role, updated_at) values ($1, $2, $3, $4, $5, $6);

