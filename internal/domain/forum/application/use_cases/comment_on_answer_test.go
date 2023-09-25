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

func TestCommentOnAnswerUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should create a comment on answer", func(t *testing.T) {
		answerCommentRepo := mock.NewMockAnswerCommentsRepositoryInterface(ctrl)
		answerCommentRepo.EXPECT().Create(gomock.Any()).AnyTimes()

		answer := enterprise.NewAnswer("Content Test", "1", "1")
		answerRepo := mock.NewMockAnswerRepositoryInterface(ctrl)
		answerRepo.EXPECT().GetByID(gomock.Any()).Return(answer, nil).AnyTimes()

		useCase := uc.NewDefaultCommentOnAnswerUseCase(answerRepo, answerCommentRepo)

		input := uc.CommentOnAnswerUseCaseInput{
			AuthorID: "1",
			AnswerID: "1",
			Content:  "Content Test",
		}

		_, err := useCase.Execute(input)

		require.Nil(t, err)
	})

	t.Run("should not create a comment on answer", func(t *testing.T) {
		answerCommentRepo := mock.NewMockAnswerCommentsRepositoryInterface(ctrl)
		answerCommentRepo.EXPECT().Create(gomock.Any()).AnyTimes()

		answerRepo := mock.NewMockAnswerRepositoryInterface(ctrl)
		answerRepo.EXPECT().GetByID(gomock.Any()).Return(nil, errors.New("any")).AnyTimes()

		useCase := uc.NewDefaultCommentOnAnswerUseCase(answerRepo, answerCommentRepo)

		input := uc.CommentOnAnswerUseCaseInput{
			AuthorID: "1",
			AnswerID: "1",
			Content:  "Content Test",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
	})
}
