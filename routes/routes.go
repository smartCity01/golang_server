package routes

import (
	"github.com/labstack/echo"
)

func Init(r *echo.Echo) {
	InitEvents(r)
	InitUsers(r)
}
