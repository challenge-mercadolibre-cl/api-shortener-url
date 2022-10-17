package identifier

import (
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/exceptions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlId(t *testing.T) {
	t.Parallel()
	t.Run("When send empty string throw error, because id is empty", func(t *testing.T) {
		_, err := NewUrlId("")
		assert.Error(t, err)
		assert.ErrorIs(t, err, exceptions.ErrUrlIdEmpty)
	})
	t.Run("When send correct id return good, because id is correct", func(t *testing.T) {
		expected := "testing"
		urlId, err := NewUrlId(expected)
		assert.NoError(t, err)
		assert.Equal(t, urlId.Value(), expected)
	})
}
