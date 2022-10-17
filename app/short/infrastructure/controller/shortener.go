package controller

import (
	"fmt"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/application/command"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/infrastructure/log"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/application/create_shortening"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/application/find_one_shortening"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/application/update_shortening"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/identifier"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/infrastructure/exceptions"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ShortenerResponse struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}
type shortenerController struct {
	commandBus     command.CommandBus
	useCaseFindOne find_one_shortening.UseCaseFindOneShortening
	useCaseCreate  create_shortening.UseCaseCreateShortening
}

func NewShortenerHandler(e *echo.Echo, commandBus command.CommandBus, useCaseFindOne find_one_shortening.UseCaseFindOneShortening, useCaseCreate create_shortening.UseCaseCreateShortening) {
	h := &shortenerController{commandBus: commandBus, useCaseFindOne: useCaseFindOne, useCaseCreate: useCaseCreate}
	e.POST("/url/shortener", h.CreateShortenerUrl)
	e.PUT("/url/shortener/:id", h.EditShortenerUrl)
	e.GET("/url/shortener/:id", h.FindOneShortener)
}

type createRequest struct {
	UserId string `json:"user_id"`
	Url    string `json:"url"`
}
type updateRequest struct {
	UserId string `json:"user_id"`
	Url    string `json:"url"`
	UrlId  string `param:"id"`
}

type findOneRequest struct {
	UrlId string `param:"id"`
}

func (s shortenerController) FindOneShortener(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(findOneRequest)
	err := c.Bind(r)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	cmd := find_one_shortening.NewCommandFindOneShortening(r.UrlId)

	useCase, err := s.useCaseFindOne.Do(ctx, cmd)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, ShortenerResponse{
		Status: http.StatusOK,
		Data:   useCase.Value(),
	})

}
func (s shortenerController) CreateShortenerUrl(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(createRequest)
	err := c.Bind(r)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	
	urlUuid, err := identifier.NewUrlUuid()

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	cmd := create_shortening.NewCommandCreateShortening(r.Url, r.UserId, urlUuid.Value())
	useCase, err := s.useCaseCreate.Do(ctx, cmd)
	if err != nil {
		log.WithError(err).Error(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, ShortenerResponse{
		Status: 201,
		Data:   useCase.UrlId().Value(),
	})
}
func (s shortenerController) EditShortenerUrl(c echo.Context) error {
	ctx := c.Request().Context()
	u := new(updateRequest)
	err := c.Bind(u)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	cmd := update_shortening.NewCommandUpdateShortening(u.Url, u.UserId, u.UrlId)
	err = s.commandBus.Dispatch(ctx, cmd)
	if err == exceptions.ErrUrlIdNotFound {
		log.WithError(err).Error(err.Error())
		return c.JSON(http.StatusNotFound, map[string]interface{}{"error": fmt.Sprintf("%s : %s", exceptions.ErrUrlIdNotFound.Error(), u.UrlId)})
	}

	if err != nil {
		log.WithError(err).Error(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, ShortenerResponse{
		Status: 201,
		Data:   u.UrlId,
	})

}
