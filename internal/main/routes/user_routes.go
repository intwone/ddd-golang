package routes

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/intwone/ddd-golang/internal/presentation/controllers"
)

func SetupUserRoutes(r *gin.Engine, controllers ctrl.UserControllers) {
	userGroup := r.Group("/api/users")
	userGroup.POST("/sign-up", controllers.SignUpController.Handle)
}
