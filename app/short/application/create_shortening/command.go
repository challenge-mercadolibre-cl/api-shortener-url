package create_shortening

import (
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/application/command"
)

const CommandTypeCreateShortening command.Type = "command.shortening.create"

type CommandCreateShortening struct {
	url    string
	userId string
}

func NewCommandCreateShortening(url string, userId string) CommandCreateShortening {
	return CommandCreateShortening{
		url:    url,
		userId: userId,
	}
}

func (c *CommandCreateShortening) UserId() string {
	return c.userId
}

func (c *CommandCreateShortening) Url() string {
	return c.url
}
func (c CommandCreateShortening) Type() command.Type {
	return CommandTypeCreateShortening
}
