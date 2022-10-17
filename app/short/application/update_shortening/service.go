package update_shortening

import (
	"context"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/aggregations"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/repository"
)

type ServiceUpdateShortening interface {
	Do(ctx context.Context, command CommandUpdateShortening) error
}

//go:generate mockery --case=snake --outpkg=servicemock --output=../mocks/servicemock --name=ServiceUpdateShortening

type serviceUpdateShortening struct {
	shortenerRepository repository.ShortenerRepository
}

func NewServiceUpdateShortening(shortenerRepository repository.ShortenerRepository) serviceUpdateShortening {
	return serviceUpdateShortening{
		shortenerRepository: shortenerRepository,
	}
}

func (g serviceUpdateShortening) Do(ctx context.Context, command CommandUpdateShortening) error {
	anUrl, err := aggregations.NewUrl(command.Url(), command.UserId(), command.UrlId())
	if err != nil {
		return err
	}

	err = g.shortenerRepository.Edit(ctx, anUrl, anUrl.UrlId())
	if err != nil {
		return err
	}
	return nil
}
