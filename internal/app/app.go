package app

import (
	"URL_Shortener/internal/config"
	"URL_Shortener/internal/handlers/rest"
	"URL_Shortener/internal/repo"
	"URL_Shortener/internal/server"
	"URL_Shortener/internal/service"
	"URL_Shortener/internal/storage"
	"URL_Shortener/internal/utils"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type App struct {
	service *service.Service
	storage.Storage
	router *gin.Engine
	server *server.Server
	cfg    *config.Config
	logger *zap.Logger
}

func New(ctx context.Context) (app *App, err error) {
	app = &App{}

	app.logger = utils.NewLogger()

	app.logger.Info("config initializing")
	app.cfg, err = config.New()
	if err != nil {
		app.logger.Error(err.Error())
		return nil, err
	}

	app.logger.Info("storage initializing")
	if app.cfg.UseMemory {
		app.Storage = storage.NewMemoryStorage()
		app.logger.Info("using in-memory storage")
	} else {
		app.Storage, err = storage.NewDatabaseConnection(&app.cfg.DB)
		if err != nil {
			app.logger.Error(err.Error())
			return nil, err
		}
		app.logger.Info("database connected")
	}

	repo := repo.New(&app.Storage)
	app.service = service.New(repo, app.Storage, app.logger)

	rest := rest.New(app.service)
	app.router = rest.Route()

	app.server, err = server.NewServer(&app.cfg.HTTP, app.router.Handler())
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) Run(ctx context.Context) error {
	app.logger.Info("server starting")

	if err := app.server.Start(); err != nil {
		return err
	}

	return nil
}

func (app *App) Shutdown(ctx context.Context) error {
	return app.server.Shutdown(ctx)
}
