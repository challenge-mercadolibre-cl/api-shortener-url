package update_shortening

import (
	"context"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/application/command"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/application/mocks/commandmocks"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/application/exceptions"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/mocks/repositorymock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCommandHandlerUpdateShortening(t *testing.T) {
	t.Parallel()
	t.Run("When execute command handler is successful, because all arguments are valid. ", func(t *testing.T) {
		ctx := context.Background()
		shortenerRepository := new(repositorymock.ShortenerRepository)
		shortenerRepository.On("Edit", ctx, mock.AnythingOfType("aggregations.Url"), mock.AnythingOfType("identifier.UrlId")).
			Return(nil)
		url := "https://google.com"
		userId := "0000-0000"
		urlId := "0cDfV"
		cmd := NewCommandUpdateShortening(url, userId, urlId)
		service := NewServiceUpdateShortening(shortenerRepository)
		handler := NewCommandHandlerUpdateShortening(service)
		err2 := handler.Handle(ctx, cmd)
		assert.NoError(t, err2)
	})
	t.Run("When execute command handler is failed, because command not found. ", func(t *testing.T) {
		ctx := context.Background()
		var commandMockType command.Type = "command.mock"
		cmdMock := commandmocks.NewCommand(t)
		cmdMock.On("Type").Return(commandMockType)
		shortenerRepository := new(repositorymock.ShortenerRepository)
		shortenerRepository.On("Edit", ctx, mock.AnythingOfType("aggregations.Url"), mock.AnythingOfType("identifier.UrlId")).
			Return(nil)
		service := NewServiceUpdateShortening(shortenerRepository)
		handler := NewCommandHandlerUpdateShortening(service)
		err2 := handler.Handle(ctx, cmdMock)
		assert.Error(t, err2)
		assert.ErrorIs(t, err2, exceptions.ErrUnexpectedCommand)
	})
}
