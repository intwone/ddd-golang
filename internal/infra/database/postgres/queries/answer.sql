-- name: CreateAnswer :exec
insert into "answers" (answer_id, author_id, question_id, content, created_at, updated_at) values ($1, $2, $3, $4, $5, $6);