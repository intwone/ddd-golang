package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
	s "github.com/intwone/ddd-golang/internal/infra/database/sqlc"
)

type QuestionSQLCRepository struct {
	db *s.Queries
}

func NewQuestionSQLCRepository(db *s.Queries) repositories.QuestionRepositoryInterface {
	return &QuestionSQLCRepository{
		db: db,
	}
}

func (r *QuestionSQLCRepository) GetBySlug(slug string) (*enterprise.Question, error) {
	result, err := r.db.GetQuestionBySlug(context.Background(), slug)

	if err != nil {
		return nil, err
	}

	attachments := enterprise.NewQuestionAttachmentsList([]interface{}{})

	var bestAnswerID string

	if result.BestAnswerID.Valid {
		bestAnswerID = result.BestAnswerID.UUID.String()
	} else {
		bestAnswerID = ""
	}

	return enterprise.NewQuestion(
		result.Title,
		result.Content,
		result.AuthorID.String(),
		enterprise.QuestionOptionalParams{
			BestAnswerID: bestAnswerID,
			Attachments:  *attachments,
		}), nil
}

func (r *QuestionSQLCRepository) GetByID(id string) (*enterprise.Question, error) {
	questionID, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	result, getQuestionByIDErr := r.db.GetQuestionByID(context.Background(), questionID)

	if getQuestionByIDErr != nil {
		return nil, getQuestionByIDErr
	}

	attachments := enterprise.NewQuestionAttachmentsList([]interface{}{})

	return enterprise.NewQuestion(
		result.Title,
		result.Content,
		result.AuthorID.String(),
		enterprise.QuestionOptionalParams{
			BestAnswerID: result.BestAnswerID.UUID.String(),
			Attachments:  *attachments,
		}), nil
}

func (r *QuestionSQLCRepository) GetManyRecent(page int64) (*[]enterprise.Question, error) {
	return nil, nil
}

func (r *QuestionSQLCRepository) Create(question *enterprise.Question) error {
	authorID, err := uuid.Parse(question.GetAuthorID())

	if err != nil {
		return err
	}

	questionID, err := uuid.Parse(question.GetID())

	if err != nil {
		return err
	}

	var bestAnswerID uuid.NullUUID
	if question.GetBestAnswerID() != "" {
		bestAnswerUUID, err := uuid.Parse(question.GetBestAnswerID())

		if err != nil {
			return err
		}

		bestAnswerID = uuid.NullUUID{UUID: bestAnswerUUID, Valid: true}
	}

	createQuestionErr := r.db.CreateQuestion(context.Background(), s.CreateQuestionParams{
		Title:        question.GetTitle(),
		Slug:         question.GetSlug().Value,
		Content:      question.GetContent(),
		QuestionID:   questionID,
		AuthorID:     authorID,
		BestAnswerID: bestAnswerID,
		UpdatedAt:    *question.GetUpdatedAt(),
	})

	if createQuestionErr != nil {
		return createQuestionErr
	}

	return nil
}

func (r *QuestionSQLCRepository) Save(question *enterprise.Question) error {
	return nil
}

func (r *QuestionSQLCRepository) DeleteByID(id string) error {
	questionID, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	deleteQuestionErr := r.db.DeleteQuestionByID(context.Background(), s.DeleteQuestionByIDParams{
		QuestionID: questionID,
		UpdatedAt:  time.Now(),
	})

	if deleteQuestionErr != nil {
		return deleteQuestionErr
	}

	return nil
}
