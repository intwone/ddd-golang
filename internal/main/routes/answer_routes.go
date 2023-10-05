package routes

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/intwone/ddd-golang/internal/presentation/controllers"
)

func SetupAnswerRoutes(r *gin.Engine, controllers ctrl.AnswerControllers) {
	userGroup := r.Group("/api/answers")
	userGroup.POST("/", controllers.AnswerQuestionController.Handle)
}
