package update_shortening

import (
	"context"
	"errors"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/exceptions"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/mocks/repositorymock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServiceUpdateShortening(t *testing.T) {
	t.Parallel()
	t.Run("When execute service is successful, because arguments are valid..", func(t *testing.T) {
		ctx := context.Background()
		shortenerRepository := new(repositorymock.ShortenerRepository)
		shortenerRepository.On("Edit", ctx, mock.AnythingOfType("aggregations.Url"), mock.AnythingOfType("identifier.UrlId")).
			Return(nil)
		url := "https://google.com"
		userId := "0000-0000"
		urlId := "0cDfV"
		cmd := NewCommandUpdateShortening(url, userId, urlId)
		service := NewServiceUpdateShortening(shortenerRepository)

		err := service.Do(ctx, cmd)
		assert.NoError(t, err)
	})
	t.Run("When execute service is failed, because url format is invalid.", func(t *testing.T) {
		ctx := context.Background()
		shortenerRepository := new(repositorymock.ShortenerRepository)
		shortenerRepository.On("Edit", ctx, mock.AnythingOfType("aggregations.Url"), mock.AnythingOfType("identifier.UrlId")).
			Return(nil)
		url := ""
		userId := "0000-0000"
		urlId := "0cDfV"
		cmd := NewCommandUpdateShortening(url, userId, urlId)
		service := NewServiceUpdateShortening(shortenerRepository)

		err := service.Do(ctx, cmd)
		assert.Error(t, err)
		assert.ErrorIs(t, err, exceptions.ErrUrlFormat)
	})
	t.Run("When execute service is failed, because url id is empty.", func(t *testing.T) {
		ctx := context.Background()
		shortenerRepository := new(repositorymock.ShortenerRepository)
		shortenerRepository.On("Edit", ctx, mock.AnythingOfType("aggregations.Url"), mock.AnythingOfType("identifier.UrlId")).
			Return(nil)
		url := "https://google.com"
		userId := "0000-0000"
		urlId := ""
		cmd := NewCommandUpdateShortening(url, userId, urlId)
		service := NewServiceUpdateShortening(shortenerRepository)

		err := service.Do(ctx, cmd)
		assert.Error(t, err)
		assert.ErrorIs(t, err, exceptions.ErrUrlIdEmpty)
	})
	t.Run("When execute service is failed, because has problem on repository.", func(t *testing.T) {
		ctx := context.Background()
		errorRepository := errors.New("error save record")
		shortenerRepository := new(repositorymock.ShortenerRepository)
		shortenerRepository.On("Edit", ctx, mock.AnythingOfType("aggregations.Url"), mock.AnythingOfType("identifier.UrlId")).
			Return(errorRepository)
		url := "https://google.com"
		userId := "0000-0000"
		urlId := "0000x"
		cmd := NewCommandUpdateShortening(url, userId, urlId)
		service := NewServiceUpdateShortening(shortenerRepository)

		err := service.Do(ctx, cmd)
		assert.Error(t, err)
		assert.ErrorIs(t, err, errorRepository)
	})
}
