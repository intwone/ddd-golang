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

func TestUpdateQuestionByIDUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should update a question", func(t *testing.T) {
		question := enterprise.NewQuestion("Title Test", "Content test", "1")
		repo := mock.NewMockQuestionRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(*question, nil).AnyTimes()
		repo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultUpdateQuestionByIDUseCase(repo)

		input := uc.UpdateQuestionByIDUseCaseInput{
			ID:       "1",
			AuthorID: "1",
			Title:    "Another Title",
			Content:  "Another Content",
		}

		_, err := useCase.Execute(input)

		require.Nil(t, err)
	})

	t.Run("should not update a question when not found question", func(t *testing.T) {
		question := enterprise.Question{}
		repo := mock.NewMockQuestionRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(question, errors.New("any")).AnyTimes()
		repo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultUpdateQuestionByIDUseCase(repo)

		input := uc.UpdateQuestionByIDUseCaseInput{
			ID:       "1",
			AuthorID: "1",
			Title:    "Another Title",
			Content:  "Another Content",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
	})

	t.Run("should not update a question when the author is not the same one who created the question", func(t *testing.T) {
		question := enterprise.NewQuestion("Title Test", "Content test", "1")
		repo := mock.NewMockQuestionRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(*question, nil).AnyTimes()
		useCase := uc.NewDefaultUpdateQuestionByIDUseCase(repo)

		input := uc.UpdateQuestionByIDUseCaseInput{
			ID:       "1",
			AuthorID: "2",
			Title:    "Another Title",
			Content:  "Another Content",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
		require.Equal(t, "not allowed", err.Error())
	})

}
