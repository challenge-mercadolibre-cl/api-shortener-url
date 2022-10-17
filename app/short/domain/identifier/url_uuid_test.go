package identifier

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlUUid(t *testing.T) {
	t.Parallel()
	t.Run("When generate correct uuid return uuid, because generate is correct", func(t *testing.T) {
		urlId, err := NewUrlUuid()
		assert.NoError(t, err)
		assert.Equal(t, len(urlId.Value()), 8)
	})
}
