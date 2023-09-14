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

func TestGetQuestionBySlugUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should get a question by slug", func(t *testing.T) {
		question := enterprise.NewQuestion("Title Test", "Content test", "1")
		repo := mock.NewMockQuestionRepositoryInterface(ctrl)
		repo.EXPECT().GetBySlug("title-test").Return(question, nil).AnyTimes()
		useCase := uc.NewDefaulGetQuestionBySlugUseCase(repo)

		input := uc.GetQuestionBySlugUseCaseInput{
			Slug: "title-test",
		}

		result, err := useCase.Execute(input)

		require.Nil(t, err)
		require.Equal(t, *question, result)
	})

	t.Run("should not get a question by slug when slug not found", func(t *testing.T) {
		question := enterprise.Question{}
		repo := mock.NewMockQuestionRepositoryInterface(ctrl)
		repo.EXPECT().GetBySlug("title-test").Return(&question, errors.New("any")).AnyTimes()
		useCase := uc.NewDefaulGetQuestionBySlugUseCase(repo)

		input := uc.GetQuestionBySlugUseCaseInput{
			Slug: "title-test",
		}

		result, err := useCase.Execute(input)

		require.NotNil(t, err)
		require.Equal(t, question, result)
	})
}
