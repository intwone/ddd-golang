package repositories

import (
	"context"

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

	return enterprise.NewQuestion(
		result.Title,
		result.Content,
		result.AuthorID.String(),
		enterprise.QuestionOptionalParams{
			BestAnswerID: result.BestAnswerID.UUID.String(),
			Attachments:  *attachments,
		}), nil
}

func (r *QuestionSQLCRepository) GetByID(id string) (*enterprise.Question, error) {
	return nil, nil
}

func (r *QuestionSQLCRepository) GetManyRecent(page int64) (*[]enterprise.Question, error) {
	return nil, nil
}

func (r *QuestionSQLCRepository) Create(question *enterprise.Question) error {
	return nil
}

func (r *QuestionSQLCRepository) Save(question *enterprise.Question) error {
	return nil
}

func (r *QuestionSQLCRepository) DeleteByID(id string) error {
	return nil
}
