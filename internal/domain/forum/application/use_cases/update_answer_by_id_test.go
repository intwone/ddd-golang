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

func TestUpdateAnswerByIDUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should update an answer", func(t *testing.T) {
		answer := enterprise.NewAnswer("Content test", "1", "1")
		repo := mock.NewMockAnswerRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(*answer, nil).AnyTimes()
		repo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultUpdateAnswerByIDUseCase(repo)

		input := uc.UpdateAnswerByIDUseCaseInput{
			ID:       "1",
			AuthorID: "1",
			Content:  "Another Content",
		}

		_, err := useCase.Execute(input)

		require.Nil(t, err)
	})

	t.Run("should not update an answer when not found answer", func(t *testing.T) {
		answer := enterprise.Answer{}
		repo := mock.NewMockAnswerRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(answer, errors.New("any")).AnyTimes()
		repo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultUpdateAnswerByIDUseCase(repo)

		input := uc.UpdateAnswerByIDUseCaseInput{
			ID:       "1",
			AuthorID: "1",
			Content:  "Another Content",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
	})

	t.Run("should not update an answer when the author is not the same one who created the answer", func(t *testing.T) {
		answer := enterprise.NewAnswer("Content test", "1", "1")
		repo := mock.NewMockAnswerRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(*answer, nil).AnyTimes()
		useCase := uc.NewDefaultUpdateAnswerByIDUseCase(repo)

		input := uc.UpdateAnswerByIDUseCaseInput{
			ID:       "1",
			AuthorID: "2",
			Content:  "Another Content",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
		require.Equal(t, "not allowed", err.Error())
	})
}
