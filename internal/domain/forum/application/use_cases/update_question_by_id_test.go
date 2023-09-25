package use_cases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
	"github.com/intwone/ddd-golang/internal/test/factories"
	mock "github.com/intwone/ddd-golang/internal/test/mocks"
	"github.com/stretchr/testify/require"
)

func TestUpdateQuestionByIDUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should update a question", func(t *testing.T) {
		attachments := enterprise.NewQuestionAttachmentsList([]interface{}{"1", "2", "3"})
		question := enterprise.NewQuestion("Title Test", "Content test", "1", enterprise.QuestionOptionalParams{ID: "1", Attachments: *attachments})
		questionRepo := mock.NewMockQuestionRepositoryInterface(ctrl)
		questionRepo.EXPECT().GetByID(gomock.Any()).Return(question, nil).AnyTimes()
		questionRepo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()

		questionAttachments := factories.QuestionAttachmentsFactory(3, "1")
		questionAttachmentsRepo := mock.NewMockQuestionAttachmentsRepositoryInterface(ctrl)
		questionAttachmentsRepo.EXPECT().GetManyByQuestionID(gomock.Any()).Return(questionAttachments, nil).AnyTimes()

		useCase := uc.NewDefaultUpdateQuestionByIDUseCase(questionRepo, questionAttachmentsRepo)

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
		questionRepo := mock.NewMockQuestionRepositoryInterface(ctrl)
		questionRepo.EXPECT().GetByID(gomock.Any()).Return(nil, errors.New("any")).AnyTimes()
		questionRepo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()

		questionAttachments := factories.QuestionAttachmentsFactory(5, "1")
		questionAttachmentsRepo := mock.NewMockQuestionAttachmentsRepositoryInterface(ctrl)
		questionAttachmentsRepo.EXPECT().GetManyByQuestionID(gomock.Any()).Return(questionAttachments, nil).AnyTimes()

		useCase := uc.NewDefaultUpdateQuestionByIDUseCase(questionRepo, questionAttachmentsRepo)

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
		questionRepo := mock.NewMockQuestionRepositoryInterface(ctrl)
		questionRepo.EXPECT().GetByID(gomock.Any()).Return(question, nil).AnyTimes()
		questionRepo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()

		questionAttachments := factories.QuestionAttachmentsFactory(5, "1")
		questionAttachmentsRepo := mock.NewMockQuestionAttachmentsRepositoryInterface(ctrl)
		questionAttachmentsRepo.EXPECT().GetManyByQuestionID(gomock.Any()).Return(questionAttachments, nil).AnyTimes()

		useCase := uc.NewDefaultUpdateQuestionByIDUseCase(questionRepo, questionAttachmentsRepo)

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
