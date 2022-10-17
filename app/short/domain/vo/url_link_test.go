package vo

import (
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/exceptions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlLink(t *testing.T) {
	t.Parallel()
	t.Run("When send incorrect format url return error, because format is invalid", func(t *testing.T) {
		_, err := NewUrlLink("abc")
		assert.Error(t, err)
		assert.ErrorIs(t, err, exceptions.ErrUrlFormat)
	})
	t.Run("When send correct format url return url, because format is valid.", func(t *testing.T) {
		expected := "https://google.com"
		anUrlLink, err := NewUrlLink(expected)
		assert.NoError(t, err)
		assert.Equal(t, anUrlLink.Value(), expected)
	})
}
