-- name: GetQuestionBySlug :one
select * from "questions" where slug = $1 limit 1; 

-- name: GetQuestionByID :one
select * from "questions" where question_id = $1 limit 1; 

-- name: GetManyQuestionRecent :many
select * from "questions"
where created_at >= now() - INTERVAL '3 days' 
order by created_at desc 
limit 10
offset $1; 

-- name: CreateQuestion :exec
insert into "questions" (question_id, author_id, best_answer_id, slug, title, content, is_active, updated_at) values ($1, $2, $3, $4, $5, $6, true, $7);

-- name: SaveQuestion :exec
update "questions" set
  best_answer_id = $1,
  updated_at = $2
where question_id = $3;

-- name: DeleteQuestionByID :exec
update "questions" set
  is_active = false,
  updated_at = $1
where question_id = $2;