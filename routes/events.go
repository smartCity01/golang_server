package routes

import (
	"../controllers"

	"github.com/labstack/echo"
)

func InitEvents(r *echo.Echo) {
	r.GET("/api/events", controllers.GetEvents)
}
