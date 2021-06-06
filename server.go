package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
//	"github.com/labstack/echo/v4/middleware"
	"lottery/controllers"
)

func main() {
	e := echo.New()

	  // ログなど
	//  echo.Use(middleware.Logger())
	//  echo.Use(middleware.Recover())

	prefix := "/api/v1"
	e.GET(prefix+"/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST(prefix+ "/users", controllers.GetAllUsers)
	e.GET(prefix+ "/users", controllers.GetAllUsers)	
	e.POST(prefix+ "/users/:userID/points", controllers.GetAllUsers)
	e.GET(prefix+ "/users/:userID/points", controllers.GetAllUsers)	
	e.HEAD(prefix+ "/users/:userID/draw", controllers.GetAllUsers)
	e.PUT(prefix+ "/users/:userID/draw/:place", controllers.GetAllUsers)	

	e.Logger.Fatal(e.Start(":12111"))
}