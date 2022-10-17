package identifier

import (
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/exceptions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserId(t *testing.T) {
	t.Parallel()
	t.Run("When send empty string throw error, because user id is empty", func(t *testing.T) {
		_, err := NewUserId("")
		assert.Error(t, err)
		assert.ErrorIs(t, err, exceptions.ErrUserIdEmpty)
	})
	t.Run("When send correct user id return good, because user id is correct", func(t *testing.T) {
		expected := "testing"
		urlId, err := NewUserId(expected)
		assert.NoError(t, err)
		assert.Equal(t, urlId.Value(), expected)
	})
}
