package use_cases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/test/factories"
	mock "github.com/intwone/ddd-golang/internal/test/mocks"
	"github.com/stretchr/testify/require"
)

func TestGetRecentQuestionsUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should get many questions", func(t *testing.T) {
		questions := factories.QuestionFactory(5)
		repo := mock.NewMockQuestionRepositoryInterface(ctrl)
		repo.EXPECT().GetManyRecent(gomock.Any()).Return(questions, nil).AnyTimes()
		useCase := uc.NewDefaulGetRecentQuestionsUseCase(repo)

		input := uc.GetRecentQuestionsUseCaseInput{
			Page: 1,
		}

		_, err := useCase.Execute(input)

		require.Nil(t, err)
	})

	t.Run("should return empty list when repo returns error", func(t *testing.T) {
		repo := mock.NewMockQuestionRepositoryInterface(ctrl)
		repo.EXPECT().GetManyRecent(gomock.Any()).Return(nil, errors.New("any")).AnyTimes()
		useCase := uc.NewDefaulGetRecentQuestionsUseCase(repo)

		input := uc.GetRecentQuestionsUseCaseInput{
			Page: 1,
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
	})
}
