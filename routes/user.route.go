package routes

import (
	"github.com/This-Is-Prince/golang-mongodb-api/controllers"
	"github.com/This-Is-Prince/golang-mongodb-api/middleware"
	"github.com/This-Is-Prince/golang-mongodb-api/services"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserRouteController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("users")
	router.Use(middleware.DeserializeUser(userService))
	router.GET("/me", uc.userController.GetMe)
}
