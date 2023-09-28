package use_cases

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/constants"
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type UpdateAnswerByIDUseCaseInput struct {
	ID             string
	AuthorID       string
	AttachmentsIDs []string
	Content        string
}

type UpdateAnswerByIDUseCaseInterface interface {
	Execute(input UpdateAnswerByIDUseCaseInput) (enterprise.Answer, error)
}

type DefaultUpdateAnswerByIDUseCase struct {
	AnswerRepository            repositories.AnswerRepositoryInterface
	AnswerAttachmentsRepository repositories.AnswerAttachmentsRepositoryInterface
}

func NewDefaultUpdateAnswerByIDUseCase(answerRepository repositories.AnswerRepositoryInterface, answerAttachmentsRepository repositories.AnswerAttachmentsRepositoryInterface) *DefaultUpdateAnswerByIDUseCase {
	return &DefaultUpdateAnswerByIDUseCase{
		AnswerRepository:            answerRepository,
		AnswerAttachmentsRepository: answerAttachmentsRepository,
	}
}

func (uc *DefaultUpdateAnswerByIDUseCase) Execute(input UpdateAnswerByIDUseCaseInput) (*enterprise.Answer, error) {
	answer, err := uc.AnswerRepository.GetByID(input.ID)

	if err != nil {
		return nil, err
	}

	if input.AuthorID != answer.GetAuthorID() {
		return nil, errors.New(constants.NotAllowedError)
	}

	currentAttachments, answerAttachmentErr := uc.AnswerAttachmentsRepository.GetManyByAnswerID(answer.GetID())

	if answerAttachmentErr != nil {
		return nil, err
	}

	attachmentsList := enterprise.NewAnswerAttachmentsList([]interface{}{})

	for _, attachment := range *currentAttachments {
		attachmentsList.Add(attachment)
	}

	newAttachments := make([]interface{}, len(input.AttachmentsIDs))

	for i, attachmentID := range input.AttachmentsIDs {
		newAttachments[i] = enterprise.NewAnswerAttachment(attachmentID, answer.GetID())
	}

	attachmentsList.Update(newAttachments)

	answer.SetContent(input.Content)

	err = uc.AnswerRepository.Save(answer)

	if err != nil {
		return nil, nil
	}

	return answer, nil
}
