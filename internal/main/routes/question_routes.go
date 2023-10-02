package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/intwone/ddd-golang/internal/infra/cryptography"
	"github.com/intwone/ddd-golang/internal/main/adapters"
	ctrl "github.com/intwone/ddd-golang/internal/presentation/controllers"
)

func SetupQuestionRoutes(r *gin.Engine, controllers ctrl.QuestionControllers, crypto cryptography.CryptographyInterface) {
	questionGroup := r.Group("/api/questions")
	questionGroup.GET("/:questionSlug", controllers.GetQuestionBySlugController.Handle)
	questionGroup.GET("/recents", controllers.GetRecentQuestionsController.Handle)
	questionGroup.POST("/", controllers.CreateQuestionController.Handle)
	questionGroup.DELETE("/:questionID", adapters.GinGonicMiddlewareAdapter(crypto), controllers.DeleteQuestionByIDController.Handle)
}
