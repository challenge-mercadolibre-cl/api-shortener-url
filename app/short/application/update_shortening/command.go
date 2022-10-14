package update_shortening

import (
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/application/command"
)

const CommandTypeUpdateShortening command.Type = "command.shortening.update"

type CommandUpdateShortening struct {
	url    string
	userId string
	urlId  string
}

func NewCommandUpdateShortening(url string, userId string, urlId string) CommandUpdateShortening {
	return CommandUpdateShortening{
		url:    url,
		userId: userId,
		urlId:  urlId,
	}
}

func (c *CommandUpdateShortening) UserId() string {
	return c.userId
}

func (c *CommandUpdateShortening) Url() string {
	return c.url
}

func (c *CommandUpdateShortening) UrlId() string {
	return c.urlId
}
func (c CommandUpdateShortening) Type() command.Type {
	return CommandTypeUpdateShortening
}
