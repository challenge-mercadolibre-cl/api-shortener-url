package create_shortening

import (
	"context"
	"errors"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/exceptions"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/mocks/repositorymock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUseCaseCreateShortening(t *testing.T) {

	t.Parallel()
	t.Run("When execute use case find one is successful, because record was found.", func(t *testing.T) {
		ctx := context.Background()
		shortenerRepository := new(repositorymock.ShortenerRepository)
		shortenerRepository.On("Save", ctx, mock.AnythingOfType("aggregations.Url")).
			Return(nil)
		url := "https://google.com"
		userId := "0000-0000"
		cmd := NewCommandCreateShortening(url, userId, "bbb")
		service := NewUseCaseCreateShortening(shortenerRepository)

		anUrl, err := service.Do(ctx, cmd)
		assert.Equal(t, anUrl.UrlLink().Value(), url)
		assert.Equal(t, anUrl.UserId().Value(), userId)
		assert.NoError(t, err)
	})
	t.Run("When execute service is failed, because has problem on repository.", func(t *testing.T) {
		ctx := context.Background()
		shortenerRepository := new(repositorymock.ShortenerRepository)
		errorRepository := errors.New("save failed")
		shortenerRepository.On("Save", ctx, mock.AnythingOfType("aggregations.Url")).
			Return(errorRepository)
		url := "https://google.com"
		userId := "0000-0000"
		cmd := NewCommandCreateShortening(url, userId, "aaa")
		service := NewUseCaseCreateShortening(shortenerRepository)

		_, err := service.Do(ctx, cmd)
		assert.Error(t, err)
		assert.ErrorIs(t, err, errorRepository)
	})
	t.Run("When execute service is failed, because url has invalid format.", func(t *testing.T) {
		ctx := context.Background()
		shortenerRepository := new(repositorymock.ShortenerRepository)
		shortenerRepository.On("Save", ctx, mock.AnythingOfType("aggregations.Url")).
			Return(nil)
		url := "abc"
		userId := "0000-0000"
		cmd := NewCommandCreateShortening(url, userId, "aaa")
		service := NewUseCaseCreateShortening(shortenerRepository)
		_, err := service.Do(ctx, cmd)
		assert.Error(t, err)
		assert.ErrorIs(t, err, exceptions.ErrUrlFormat)
	})
}
