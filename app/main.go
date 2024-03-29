package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/challenge-mercadolibre-cl/api-shortener-url/app/docs"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/infrastructure/bus/inmemory"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/infrastructure/rest"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/infrastructure/utils"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/application/create_shortening"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/application/find_one_shortening"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/application/update_shortening"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/infrastructure/controller"
	"github.com/challenge-mercadolibre-cl/api-shortener-url/app/short/infrastructure/persistence"
	"github.com/go-redis/redis/v8"
)

func main() {
	config := utils.ReadConfig()
	server := rest.New()
	commandBus := inmemory.NewCommandBusMemory()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
		Username: config.Redis.Username,
	})
	shortenerRepository := persistence.NewShortenerRepositoryRedis(rdb)

	shorteningUpdateService := update_shortening.NewServiceUpdateShortening(shortenerRepository)
	shorteningUpdateCommandHandler := update_shortening.NewCommandHandlerUpdateShortening(shorteningUpdateService)

	shorteningFindOneUseCase := find_one_shortening.NewUseCaseFindOneShortening(shortenerRepository)
	shorteningCreateUseCase := create_shortening.NewUseCaseCreateShortening(shortenerRepository)
	commandBus.Register(update_shortening.CommandTypeUpdateShortening, shorteningUpdateCommandHandler)

	controller.NewShortenerHandler(server, commandBus, shorteningFindOneUseCase, shorteningCreateUseCase)
	go func() {
		if err := server.StartServer(rest.Setup(config.Server.Host, config.Server.Port)); err != http.ErrServerClosed {
			server.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT)
	<-quit

	ctxServer, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctxServer); err != nil {
		server.Logger.Fatal(err)
	}

}
