package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application"
	mock "github.com/intwone/ddd-golang/internal/test/mocks"
	"github.com/stretchr/testify/require"
)

func TestAnswerQuestionUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should create an answer", func(t *testing.T) {
		repo := mock.NewMockRepositoryInterface(ctrl)
		repo.EXPECT().Create(gomock.Any()).AnyTimes()
		useCase := uc.NewDefaultAnswerQuestionUseCase(repo)

		input := uc.AnswerQuestionUseCaseInput{
			InstructorID: "1",
			QuestionID:   "1",
			Content:      "Content",
		}

		answerQuestion, err := useCase.Execute(input)

		require.Nil(t, err)
		require.Equal(t, answerQuestion.GetContent(), input.Content)
	})
}
