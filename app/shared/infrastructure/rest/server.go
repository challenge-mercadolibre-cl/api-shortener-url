package rest

import (
	"fmt"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/infrastructure/log"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

const (
	httpReadTimeout  = 3 * time.Minute
	httpWriteTimeout = 3 * time.Minute
)

func New() *echo.Echo {
	server := echo.New()
	server.Use(log.EchoLogger())
	server.Use(echoMiddleware.Logger())
	server.Use(echoMiddleware.Recover())
	server.Use(echoMiddleware.CORS())
	server.Validator = NewValidator()
	server.HideBanner = true
	NewHealthHandler(server)
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
