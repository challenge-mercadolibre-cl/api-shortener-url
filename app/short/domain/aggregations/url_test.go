package aggregations

import (
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/exceptions"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/identifier"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrl(t *testing.T) {
	t.Parallel()
	t.Run("When send empty string in user id throw error, because user id is empty", func(t *testing.T) {
		_, err := NewUrl("https://google.cl", "", "abc")
		assert.Error(t, err)
		assert.ErrorIs(t, err, exceptions.ErrUserIdEmpty)
	})
	t.Run("When send incorrect format link throw error, because link is invalid", func(t *testing.T) {
		_, err := NewUrl("abc", "1", "abc")
		assert.Error(t, err)
		assert.ErrorIs(t, err, exceptions.ErrUrlFormat)
	})
	t.Run("When send correct format link and user Id is successful, because arguments are valid", func(t *testing.T) {
		anUrl, err := NewUrl("https://google.com", "1", "abc")
		assert.NoError(t, err)
		assert.Equal(t, anUrl.UserId().Value(), "1")
		assert.Equal(t, anUrl.UrlLink().Value(), "https://google.com")
		assert.IsType(t, anUrl.UrlId(), identifier.UrlId{})
	})
}
