package use_cases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/intwone/ddd-golang/internal/constants"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
	"github.com/intwone/ddd-golang/internal/test/factories"
	mock "github.com/intwone/ddd-golang/internal/test/mocks"
	"github.com/stretchr/testify/require"
)

func TestUpdateAnswerByIDUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should update an answer", func(t *testing.T) {
		answer := enterprise.NewAnswer("Content test", "1", "1")
		answerRepo := mock.NewMockAnswerRepositoryInterface(ctrl)
		answerRepo.EXPECT().GetByID(gomock.Any()).Return(answer, nil).AnyTimes()
		answerRepo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()

		answerAttachments := factories.AnswerAttachmentsFactory(3, "1")
		answerAttachmentsRepo := mock.NewMockAnswerAttachmentsRepositoryInterface(ctrl)
		answerAttachmentsRepo.EXPECT().GetManyByAnswerID(gomock.Any()).Return(answerAttachments, nil).AnyTimes()

		useCase := uc.NewDefaultUpdateAnswerByIDUseCase(answerRepo, answerAttachmentsRepo)

		input := uc.UpdateAnswerByIDUseCaseInput{
			ID:       "1",
			AuthorID: "1",
			Content:  "Another Content",
		}

		_, err := useCase.Execute(input)

		require.Nil(t, err)
	})

	t.Run("should not update an answer when not found answer", func(t *testing.T) {
		answerRepo := mock.NewMockAnswerRepositoryInterface(ctrl)
		answerRepo.EXPECT().GetByID(gomock.Any()).Return(nil, errors.New("any")).AnyTimes()
		answerRepo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()

		answerAttachments := factories.AnswerAttachmentsFactory(3, "1")
		answerAttachmentsRepo := mock.NewMockAnswerAttachmentsRepositoryInterface(ctrl)
		answerAttachmentsRepo.EXPECT().GetManyByAnswerID(gomock.Any()).Return(answerAttachments, nil).AnyTimes()

		useCase := uc.NewDefaultUpdateAnswerByIDUseCase(answerRepo, answerAttachmentsRepo)

		input := uc.UpdateAnswerByIDUseCaseInput{
			ID:       "1",
			AuthorID: "1",
			Content:  "Another Content",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
	})

	t.Run("should not update an answer when the author is not the same one who created the answer", func(t *testing.T) {
		answer := enterprise.NewAnswer("Content test", "1", "1")
		answerRepo := mock.NewMockAnswerRepositoryInterface(ctrl)
		answerRepo.EXPECT().GetByID(gomock.Any()).Return(answer, nil).AnyTimes()
		answerRepo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()

		answerAttachments := factories.AnswerAttachmentsFactory(3, "1")
		answerAttachmentsRepo := mock.NewMockAnswerAttachmentsRepositoryInterface(ctrl)
		answerAttachmentsRepo.EXPECT().GetManyByAnswerID(gomock.Any()).Return(answerAttachments, nil).AnyTimes()

		useCase := uc.NewDefaultUpdateAnswerByIDUseCase(answerRepo, answerAttachmentsRepo)

		input := uc.UpdateAnswerByIDUseCaseInput{
			ID:       "1",
			AuthorID: "2",
			Content:  "Another Content",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
		require.Equal(t, constants.NotAllowedError, err.Error())
	})
}
