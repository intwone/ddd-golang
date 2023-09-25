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

func TestDeleteQuestionCommentByIDUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should delete a question comment", func(t *testing.T) {
		questionComment := enterprise.NewQuestionComment("Content test", "1", "1")
		repo := mock.NewMockQuestionCommentsRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(questionComment, nil).AnyTimes()
		repo.EXPECT().DeleteByID(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultDeleteQuestionCommentByIDUseCase(repo)

		input := uc.DeleteQuestionCommentByIDUseCaseInput{
			ID:       "1",
			AuthorID: "1",
		}

		result := useCase.Execute(input)

		require.Nil(t, result)
	})

	t.Run("should not delete a question comment when not found questionComment", func(t *testing.T) {
		repo := mock.NewMockQuestionCommentsRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(nil, errors.New("any")).AnyTimes()
		repo.EXPECT().DeleteByID(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultDeleteQuestionCommentByIDUseCase(repo)

		input := uc.DeleteQuestionCommentByIDUseCaseInput{
			ID:       "1",
			AuthorID: "1",
		}

		result := useCase.Execute(input)

		require.NotNil(t, result)
	})

	t.Run("should not delete a questionComment when the author is not the same one who created the questionComment", func(t *testing.T) {
		questionComment := enterprise.NewQuestionComment("Content test", "1", "1")
		repo := mock.NewMockQuestionCommentsRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(questionComment, nil).AnyTimes()
		useCase := uc.NewDefaultDeleteQuestionCommentByIDUseCase(repo)

		input := uc.DeleteQuestionCommentByIDUseCaseInput{
			ID:       "1",
			AuthorID: "2",
		}

		result := useCase.Execute(input)

		require.NotNil(t, result)
		require.Equal(t, "not allowed", result.Error())
	})
}
