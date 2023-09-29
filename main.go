package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

	db, err := sql.Open(os.Getenv("DATABASE_DIALECT"), os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	dt := s.New(db)

	router := gin.Default()

	// Hasher
	bcryptHasher := hasher.NewBcryptHasher()

	// Cryptography
	secret := os.Getenv("JWT_SECRET")
	secretByte := []byte(secret)
	jwtCryptography := cryptography.NewJWTCryptography(secretByte)

	// Question
	questionSQLCRepository := repositories.NewQuestionSQLCRepository(dt)
	getQuestionBySlugUseCase := uc.NewDefaulGetQuestionBySlugUseCase(questionSQLCRepository)
	getQuestionBySlugController := ctrl.NewDefaultGetQuestionBySlugController(getQuestionBySlugUseCase)
	createQuestionUseCase := uc.NewDefaultCreateQuestionUseCase(questionSQLCRepository)
	createQuestionController := ctrl.NewDefaultCreateQuestionController(createQuestionUseCase)
	deleteQuestionByIDUseCase := uc.NewDefaultDeleteQuestionByIDUseCase(questionSQLCRepository)
	deleteQuestionByIDController := ctrl.NewDefaultDeleteQuestionByIDController(deleteQuestionByIDUseCase)

	questionControllers := ctrl.QuestionControllers{
		CreateQuestionController:     createQuestionController,
		GetQuestionBySlugController:  getQuestionBySlugController,
		DeleteQuestionByIDController: deleteQuestionByIDController,
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

	routes.SetupQuestionRoutes(router, questionControllers)
	routes.SetupUserRoutes(router, userControllers)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
