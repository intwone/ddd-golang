package use_cases_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	mock "github.com/intwone/ddd-golang/internal/test/mocks"
	"github.com/stretchr/testify/require"
)

func TestCreateQuestionUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should create a question", func(t *testing.T) {
		repo := mock.NewMockQuestionRepositoryInterface(ctrl)
		repo.EXPECT().Create(gomock.Any()).AnyTimes()
		useCase := uc.NewDefaultCreateQuestionUseCase(repo)

		input := uc.CreateQuestionUseCaseInput{
			Title:    "Title Example",
			Content:  "Content",
			AuthorID: "1",
		}

		createQuestion, err := useCase.Execute(input)

		require.Nil(t, err)
		require.NotNil(t, createQuestion.GetAuthorID().Value)
		require.Equal(t, createQuestion.GetTitle(), input.Title)
		require.Equal(t, createQuestion.GetContent(), input.Content)
		require.Equal(t, createQuestion.GetSlug().Value, "title-example")
	})
}
