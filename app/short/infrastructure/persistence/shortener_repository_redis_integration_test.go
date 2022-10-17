package persistence

import (
	"context"
	"fmt"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/aggregations"
	exceptionDomain "github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/exceptions"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/domain/identifier"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/infrastructure/exceptions"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
	"time"
)

type redisContainer struct {
	testcontainers.Container
	URI string
}

func setupRedis(ctx context.Context) (*redisContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "redis:6",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("* Ready to accept connections"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "6379")
	if err != nil {
		return nil, err
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("redis://%s:%s", hostIP, mappedPort.Port())

	return &redisContainer{Container: container, URI: uri}, nil
}

func flushRedis(ctx context.Context, client redis.Client) error {
	return client.FlushAll(ctx).Err()
}

func TestShortenerRepositoryRedis(t *testing.T) {
	ctx := context.Background()
	redisContainer, err := setupRedis(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer redisContainer.Terminate(ctx)

	options, err := redis.ParseURL(redisContainer.URI)
	if err != nil {
		t.Fatal(err)
	}
	redisClient := redis.NewClient(options)
	repository := NewShortenerRepositoryRedis(redisClient)

	badUri := fmt.Sprintf("redis://%s:%s", "localhost", "6320")
	optionsBad, err := redis.ParseURL(badUri)
	require.NoError(t, err)
	badRedisClient := redis.NewClient(optionsBad)
	repositoryBad := NewShortenerRepositoryRedis(badRedisClient)

	t.Run("When save record is successful, because arguments are valid", func(t *testing.T) {
		anUrlId := "xxx-xxxx"
		anUrlLink := "https://google.com"
		anUserId := "000"
		anUrl, err := aggregations.NewUrl(anUrlLink, anUserId, anUrlId)
		require.NoError(t, err)
		err = repository.Save(ctx, anUrl)
		assert.NoError(t, err)
	})
	t.Run("When force save record is successful, because arguments are valid", func(t *testing.T) {
		require.NoError(t, err)
		err := redisClient.Set(ctx, "key-testing", "abc", 0).Err()
		assert.NoError(t, err)
	})

	t.Run("When save record is failed, because repository bad configuration.", func(t *testing.T) {
		anUrlId := "xxx-xxxx"
		anUrlLink := "https://google.com"
		anUserId := "000"
		anUrl, err := aggregations.NewUrl(anUrlLink, anUserId, anUrlId)
		require.NoError(t, err)
		err = repositoryBad.Save(ctx, anUrl)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "connection refused")
	})
	t.Run("When edit record is successful, because record exist.", func(t *testing.T) {
		anUrlId := "xxx-xxxx"
		anUrlLink := "https://google.com"
		anUserId := "000"
		anUrlIdIdentifier, err := identifier.NewUrlId(anUrlId)
		require.NoError(t, err)
		anUrl, err := aggregations.NewUrl(anUrlLink, anUserId, anUrlId)
		require.NoError(t, err)
		err = repository.Edit(ctx, anUrl, anUrlIdIdentifier)
		assert.NoError(t, err)
	})
	t.Run("When edit record is failed, because record not found.", func(t *testing.T) {
		anUrlId := "123c"
		anUrlLink := "https://google.com"
		anUserId := "000"
		anUrlIdIdentifier, err := identifier.NewUrlId(anUrlId)
		require.NoError(t, err)
		anUrl, err := aggregations.NewUrl(anUrlLink, anUserId, anUrlId)
		require.NoError(t, err)
		err = repository.Edit(ctx, anUrl, anUrlIdIdentifier)
		assert.Error(t, err)
		assert.ErrorIs(t, err, exceptions.ErrUrlIdNotFound)
	})
	t.Run("When edit record is failed, because repository bad configuration.", func(t *testing.T) {
		anUrlId := "123c"
		anUrlLink := "https://google.com"
		anUserId := "000"
		anUrlIdIdentifier, err := identifier.NewUrlId(anUrlId)
		require.NoError(t, err)
		anUrl, err := aggregations.NewUrl(anUrlLink, anUserId, anUrlId)
		require.NoError(t, err)
		err = repositoryBad.Edit(ctx, anUrl, anUrlIdIdentifier)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "connection refused")
	})

	t.Run("When get record is successful, because record exist.", func(t *testing.T) {
		anUrlId := "xxx-xxxx"
		anUrlLinkExpected := "https://google.com"

		anUrlIdIdentifier, err := identifier.NewUrlId(anUrlId)
		require.NoError(t, err)
		anUrlLink, err := repository.Get(ctx, anUrlIdIdentifier)
		assert.NoError(t, err)
		assert.Equal(t, anUrlLink.Value(), anUrlLinkExpected)
	})
	t.Run("When get record is failed, because record not found.", func(t *testing.T) {
		anUrlId := "xxx"
		anUrlIdIdentifier, err := identifier.NewUrlId(anUrlId)
		require.NoError(t, err)
		_, err = repository.Get(ctx, anUrlIdIdentifier)
		assert.Error(t, err)
		assert.ErrorIs(t, err, exceptions.ErrUrlIdNotFound)
	})
	t.Run("When get record is failed, because repository bad configuration.", func(t *testing.T) {
		anUrlId := "xxx"
		anUrlIdIdentifier, err := identifier.NewUrlId(anUrlId)
		require.NoError(t, err)
		_, err = repositoryBad.Get(ctx, anUrlIdIdentifier)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "connection refused")
	})
	t.Run("When get record is failed, because url link has invalid format", func(t *testing.T) {
		anUrlId := "key-testing"
		anUrlIdIdentifier, err := identifier.NewUrlId(anUrlId)
		require.NoError(t, err)
		_, err = repository.Get(ctx, anUrlIdIdentifier)
		assert.Error(t, err)
		assert.ErrorIs(t, err, exceptionDomain.ErrUrlFormat)

	})
}

func destroyContainer(compose *testcontainers.LocalDockerCompose) {
	compose.Down()
	time.Sleep(1 * time.Second)
}
