package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type CreateQuestionUseCaseInput struct {
	AuthorID       string
	AttachmentsIDs []string
	Title          string
	Content        string
}

type CreateQuestionUseCaseInterface interface {
	Execute(input CreateQuestionUseCaseInput) (*enterprise.Question, error)
}

type DefaultCreateQuestionUseCase struct {
	QuestionRepository repositories.QuestionRepositoryInterface
}

func NewDefaultCreateQuestionUseCase(questionRepository repositories.QuestionRepositoryInterface) *DefaultCreateQuestionUseCase {
	return &DefaultCreateQuestionUseCase{
		QuestionRepository: questionRepository,
	}
}

func (uc *DefaultCreateQuestionUseCase) Execute(input CreateQuestionUseCaseInput) (*enterprise.Question, error) {
	newQuestion := enterprise.NewQuestion(input.Title, input.Content, input.AuthorID)

	attachments := make([]*enterprise.QuestionAttachment, len(input.AttachmentsIDs))

	for i, attachmentID := range input.AttachmentsIDs {
		attachments[i] = enterprise.NewQuestionAttachment(attachmentID, newQuestion.GetID())
	}

	attachmentsList := enterprise.NewQuestionAttachmentsList([]interface{}{})

	for _, attachment := range attachments {
		attachmentsList.Add(attachment)
	}

	newQuestion.SetAttachments(*attachmentsList)

	err := uc.QuestionRepository.Create(newQuestion)

	if err != nil {
		return nil, err
	}

	return newQuestion, nil
}
