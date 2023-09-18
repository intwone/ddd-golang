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

func TestGetQuestionCommentsByIDAnswersUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should get many question comments", func(t *testing.T) {
		questionComments := factories.QuestionCommentsFactory(5)
		repo := mock.NewMockQuestionCommentsRepositoryInterface(ctrl)
		repo.EXPECT().GetManyByID(gomock.Any(), gomock.Any()).Return(questionComments, nil).AnyTimes()
		useCase := uc.NewDefaulGetQuestionCommentsByIDUseCase(repo)

		input := uc.GetQuestionCommentsByIDUseCaseInput{
			Page: 1,
			ID:   "1",
		}

		_, err := useCase.Execute(input)

		require.Nil(t, err)
	})

	t.Run("should return empty list when repo returns error", func(t *testing.T) {
		repo := mock.NewMockQuestionCommentsRepositoryInterface(ctrl)
		repo.EXPECT().GetManyByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("any")).AnyTimes()
		useCase := uc.NewDefaulGetQuestionCommentsByIDUseCase(repo)

		input := uc.GetQuestionCommentsByIDUseCaseInput{
			Page: 1,
			ID:   "1",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
	})
}
