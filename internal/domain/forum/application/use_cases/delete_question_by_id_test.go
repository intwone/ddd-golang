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

func TestDeleteQuestionByIDUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should delete a question", func(t *testing.T) {
		question := enterprise.NewQuestion("Title Test", "Content test", "1")
		repo := mock.NewMockQuestionRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(*question, nil).AnyTimes()
		repo.EXPECT().DeleteByID(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultDeleteQuestionByIDUseCase(repo)

		input := uc.DeleteQuestionByIDUseCaseInput{
			ID:       "1",
			AuthorID: "1",
		}

		result := useCase.Execute(input)

		require.Nil(t, result)
	})

	t.Run("should not delete a question when not found question", func(t *testing.T) {
		question := enterprise.Question{}
		repo := mock.NewMockQuestionRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(question, errors.New("any")).AnyTimes()
		repo.EXPECT().DeleteByID(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultDeleteQuestionByIDUseCase(repo)

		input := uc.DeleteQuestionByIDUseCaseInput{
			ID:       "1",
			AuthorID: "1",
		}

		result := useCase.Execute(input)

		require.NotNil(t, result)
	})

	t.Run("should not delete a question when the author is not the same one who created the question", func(t *testing.T) {
		question := enterprise.NewQuestion("Title Test", "Content test", "1")
		repo := mock.NewMockQuestionRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(*question, nil).AnyTimes()
		useCase := uc.NewDefaultDeleteQuestionByIDUseCase(repo)

		input := uc.DeleteQuestionByIDUseCaseInput{
			ID:       "1",
			AuthorID: "2",
		}

		result := useCase.Execute(input)

		require.NotNil(t, result)
		require.Equal(t, "not allowed", result.Error())
	})
}
