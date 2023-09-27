package controllers

type QuestionControllers struct {
	CreateQuestionController    *DefaultCreateQuestionControllerInterface
	GetQuestionBySlugController *DefaultGetQuestionBySlugControllerInterface
}

type UserControllers struct {
	CreateUserController *DefaultCreateUserControllerInterface
}
