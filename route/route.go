package route

import (
	"go-hero/handler"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Routes
	v1 := e.Group("/api/v1/hero")
	{
		v1.GET("/search/:name", handler.GetHeros)
		v1.GET("/detail/:id", handler.GetDetail)
	}
	return e
}
