package create_shortening

import (
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/application/command"
)

const CommandTypeCreateShortening command.Type = "command.shortening.create"

type CommandCreateShortening struct {
	url     string
	userId  string
	urlUuid string
}

func NewCommandCreateShortening(url string, userId string, urlUuid string) CommandCreateShortening {
	return CommandCreateShortening{
		url:     url,
		userId:  userId,
		urlUuid: urlUuid,
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

func (c CommandCreateShortening) UrlUuid() string {
	return c.urlUuid
}
