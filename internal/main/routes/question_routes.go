package routes

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/intwone/ddd-golang/internal/presentation/controllers"
)

func QuestionRoutes(r *gin.RouterGroup, controller ctrl.GetQuestionBySlugInterface) {
	r.GET("/:slug", controller.Handle)
}
