package routes

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/intwone/ddd-golang/internal/presentation/controllers"
)

func QuestionRoutes(r *gin.Engine, controller ctrl.DefaultGetQuestionBySlugInterface) {
	questionGroup := r.Group("/api/questions")
	questionGroup.GET("/:slug", controller.HandleGetQuestionBySlug)
}
