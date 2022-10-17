package create_shortening

import (
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/application/command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommandCreateShortening(t *testing.T) {
	t.Parallel()
	t.Run("When execute command validate all arguments and is successful, because argument are valid.", func(t *testing.T) {
		url := "https://google.com"
		userId := "0000-0000"
		const typeCmd command.Type = "command.shortening.create"
		cmd := NewCommandCreateShortening(url, userId, "aaa")
		assert.Equal(t, cmd.Type(), typeCmd)
		assert.Equal(t, cmd.UserId(), userId)
		assert.Equal(t, cmd.Url(), url)
	})
}
