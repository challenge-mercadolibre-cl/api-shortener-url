package find_one_shortening

import (
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/application/command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommandFindOneShortening(t *testing.T) {
	t.Parallel()
	t.Run("When execute command validate all arguments and is successful, because argument are valid.", func(t *testing.T) {
		urlId := "0cDfV"
		const typeCmd command.Type = "command.shortening.find.one"
		cmd := NewCommandFindOneShortening(urlId)
		assert.Equal(t, cmd.Type(), typeCmd)
		assert.Equal(t, cmd.UrlId(), urlId)
	})
}
