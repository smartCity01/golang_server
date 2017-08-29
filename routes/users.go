package routes

import (
	"../controllers"

	"github.com/labstack/echo"
)

func InitUsers(r *echo.Echo) {
	r.GET("/api/users/userinfo/:id", controllers.GetUser)
}
