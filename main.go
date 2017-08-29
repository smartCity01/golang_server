package main

import (
	"./routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := "1337"
	server := echo.New()

	startServer(port, server)
}

func startServer(port string, server *echo.Echo) {

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	routes.Init(server)
	server.Logger.Fatal(server.Start(":" + port))
}
