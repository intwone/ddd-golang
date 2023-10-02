package mappers

import (
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
	"github.com/intwone/ddd-golang/internal/presentation/dtos"
)

func QuestionDTOMapper(question enterprise.Question) dtos.QuestionDTO {
	return dtos.QuestionDTO{
		ID:           question.GetID(),
		AuthorID:     question.GetAuthorID(),
		BestAnswerID: question.GetBestAnswerID(),
		Slug:         question.GetSlug().Value,
		Title:        question.GetTitle(),
		Content:      question.GetContent(),
		CreatedAt:    *question.GetCreatedAt(),
		UpdatedAt:    question.GetUpdatedAt(),
	}
}

func QuestionsDTOMapper(questions []enterprise.Question) []dtos.QuestionDTO {
	var questionDTOs []dtos.QuestionDTO

	for _, question := range questions {
		questionDTO := dtos.QuestionDTO{
			ID:           question.GetID(),
			AuthorID:     question.GetAuthorID(),
			BestAnswerID: question.GetBestAnswerID(),
			Slug:         question.GetSlug().Value,
			Title:        question.GetTitle(),
			Content:      question.GetContent(),
			CreatedAt:    *question.GetCreatedAt(),
			UpdatedAt:    question.GetUpdatedAt(),
		}

		questionDTOs = append(questionDTOs, questionDTO)
	}

	return questionDTOs
}
