package factories

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

func AnswerCommentsFactory(count int) *[]enterprise.AnswerComment {
	fake := gofakeit.New(0)
	answerComments := make([]enterprise.AnswerComment, count)

	for i := 0; i < count; i++ {
		answerComments[i].SetContent(fake.Paragraph(2, 4, 3, " "))
	}

	return &answerComments
}
