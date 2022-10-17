package update_shortening

import (
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/application/command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommandUpdateShortening(t *testing.T) {
	t.Parallel()
	t.Run("When execute command validate all arguments and is successful, because argument are valid.", func(t *testing.T) {
		url := "https://google.com"
		userId := "0000-0000"
		urlId := "0cDfV"
		const typeCmd command.Type = "command.shortening.update"
		cmd := NewCommandUpdateShortening(url, userId, urlId)
		assert.Equal(t, cmd.Type(), typeCmd)
		assert.Equal(t, cmd.UrlId(), urlId)
		assert.Equal(t, cmd.UserId(), userId)
		assert.Equal(t, cmd.Url(), url)
	})
}
