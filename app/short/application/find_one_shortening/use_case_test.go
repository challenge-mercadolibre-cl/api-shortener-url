package find_one_shortening

import (
	"context"
	exceptionDomain "github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/exceptions"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/mocks/repositorymock"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/vo"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/infrastructure/exceptions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUseCaseFindOneShortening(t *testing.T) {

	t.Parallel()
	t.Run("When execute use case find one is successful, because record was found.", func(t *testing.T) {
		urlId := "0cDfV"
		ctx := context.Background()
		urlLink, err := vo.NewUrlLink("https://google.com")
		require.NoError(t, err)
		shortenerRepository := new(repositorymock.ShortenerRepository)
		shortenerRepository.On("Get", ctx, mock.AnythingOfType("identifier.UrlId")).
			Return(urlLink, nil)
		cmd := NewCommandFindOneShortening(urlId)
		service := NewUseCaseFindOneShortening(shortenerRepository)

		anUrlLink, err := service.Do(ctx, cmd)
		assert.Equal(t, anUrlLink.Value(), urlLink.Value())
		assert.NoError(t, err)
	})
	t.Run("When execute use case find one is failed, because record wasn't found.", func(t *testing.T) {
		urlId := "0cDfV"
		ctx := context.Background()
		shortenerRepository := new(repositorymock.ShortenerRepository)
		shortenerRepository.On("Get", ctx, mock.AnythingOfType("identifier.UrlId")).
			Return(vo.UrlLink{}, exceptions.ErrUrlIdNotFound)
		cmd := NewCommandFindOneShortening(urlId)
		service := NewUseCaseFindOneShortening(shortenerRepository)
		_, err := service.Do(ctx, cmd)
		assert.Error(t, err)
		assert.ErrorIs(t, err, exceptions.ErrUrlIdNotFound)
	})
	t.Run("When execute use case find one is failed, because url id is empty.", func(t *testing.T) {
		ctx := context.Background()
		shortenerRepository := new(repositorymock.ShortenerRepository)
		shortenerRepository.On("Get", ctx, mock.AnythingOfType("identifier.UrlId")).
			Return(vo.UrlLink{}, exceptionDomain.ErrUrlIdEmpty)
		cmd := NewCommandFindOneShortening("")
		service := NewUseCaseFindOneShortening(shortenerRepository)
		_, err := service.Do(ctx, cmd)
		assert.Error(t, err)
		assert.ErrorIs(t, err, exceptionDomain.ErrUrlIdEmpty)
	})
}
