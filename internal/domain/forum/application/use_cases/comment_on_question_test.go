package use_cases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
	mock "github.com/intwone/ddd-golang/internal/test/mocks"
	"github.com/stretchr/testify/require"
)

func TestCommentOnQuestionUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should create a comment on question", func(t *testing.T) {
		questionCommentRepo := mock.NewMockQuestionCommentsRepositoryInterface(ctrl)
		questionCommentRepo.EXPECT().Create(gomock.Any()).AnyTimes()

		question := enterprise.NewQuestion("Title Test", "Content test", "1")
		questionRepo := mock.NewMockQuestionRepositoryInterface(ctrl)
		questionRepo.EXPECT().GetByID(gomock.Any()).Return(*question, nil).AnyTimes()

		useCase := uc.NewDefaultCommentOnQuestionUseCase(questionRepo, questionCommentRepo)

		input := uc.CommentOnQuestionUseCaseInput{
			AuthorID:   "1",
			QuestionID: "1",
			Content:    "Content Test",
		}

		_, err := useCase.Execute(input)

		require.Nil(t, err)
	})

	t.Run("should not create a comment on question", func(t *testing.T) {
		questionCommentRepo := mock.NewMockQuestionCommentsRepositoryInterface(ctrl)
		questionCommentRepo.EXPECT().Create(gomock.Any()).AnyTimes()

		question := enterprise.Question{}
		questionRepo := mock.NewMockQuestionRepositoryInterface(ctrl)
		questionRepo.EXPECT().GetByID(gomock.Any()).Return(question, errors.New("any")).AnyTimes()

		useCase := uc.NewDefaultCommentOnQuestionUseCase(questionRepo, questionCommentRepo)

		input := uc.CommentOnQuestionUseCaseInput{
			AuthorID:   "1",
			QuestionID: "1",
			Content:    "Content Test",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
	})
}
