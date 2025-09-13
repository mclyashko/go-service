package di

import (
	"fmt"
	"log"

	"github.com/mclyashko/go-service/internal/config"
	"github.com/mclyashko/go-service/internal/http"
	"github.com/mclyashko/go-service/internal/http/handler/hello_world"
	jokeHandler "github.com/mclyashko/go-service/internal/http/handler/joke"
	"github.com/mclyashko/go-service/internal/http/handler/static"

	jokeProvider "github.com/mclyashko/go-service/internal/provider/joke"
)

type Container struct {
	Config config.Config
	Logger *log.Logger

	JokeProvider jokeProvider.JokeProvider

	HelloHandler  hello_world.Handler
	JokeHandler   jokeHandler.Handler
	StaticHandler static.Handler

	Server *http.Server
}

func NewContainer() *Container {
	cfg := config.Load()
	logger := log.New(log.Writer(), fmt.Sprintf("[%s]", cfg.Name), log.LstdFlags|log.Lshortfile)

	jokeProvider := jokeProvider.NewHTTPJokeProvider()

	helloHandler := hello_world.New()
	jokeHandler := jokeHandler.New(logger, jokeProvider)
	staticHandler := static.New()

	srv := http.NewServer(
		cfg.Addr,
		helloHandler,
		jokeHandler,
		staticHandler,
	)

	return &Container{
		Config: cfg,
		Logger: logger,

		JokeProvider: jokeProvider,

		HelloHandler:  helloHandler,
		JokeHandler:   jokeHandler,
		StaticHandler: staticHandler,

		Server: srv,
	}
}
