package find_one_shortening

import (
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/application/command"
)

const CommandTypeCreateShortening command.Type = "command.shortening.find.one"

type CommandFindOneShortening struct {
	urlId string
}

func NewCommandFindOneShortening(urlId string) CommandFindOneShortening {
	return CommandFindOneShortening{
		urlId: urlId,
	}
}

func (c *CommandFindOneShortening) UrlId() string {
	return c.urlId
}
func (c CommandFindOneShortening) Type() command.Type {
	return CommandTypeCreateShortening
}
