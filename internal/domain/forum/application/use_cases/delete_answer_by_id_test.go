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

func TestDeleteAnswerByIDUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should delete an answer", func(t *testing.T) {
		answer := enterprise.NewAnswer("Content test", "1", "1")
		repo := mock.NewMockAnswerRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(*answer, nil).AnyTimes()
		repo.EXPECT().DeleteByID(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultDeleteAnswerByIDUseCase(repo)

		input := uc.DeleteAnswerByIDUseCaseInput{
			ID:       "1",
			AuthorID: "1",
		}

		result := useCase.Execute(input)

		require.Nil(t, result)
	})

	t.Run("should not delete an answer when not found answer", func(t *testing.T) {
		answer := enterprise.Answer{}
		repo := mock.NewMockAnswerRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(answer, errors.New("any")).AnyTimes()
		repo.EXPECT().DeleteByID(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultDeleteAnswerByIDUseCase(repo)

		input := uc.DeleteAnswerByIDUseCaseInput{
			ID:       "1",
			AuthorID: "1",
		}

		result := useCase.Execute(input)

		require.NotNil(t, result)
	})

	t.Run("should not delete an answer when the author is not the same one who created the answer", func(t *testing.T) {
		answer := enterprise.NewAnswer("Content Test", "1", "1")
		repo := mock.NewMockAnswerRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(*answer, nil).AnyTimes()
		useCase := uc.NewDefaultDeleteAnswerByIDUseCase(repo)

		input := uc.DeleteAnswerByIDUseCaseInput{
			ID:       "1",
			AuthorID: "2",
		}

		result := useCase.Execute(input)

		require.NotNil(t, result)
		require.Equal(t, "not allowed", result.Error())
	})
}
