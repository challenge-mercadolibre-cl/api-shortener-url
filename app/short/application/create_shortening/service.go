package create_shortening

import (
	"context"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/aggregations"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/repository"
)

type ServiceCreateShortening interface {
	Do(ctx context.Context, command CommandCreateShortening) (aggregations.Url, error)
}

type serviceCreateShortening struct {
	shortenerRepository repository.ShortenerRepository
}

func NewServiceCreateShortening(shortenerRepository repository.ShortenerRepository) serviceCreateShortening {
	return serviceCreateShortening{
		shortenerRepository: shortenerRepository,
	}
}

func (g serviceCreateShortening) Do(ctx context.Context, command CommandCreateShortening) (aggregations.Url, error) {
	anUrl, err := aggregations.NewUrl(command.Url(), command.UserId())
	if err != nil {
		return aggregations.Url{}, err
	}
	err = g.shortenerRepository.Save(ctx, anUrl)
	if err != nil {
		return aggregations.Url{}, err
	}
	return anUrl, nil
}
