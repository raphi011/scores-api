package app

import (
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/raphi011/scores-api/events"
)

// App wraps all the services and configuration needed
// to serve the api.
type App struct {
	conf        *oauth2.Config
	services    *handlerServices
	eventBroker *events.Broker
	version     string
	production  bool
}

// Option is used to configure a new Router.
type Option func(*App)

// New creates a new router and configures it with `opts`.
func New(opts ...Option) *App {
	app := &App{}

	for _, o := range opts {
		o(app)
	}

	return app
}

// Run builds the router and opens the port.
func (r *App) Run() {
	router := r.Build()

	err := router.Run()

	if err != nil {
		zap.S().Errorf("could not start router: %+v", err)
	}
}
