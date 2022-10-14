package persistence

import (
	"context"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/aggregations"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/identifier"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/vo"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/infrastructure/exceptions"
	"github.com/go-redis/redis/v8"
)

type shortenerRepositoryRedis struct {
	redis *redis.Client
}

func NewShortenerRepositoryRedis(redis *redis.Client) shortenerRepositoryRedis {
	return shortenerRepositoryRedis{redis: redis}
}

func (s shortenerRepositoryRedis) Save(ctx context.Context, url aggregations.Url) error {
	err := s.redis.Set(ctx, url.UrlId().Value(), url.UrlLink().Value(), 0).Err()
	if err != nil {
		return err
	}
	return nil
}
func (s shortenerRepositoryRedis) Edit(ctx context.Context, url aggregations.Url, urlId identifier.UrlId) error {
	_, err := s.redis.Get(ctx, urlId.Value()).Result()
	if err == redis.Nil {
		return exceptions.ErrUrlIdNotFound
	}
	if err != nil {
		return err
	}
	err = s.Save(ctx, url)
	if err != nil {
		return err
	}
	return nil
}
func (s shortenerRepositoryRedis) Get(ctx context.Context, urlId identifier.UrlId) (vo.UrlLink, error) {
	data, err := s.redis.Get(ctx, urlId.Value()).Result()
	if err == redis.Nil {
		return vo.UrlLink{}, exceptions.ErrUrlIdNotFound
	}
	if err != nil {
		return vo.UrlLink{}, err
	}
	anUrlLink, err := vo.NewUrlLink(data)
	if err != nil {
		return vo.UrlLink{}, err
	}
	return anUrlLink, nil
}
