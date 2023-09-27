package routes

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/intwone/ddd-golang/internal/presentation/controllers"
)

func SetupQuestionRoutes(r *gin.Engine, controllers ctrl.QuestionControllers) {
	questionGroup := r.Group("/api/questions")
	questionGroup.GET("/:slug", controllers.GetQuestionBySlugController.Handle)
	questionGroup.POST("/", controllers.CreateQuestionController.Handle)
	questionGroup.DELETE("/:id", controllers.DeleteQuestionByIDController.Handle)
}
