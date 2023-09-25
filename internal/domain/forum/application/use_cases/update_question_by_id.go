package use_cases

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type UpdateQuestionByIDUseCaseInput struct {
	ID             string
	AuthorID       string
	AttachmentsIDs []string
	Title          string
	Content        string
}

type UpdateQuestionByIDUseCaseInterface interface {
	Execute(input UpdateQuestionByIDUseCaseInput) (enterprise.Question, error)
}

type DefaultUpdateQuestionByIDUseCase struct {
	QuestionRepository            repositories.QuestionRepositoryInterface
	QuestionAttachmentsRepository repositories.QuestionAttachmentsRepositoryInterface
}

func NewDefaultUpdateQuestionByIDUseCase(questionRepository repositories.QuestionRepositoryInterface, questionAttachmentsRepository repositories.QuestionAttachmentsRepositoryInterface) *DefaultUpdateQuestionByIDUseCase {
	return &DefaultUpdateQuestionByIDUseCase{
		QuestionRepository:            questionRepository,
		QuestionAttachmentsRepository: questionAttachmentsRepository,
	}
}

func (uc *DefaultUpdateQuestionByIDUseCase) Execute(input UpdateQuestionByIDUseCaseInput) (*enterprise.Question, error) {
	question, err := uc.QuestionRepository.GetByID(input.ID)

	if err != nil {
		return nil, err
	}

	if input.AuthorID != question.GetAuthorID() {
		return nil, errors.New("not allowed")
	}

	currentAttachments, questionAttachmentErr := uc.QuestionAttachmentsRepository.GetManyByQuestionID(question.GetID())

	if questionAttachmentErr != nil {
		return nil, err
	}

	attachmentsList := enterprise.NewQuestionAttachmentsList([]interface{}{})

	for _, attachment := range *currentAttachments {
		attachmentsList.Add(attachment)
	}

	newAttachments := make([]interface{}, len(input.AttachmentsIDs))

	for i, attachmentID := range input.AttachmentsIDs {
		newAttachments[i] = enterprise.NewQuestionAttachment(attachmentID, question.GetID())
	}

	attachmentsList.Update(newAttachments)

	question.SetTitle(input.Title)
	question.SetContent(input.Content)
	question.SetAttachments(*attachmentsList)

	err = uc.QuestionRepository.Save(question)

	if err != nil {
		return nil, err
	}

	return question, nil
}
