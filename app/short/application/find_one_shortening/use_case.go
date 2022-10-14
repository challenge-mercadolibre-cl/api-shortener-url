package find_one_shortening

import (
	"context"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/identifier"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/repository"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/vo"
)

type UseCaseFindOneShortening interface {
	Do(ctx context.Context, command CommandFindOneShortening) (vo.UrlLink, error)
}

type useCaseFindOneShortening struct {
	shortenerRepository repository.ShortenerRepository
}

func NewUseCaseFindOneShortening(shortenerRepository repository.ShortenerRepository) useCaseFindOneShortening {
	return useCaseFindOneShortening{
		shortenerRepository: shortenerRepository,
	}
}

func (g useCaseFindOneShortening) Do(ctx context.Context, command CommandFindOneShortening) (vo.UrlLink, error) {
	anUrlId, err := identifier.NewUrlId(command.UrlId())
	if err != nil {
		return vo.UrlLink{}, err
	}
	anUrlLink, err := g.shortenerRepository.Get(ctx, anUrlId)
	if err != nil {
		return vo.UrlLink{}, err
	}
	return anUrlLink, nil
}
