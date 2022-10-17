package repository

import (
	"context"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/aggregations"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/identifier"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/vo"
)

type ShortenerRepository interface {
	Save(ctx context.Context, url aggregations.Url) error
	Edit(ctx context.Context, url aggregations.Url, urlId identifier.UrlId) error
	Get(ctx context.Context, urlId identifier.UrlId) (vo.UrlLink, error)
}

//go:generate mockery --case=snake --outpkg=repositorymock --output=../mocks/repositorymock --name=ShortenerRepository
