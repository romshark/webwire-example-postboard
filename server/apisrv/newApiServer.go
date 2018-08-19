package apisrv

import (
	"log"
	"net/http"
	"regexp"
	"sync"

	"github.com/pkg/errors"
	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-messenger/server/apisrv/authorizer"
	"github.com/qbeon/webwire-messenger/server/apisrv/config"
	"github.com/qbeon/webwire-messenger/server/apisrv/dam"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"
	"github.com/qbeon/webwire-messenger/server/apisrv/logger"
	"github.com/qbeon/webwire-messenger/server/apisrv/resolver"
	"github.com/qbeon/webwire-messenger/server/apisrv/validator"
)

// NewApiServer initializes a new API server instance
func NewApiServer(
	conf config.Config,
	engine engiface.Engine,
) (ApiServer, error) {
	// Initialize a logger module instance
	logger, err := logger.New(&conf)
	if err != nil {
		return nil, errors.Wrap(err, "logger initialization failed")
	}

	// Initialize a validator module instance
	validator := validator.NewValidator(
		// TODO: add UTF8 support
		regexp.MustCompile("^[a-zA-Z]{1,64}$"), // name rule
		1,   // minimum message length
		255, // maximum message length
		1,   // minimum reaction description length
		256, // maximum reaction description length
		10,  // maximum messages query limit
	)

	// Initialize a new API server instance
	newApiServer := &apiServer{
		conf: conf,
		lock: &sync.RWMutex{},
		stop: dam.New(1),
		// Initialize a resolver module instance
		resolver: resolver.New(
			logger,
			validator,
			// Initialize an authorizer module instance
			authorizer.New(),
			engine,
		),
		log:    logger,
		engine: engine,
	}

	// Initialize a webwire server instance
	newApiServer.wwrSrv, err = wwr.NewHeadlessServer(
		newApiServer,
		wwr.ServerOptions{
			// Enable webwire sessions
			Sessions: wwr.Enabled,

			// Make the API server responsible for handling the sessions
			SessionManager: newApiServer,

			// Define the session info parser
			SessionInfoParser: engiface.ParseSessionInfo,
			Heartbeat:         wwr.Enabled,

			// Use the log writers provided by the logger instance
			ErrorLog: log.New(
				logger.ErrorLogWriter(),
				"WWR_ERR: ",
				log.Ldate|log.Ltime|log.Lshortfile,
			),
			WarnLog: log.New(
				logger.ErrorLogWriter(),
				"WWR_WARN: ",
				log.Ldate|log.Ltime|log.Lshortfile,
			),
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "webwire server initialization failed")
	}

	// Initialize the example HTTP endpoint
	newApiServer.exampleHTTPEndpoint = NewExampleHTTPEndpoint(newApiServer)

	// Initialize the HTTP endpoint server
	// that's hosting the underlying webwire server
	newApiServer.httpSrv = &http.Server{
		Addr:    conf.ServerAddress,
		Handler: newApiServer,
	}

	// Initialize the metrics server
	metricsHandler := NewMetricsHandler()
	newApiServer.metricsSrv = &http.Server{
		Addr:    conf.MetricsServerAddress,
		Handler: metricsHandler,
	}

	return newApiServer, nil
}
