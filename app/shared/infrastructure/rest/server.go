package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/infrastructure/log"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.elastic.co/apm/module/apmechov4/v2"
	"net/http"
	"time"
)

const (
	httpReadTimeout  = 3 * time.Minute
	httpWriteTimeout = 3 * time.Minute
)

func New() *echo.Echo {
	server := echo.New()
	server.Use(apmechov4.Middleware())
	server.Use(log.EchoLogger())
	server.Use(echoMiddleware.Logger())
	server.Use(echoMiddleware.Recover())
	server.Use(echoMiddleware.CORS())
	server.Validator = NewValidator()
	server.HideBanner = true
	NewHealthHandler(server)
	server.GET("/swagger/*", echoSwagger.EchoWrapHandler(func(c *echoSwagger.Config) { c.URL = "./swagger/doc.json" }))
	return server
}

func Setup(host string, port string) *http.Server {
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		ReadTimeout:  httpReadTimeout,
		WriteTimeout: httpWriteTimeout,
	}
	return server

}
