package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
	s "github.com/intwone/ddd-golang/internal/infra/database/sqlc"
)

type AnswerSQLCRepository struct {
	db *s.Queries
}

func NewAnswerSQLCRepository(db *s.Queries) repositories.AnswerRepositoryInterface {
	return &AnswerSQLCRepository{
		db: db,
	}
}

func (r *AnswerSQLCRepository) GetByID(id string) (*enterprise.Answer, error) {
	return nil, nil
}

func (r *AnswerSQLCRepository) GetManyByQuestionID(page int64, questionID string) (*[]enterprise.Answer, error) {
	return nil, nil
}

func (r *AnswerSQLCRepository) Create(answer *enterprise.Answer) error {
	authorID, err := uuid.Parse(answer.GetAuthorID())

	if err != nil {
		return err
	}

	answerID, err := uuid.Parse(answer.GetID())

	if err != nil {
		return err
	}

	questionID, err := uuid.Parse(answer.GetQuestionID())

	if err != nil {
		return err
	}

	createAnswerErr := r.db.CreateAnswer(context.Background(), s.CreateAnswerParams{
		AuthorID:   authorID,
		AnswerID:   answerID,
		QuestionID: questionID,
		Content:    answer.GetContent(),
		CreatedAt:  answer.GetCreatedAt(),
		UpdatedAt:  *answer.GetUpdatedAt(),
	})

	if createAnswerErr != nil {
		return createAnswerErr
	}

	return nil
}

func (r *AnswerSQLCRepository) Save(answer *enterprise.Answer) error {
	return nil
}

func (r *AnswerSQLCRepository) DeleteByID(id string) error {
	return nil
}
