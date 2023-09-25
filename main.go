package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/infra/database/postgres/repositories"
	s "github.com/intwone/ddd-golang/internal/infra/database/sqlc"
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

	questionSQLCRepository := repositories.NewQuestionSQLCRepository(dt)

	getQuestionBySlugUseCase := uc.NewDefaulGetQuestionBySlugUseCase(questionSQLCRepository)
	getQuestionController := ctrl.NewDefaultGetQuestionBySlug(getQuestionBySlugUseCase)

	// createQuestionUseCase := uc.NewDefaultCreateQuestionUseCase(questionSQLCRepository)
	// createQuestionController := ctrl.NewDefaultCreateQuestion(createQuestionUseCase)

	routes.QuestionRoutes(router, *getQuestionController)
	// routes.QuestionRoutes(router, *getQuestionController)

	if err := router.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
