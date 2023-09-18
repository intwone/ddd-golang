package factories

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

func QuestionCommentsFactory(count int) []enterprise.QuestionComment {
	fake := gofakeit.New(0)
	questionComments := make([]enterprise.QuestionComment, count)

	for i := 0; i < count; i++ {
		questionComments[i].SetContent(fake.Paragraph(2, 4, 3, " "))
	}

	return questionComments
}
