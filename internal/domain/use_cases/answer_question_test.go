package usecases

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnswerQuestionUseCase_Execute(t *testing.T) {
	t.Run("create an answer", func(t *testing.T) {
		input := AnswerQuestionUseCaseInput{
			InstructorID: "1",
			QuestionID:   "1",
			Content:      "Content",
		}

		useCase := &DefaultAnswerQuestionUseCase{}

		answerQuestion, _ := useCase.Execute(input)

		require.Equal(t, answerQuestion.Content, input.Content)
	})
}
