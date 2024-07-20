package main

import (
	connectcors "connectrpc.com/cors"
	"context"
	"github.com/annp1987/sms-app/config"
	"github.com/annp1987/sms-app/logger"
	"github.com/annp1987/sms-app/proto/greetconnect"
	"github.com/annp1987/sms-app/server"
	"github.com/rs/cors"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			logger.ProvideLogger,
			server.NewServer),
		fx.Invoke(RegisterWebServer))
	app.Run()
}

func RegisterWebServer(lifeCycle fx.Lifecycle, greeter *server.GreetServer, conf *config.Config, logger *zap.Logger) {
	lifeCycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				mux := http.NewServeMux()
				path, handler := greetconnect.NewGreetServiceHandler(greeter)
				mux.Handle(path, withCORS(handler))
				http.ListenAndServe(
					":8080",
					// Use h2c so we can serve HTTP/2 without TLS.
					h2c.NewHandler(mux, &http2.Server{}),
				)
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			logger.Info("stopping server ...")
			return nil
		},
	})
	logger.Info("web server started on ", zap.String("port", conf.ServicePort))
}

func withCORS(h http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}
