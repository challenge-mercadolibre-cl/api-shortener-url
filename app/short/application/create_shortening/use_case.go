package create_shortening

import (
	"context"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/aggregations"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/repository"
)

type UseCaseCreateShortening interface {
	Do(ctx context.Context, command CommandCreateShortening) (aggregations.Url, error)
}

type useCaseCreateShortening struct {
	shortenerRepository repository.ShortenerRepository
}

func NewUseCaseCreateShortening(shortenerRepository repository.ShortenerRepository) useCaseCreateShortening {
	return useCaseCreateShortening{
		shortenerRepository: shortenerRepository,
	}
}

func (g useCaseCreateShortening) Do(ctx context.Context, command CommandCreateShortening) (aggregations.Url, error) {
	anUrl, err := aggregations.NewUrl(command.Url(), command.UserId(), command.UrlUuid())
	if err != nil {
		return aggregations.Url{}, err
	}
	err = g.shortenerRepository.Save(ctx, anUrl)
	if err != nil {
		return aggregations.Url{}, err
	}
	return anUrl, nil
}
