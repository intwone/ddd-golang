package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

type AnswerQuestionUseCaseInput struct {
	InstructorID   string
	QuestionID     string
	AttachmentsIDs []string
	Content        string
}

type AnswerQuestionUseCaseInterface interface {
	Execute(input AnswerQuestionUseCaseInput) (enterprise.Answer, error)
}

type DefaultAnswerQuestionUseCase struct {
	AnswersRepository repositories.AnswerRepositoryInterface
}

func NewDefaultAnswerQuestionUseCase(answersRepository repositories.AnswerRepositoryInterface) *DefaultAnswerQuestionUseCase {
	return &DefaultAnswerQuestionUseCase{
		AnswersRepository: answersRepository,
	}
}

func (uc *DefaultAnswerQuestionUseCase) Execute(input AnswerQuestionUseCaseInput) (enterprise.Answer, error) {
	newAnswer := enterprise.NewAnswer(input.Content, input.InstructorID, input.QuestionID)

	attachments := make([]*enterprise.AnswerAttachment, len(input.AttachmentsIDs))

	for i, attachmentID := range input.AttachmentsIDs {
		attachments[i] = enterprise.NewAnswerAttachment(attachmentID, newAnswer.GetID())
	}

	attachmentsList := enterprise.NewAnswerAttachmentsList([]interface{}{})

	for _, attachment := range attachments {
		attachmentsList.Add(attachment)
	}

	newAnswer.SetAttachments(*attachmentsList)

	uc.AnswersRepository.Create(newAnswer)

	return *newAnswer, nil
}
