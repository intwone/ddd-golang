package factories

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

func QuestionFactory(count int) []enterprise.Question {
	fake := gofakeit.New(0)
	questions := make([]enterprise.Question, count)

	for i := 0; i < count; i++ {
		questions[i].SetTitle(fake.Sentence(3))
		questions[i].SetContent(fake.Paragraph(2, 4, 3, " "))
	}

	return questions
}
