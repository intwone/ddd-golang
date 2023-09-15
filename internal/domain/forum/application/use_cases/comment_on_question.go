package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type CommentOnQuestionUseCaseInput struct {
	AuthorID   string
	QuestionID string
	Content    string
}

type CommentOnQuestionUseCaseInterface interface {
	Execute(input CommentOnQuestionUseCaseInput) (enterprise.QuestionComment, error)
}

type DefaultCommentOnQuestionUseCase struct {
	QuestionRepository        repositories.QuestionRepositoryInterface
	QuestionCommentRepository repositories.QuestionCommentsRepositoryInterface
}

func NewDefaultCommentOnQuestionUseCase(questionRepository repositories.QuestionRepositoryInterface, questionComentsRepository repositories.QuestionCommentsRepositoryInterface) *DefaultCommentOnQuestionUseCase {
	return &DefaultCommentOnQuestionUseCase{
		QuestionRepository:        questionRepository,
		QuestionCommentRepository: questionComentsRepository,
	}
}

func (uc *DefaultCommentOnQuestionUseCase) Execute(input CommentOnQuestionUseCaseInput) (enterprise.QuestionComment, error) {
	_, err := uc.QuestionRepository.GetByID(input.QuestionID)

	if err != nil {
		return enterprise.QuestionComment{}, err
	}

	questionComment := enterprise.NewQuestionComment(input.Content, input.QuestionID)

	createErr := uc.QuestionCommentRepository.Create(questionComment)

	if createErr != nil {
		return enterprise.QuestionComment{}, createErr
	}

	return *questionComment, nil
}
