package update_shortening

import (
	"context"
	"fmt"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/application/command"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/application/exceptions"
)

type CommandHandlerUpdateShortening struct {
	applicationService ServiceUpdateShortening
}

func NewCommandHandlerUpdateShortening(applicationService ServiceUpdateShortening) CommandHandlerUpdateShortening {
	return CommandHandlerUpdateShortening{
		applicationService: applicationService,
	}
}

func (h CommandHandlerUpdateShortening) Handle(ctx context.Context, cmd command.Command) error {
	command, ok := cmd.(CommandUpdateShortening)
	if !ok {
		return fmt.Errorf("%w: %s", exceptions.ErrUnexpectedCommand, cmd.Type())
	}
	return h.applicationService.Do(ctx, command)
}
