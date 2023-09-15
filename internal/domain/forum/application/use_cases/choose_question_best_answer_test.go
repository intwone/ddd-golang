package use_cases_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
	mock "github.com/intwone/ddd-golang/internal/test/mocks"
	"github.com/stretchr/testify/require"
)

func TestChooseQuestionBestAnswerUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should choose the best answer", func(t *testing.T) {
		answer := enterprise.NewAnswer("Content test", "1", "1", "1")
		answersRepo := mock.NewMockAnswerRepositoryInterface(ctrl)
		answersRepo.EXPECT().GetByID(gomock.Any()).Return(*answer, nil).AnyTimes()

		question := enterprise.NewQuestion("Title Test", "Content test", "1", "1")
		questionsRepo := mock.NewMockQuestionRepositoryInterface(ctrl)
		questionsRepo.EXPECT().GetByID(gomock.Any()).Return(*question, nil).AnyTimes()
		questionsRepo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()

		useCase := uc.NewDefaultChooseQuestionBestAnswerUseCase(questionsRepo, answersRepo)

		input := uc.ChooseQuestionBestAnswerUseCaseInput{
			AnswerID: "1",
			AuthorID: "1",
		}

		_, err := useCase.Execute(input)

		require.Nil(t, err)
	})

	t.Run("should not choose the best answer if the answer was from the author himself", func(t *testing.T) {
		answer := enterprise.NewAnswer("Content test", "1", "1", "1")
		answersRepo := mock.NewMockAnswerRepositoryInterface(ctrl)
		answersRepo.EXPECT().GetByID(gomock.Any()).Return(*answer, nil).AnyTimes()

		question := enterprise.NewQuestion("Title Test", "Content test", "2", "1")
		questionsRepo := mock.NewMockQuestionRepositoryInterface(ctrl)
		questionsRepo.EXPECT().GetByID(gomock.Any()).Return(*question, nil).AnyTimes()
		questionsRepo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()

		useCase := uc.NewDefaultChooseQuestionBestAnswerUseCase(questionsRepo, answersRepo)

		input := uc.ChooseQuestionBestAnswerUseCaseInput{
			AnswerID: "1",
			AuthorID: "1",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
		require.Equal(t, "not allowed", err.Error())
	})
}
