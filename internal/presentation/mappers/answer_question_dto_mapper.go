package mappers

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
	"github.com/intwone/ddd-golang/internal/presentation/dtos"
)

func AnswerDTOMapper(answer enterprise.Answer) dtos.AnswerResponseDTO {
	return dtos.AnswerResponseDTO{
		ID:         answer.GetID(),
		AuthorID:   answer.GetAuthorID(),
		QuestionID: answer.GetQuestionID(),
		Content:    answer.GetContent(),
		CreatedAt:  answer.GetCreatedAt(),
		UpdatedAt:  answer.GetUpdatedAt(),
	}
}
