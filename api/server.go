package api

import (
	"log"
	"net/http"

	"github.com/AhmedMohamed800/pong/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	e := echo.New()


	// Middleware
	e.Use(middleware.Logger())
 	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    	AllowOrigins: []string{"*"},
    	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
  	}))


	// Routes
	e.GET("/", func (c echo.Context) error {
    	db, err := storage.InitDB()
		if err != nil {
    		log.Fatal(err)
    	}
		defer db.Close()


		return c.String(http.StatusOK, "hello")
	})

	if err := e.Start(":1323"); err != nil {
		e.Logger.Info("Shutting down the server")
	}
}