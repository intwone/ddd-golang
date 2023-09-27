package controllers

type QuestionControllers struct {
	CreateQuestionController     *DefaultCreateQuestionControllerInterface
	GetQuestionBySlugController  *DefaultGetQuestionBySlugControllerInterface
	DeleteQuestionByIDController *DefaultDeleteQuestionByIDControllerInterface
}

type UserControllers struct {
	CreateUserController *DefaultCreateUserControllerInterface
}
