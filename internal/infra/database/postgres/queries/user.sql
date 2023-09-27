-- name: CreateUser :exec
insert into "users" (user_id, name, role) values ($1, $2, $3);

