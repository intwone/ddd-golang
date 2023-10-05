package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/intwone/ddd-golang/internal/constants"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/infra/cryptography"
	"github.com/intwone/ddd-golang/internal/infra/database/postgres/repositories"
	s "github.com/intwone/ddd-golang/internal/infra/database/sqlc"
	"github.com/intwone/ddd-golang/internal/infra/hasher"
	"github.com/intwone/ddd-golang/internal/main/routes"
	ctrl "github.com/intwone/ddd-golang/internal/presentation/controllers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()

	db, err := sql.Open(os.Getenv(constants.DATABASE_DIALECT), os.Getenv(constants.DATABASE_URL))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	dt := s.New(db)

	router := gin.Default()

	// Hasher
	bcryptHasher := hasher.NewBcryptHasher()

	// Cryptography
	secret := os.Getenv(constants.JWT_SECRET)
	jwtCryptography := cryptography.NewJWTCryptography(secret)

	// Answer
	answerSQLCRepository := repositories.NewAnswerSQLCRepository(dt)
	answerQuestionUseCase := uc.NewDefaultAnswerQuestionUseCase(answerSQLCRepository)
	answerQuestionController := ctrl.NewDefaultAnswerQuestionController(answerQuestionUseCase)

	answerControllers := ctrl.AnswerControllers{
		AnswerQuestionController: answerQuestionController,
	}

	// Question
	questionSQLCRepository := repositories.NewQuestionSQLCRepository(dt)
	getQuestionBySlugUseCase := uc.NewDefaulGetQuestionBySlugUseCase(questionSQLCRepository)
	getQuestionBySlugController := ctrl.NewDefaultGetQuestionBySlugController(getQuestionBySlugUseCase)
	createQuestionUseCase := uc.NewDefaultCreateQuestionUseCase(questionSQLCRepository)
	createQuestionController := ctrl.NewDefaultCreateQuestionController(createQuestionUseCase)
	deleteQuestionByIDUseCase := uc.NewDefaultDeleteQuestionByIDUseCase(questionSQLCRepository)
	deleteQuestionByIDController := ctrl.NewDefaultDeleteQuestionByIDController(deleteQuestionByIDUseCase)
	getRecentQuestionsUseCase := uc.NewDefaulGetRecentQuestionsUseCase(questionSQLCRepository)
	getRecentQuestionsController := ctrl.NewDefaultGetRecentQuestionsController(getRecentQuestionsUseCase)

	questionControllers := ctrl.QuestionControllers{
		CreateQuestionController:     createQuestionController,
		GetQuestionBySlugController:  getQuestionBySlugController,
		DeleteQuestionByIDController: deleteQuestionByIDController,
		GetRecentQuestionsController: getRecentQuestionsController,
	}

	// User
	userSQLCRepository := repositories.NewUserSQLCRepository(dt)
	createUserUseCase := uc.NewDefaultCreateUserUseCase(userSQLCRepository, bcryptHasher)
	getUserByEmailUseCase := uc.NewDefaulGetUserByEmailUseCase(userSQLCRepository)
	signUpController := ctrl.NewDefaultSignUpController(createUserUseCase, getUserByEmailUseCase)
	authenticateUseCase := uc.NewDefaulAuthenticateUseCase(userSQLCRepository, bcryptHasher, jwtCryptography)
	signInController := ctrl.NewDefaultSignInController(authenticateUseCase)

	userControllers := ctrl.UserControllers{
		SignUpController: signUpController,
		SignInController: signInController,
	}

	routes.SetupQuestionRoutes(router, questionControllers, jwtCryptography)
	routes.SetupUserRoutes(router, userControllers)
	routes.SetupAnswerRoutes(router, answerControllers)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
