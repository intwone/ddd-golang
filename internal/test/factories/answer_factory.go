package factories

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

func AnswerFactory(count int) []enterprise.Answer {
	fake := gofakeit.New(0)
	answers := make([]enterprise.Answer, count)

	for i := 0; i < count; i++ {
		answers[i].SetContent(fake.Paragraph(2, 4, 3, " "))
	}

	return answers
}
