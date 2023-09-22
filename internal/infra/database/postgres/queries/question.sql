-- name: GetQuestionBySlug :one
select * from question 
where slug = ?; 

-- name: GetQuestionByID :one
select * from question 
where question_id = ?; 

-- name: GetManyQuestionRecent :many
select * from question 
where created_at >= now() - INTERVAL '3 days' 
order by created_at desc 
limit 10
offset ?; 

-- name: CreateQuestion :exec
insert into question (question_id, author_id, best_answer_id, slug, title, content, is_active, created_at, updated_at) values (?, ?, ?, ?, ?, true, ?, ?);

-- name: SaveQuestion :exec
update question set
  best_answer_id = ?,
  updated_at = ?
where question_id = ?;

-- name: DeleteQuestionByID :exec
update question set
  is_active = false,
  updated_at = ?
where question_id = ?;