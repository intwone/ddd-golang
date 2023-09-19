package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type CreateQuestionUseCaseInput struct {
	AuthorID       string
	Title          string
	Content        string
	attachmentsIDs []string
}

type CreateQuestionUseCaseInterface interface {
	Execute(input CreateQuestionUseCaseInput) (enterprise.Question, error)
}

type DefaultCreateQuestionUseCase struct {
	QuestionRepository repositories.QuestionRepositoryInterface
}

func NewDefaultCreateQuestionUseCase(questionRepository repositories.QuestionRepositoryInterface) *DefaultCreateQuestionUseCase {
	return &DefaultCreateQuestionUseCase{
		QuestionRepository: questionRepository,
	}
}

func (uc *DefaultCreateQuestionUseCase) Execute(input CreateQuestionUseCaseInput) (enterprise.Question, error) {
	newQuestion := enterprise.NewQuestion(input.Title, input.Content, input.AuthorID)

	var attachments []enterprise.QuestionAttachment

	for _, attachmentID := range input.attachmentsIDs {
		attachment := enterprise.NewQuestionAttachment(attachmentID, newQuestion.GetID())
		attachments = append(attachments, *attachment)
	}

	newQuestion.SetAttachments(attachments)

	err := uc.QuestionRepository.Create(newQuestion)

	if err != nil {
		return enterprise.Question{}, err
	}

	return *newQuestion, nil
}
