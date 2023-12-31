// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: question.sql

package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createQuestion = `-- name: CreateQuestion :exec
insert into "questions" (question_id, author_id, best_answer_id, slug, title, content, is_active, updated_at) values ($1, $2, $3, $4, $5, $6, true, $7)
`

type CreateQuestionParams struct {
	QuestionID   uuid.UUID
	AuthorID     uuid.UUID
	BestAnswerID uuid.NullUUID
	Slug         string
	Title        string
	Content      string
	UpdatedAt    time.Time
}

func (q *Queries) CreateQuestion(ctx context.Context, arg CreateQuestionParams) error {
	_, err := q.db.ExecContext(ctx, createQuestion,
		arg.QuestionID,
		arg.AuthorID,
		arg.BestAnswerID,
		arg.Slug,
		arg.Title,
		arg.Content,
		arg.UpdatedAt,
	)
	return err
}

const deleteQuestionByID = `-- name: DeleteQuestionByID :exec
update "questions" set
  is_active = false,
  updated_at = $1
where question_id = $2
`

type DeleteQuestionByIDParams struct {
	UpdatedAt  time.Time
	QuestionID uuid.UUID
}

func (q *Queries) DeleteQuestionByID(ctx context.Context, arg DeleteQuestionByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteQuestionByID, arg.UpdatedAt, arg.QuestionID)
	return err
}

const getManyQuestionRecent = `-- name: GetManyQuestionRecent :many
select question_id, author_id, best_answer_id, slug, title, content, is_active, created_at, updated_at from "questions"
where created_at >= now() - INTERVAL '1 days' 
order by created_at desc 
limit 20
offset $1
`

func (q *Queries) GetManyQuestionRecent(ctx context.Context, offset int32) ([]Question, error) {
	rows, err := q.db.QueryContext(ctx, getManyQuestionRecent, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Question
	for rows.Next() {
		var i Question
		if err := rows.Scan(
			&i.QuestionID,
			&i.AuthorID,
			&i.BestAnswerID,
			&i.Slug,
			&i.Title,
			&i.Content,
			&i.IsActive,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getQuestionByID = `-- name: GetQuestionByID :one
select question_id, author_id, best_answer_id, slug, title, content, is_active, created_at, updated_at from "questions" where question_id = $1 limit 1
`

func (q *Queries) GetQuestionByID(ctx context.Context, questionID uuid.UUID) (Question, error) {
	row := q.db.QueryRowContext(ctx, getQuestionByID, questionID)
	var i Question
	err := row.Scan(
		&i.QuestionID,
		&i.AuthorID,
		&i.BestAnswerID,
		&i.Slug,
		&i.Title,
		&i.Content,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getQuestionBySlug = `-- name: GetQuestionBySlug :one
select question_id, author_id, best_answer_id, slug, title, content, is_active, created_at, updated_at from "questions" where slug = $1 limit 1
`

func (q *Queries) GetQuestionBySlug(ctx context.Context, slug string) (Question, error) {
	row := q.db.QueryRowContext(ctx, getQuestionBySlug, slug)
	var i Question
	err := row.Scan(
		&i.QuestionID,
		&i.AuthorID,
		&i.BestAnswerID,
		&i.Slug,
		&i.Title,
		&i.Content,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const saveQuestion = `-- name: SaveQuestion :exec
update "questions" set
  best_answer_id = $1,
  updated_at = $2
where question_id = $3
`

type SaveQuestionParams struct {
	BestAnswerID uuid.NullUUID
	UpdatedAt    time.Time
	QuestionID   uuid.UUID
}

func (q *Queries) SaveQuestion(ctx context.Context, arg SaveQuestionParams) error {
	_, err := q.db.ExecContext(ctx, saveQuestion, arg.BestAnswerID, arg.UpdatedAt, arg.QuestionID)
	return err
}
