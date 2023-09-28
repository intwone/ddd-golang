package use_cases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/intwone/ddd-golang/internal/constants"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
	mock "github.com/intwone/ddd-golang/internal/test/mocks"
	"github.com/stretchr/testify/require"
)

func TestDeleteAnswerCommentByIDUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should delete an answer comment", func(t *testing.T) {
		answerComment := enterprise.NewAnswerComment("Content test", "1", "1")
		repo := mock.NewMockAnswerCommentsRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(answerComment, nil).AnyTimes()
		repo.EXPECT().DeleteByID(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultDeleteAnswerCommentByIDUseCase(repo)

		input := uc.DeleteAnswerCommentByIDUseCaseInput{
			ID:       "1",
			AuthorID: "1",
		}

		result := useCase.Execute(input)

		require.Nil(t, result)
	})

	t.Run("should not delete a question comment when not found answer comment", func(t *testing.T) {
		repo := mock.NewMockAnswerCommentsRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(nil, errors.New("any")).AnyTimes()
		repo.EXPECT().DeleteByID(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultDeleteAnswerCommentByIDUseCase(repo)

		input := uc.DeleteAnswerCommentByIDUseCaseInput{
			ID:       "1",
			AuthorID: "1",
		}

		result := useCase.Execute(input)

		require.NotNil(t, result)
	})

	t.Run("should not delete a answer comment when the author is not the same one who created the answer comment", func(t *testing.T) {
		answerComment := enterprise.NewAnswerComment("Content test", "1", "1")
		repo := mock.NewMockAnswerCommentsRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(answerComment, nil).AnyTimes()
		useCase := uc.NewDefaultDeleteAnswerCommentByIDUseCase(repo)

		input := uc.DeleteAnswerCommentByIDUseCaseInput{
			ID:       "1",
			AuthorID: "2",
		}

		result := useCase.Execute(input)

		require.NotNil(t, result)
		require.Equal(t, constants.NotAllowedError, result.Error())
	})
}
