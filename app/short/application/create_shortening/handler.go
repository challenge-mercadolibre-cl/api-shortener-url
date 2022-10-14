package create_shortening

import (
	"context"
	"fmt"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/application/command"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/application/exceptions"
)

type CommandHandlerCreateShortening struct {
	applicationService ServiceCreateShortening
}

func NewCommandHandlerCreateShortening(applicationService ServiceCreateShortening) CommandHandlerCreateShortening {
	return CommandHandlerCreateShortening{
		applicationService: applicationService,
	}
}

func (h CommandHandlerCreateShortening) Handle(ctx context.Context, cmd command.Command) error {
	command, ok := cmd.(CommandCreateShortening)
	if !ok {
		return fmt.Errorf("%w: %s", exceptions.ErrUnexpectedCommand, cmd.Type())
	}
	return h.applicationService.Do(ctx, command)
}
