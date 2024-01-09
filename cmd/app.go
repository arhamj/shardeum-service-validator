package main

import (
	"context"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	defaultMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/shardeum/service-validator/common/db/sqlite"
	"github.com/shardeum/service-validator/config"
	"github.com/shardeum/service-validator/pkg/controllers"
	"github.com/shardeum/service-validator/pkg/middleware"
	"github.com/shardeum/service-validator/pkg/repo"
	"github.com/shardeum/service-validator/pkg/service"
	"github.com/shardeum/service-validator/pkg/util"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type app struct {
	validator *validator.Validate

	config config.Config

	db *gorm.DB

	accountRepo repo.AccountsEntryRepo

	accountService   service.AccountService
	validatorService service.ValidatorService

	accountController   controllers.AccountController
	validatorController controllers.ValidatorController

	scheduler *gocron.Scheduler

	appServer *echo.Echo
}

func (a *app) initConfig() {
	configFile := "./config/config.json"
	optCustomConfigFile, exists := os.LookupEnv("CONFIG_FILE_PATH")
	if exists {
		configFile = optCustomConfigFile
	} else {
		log.Warn("CONFIG_FILE_PATH env var not set, using default config file path")
	}
	cfg, err := config.NewConfig(configFile)
	if err != nil {
		log.Fatal("init config failed: ", err)
	}
	a.config = cfg
}

func (a *app) initValidator() {
	a.validator = validator.New()
}

func (a *app) initDB() {
	db, err := sqlite.NewGormDB(a.config.AppSqliteConfig)
	if err != nil {
		log.WithField("err", err).Fatal("init db failed")
	}
	log.Info("successfully initialised app sqlite database")

	a.db = db
}

func (a *app) initRepo() {
	a.accountRepo = repo.NewAccountsEntryRepo(a.db)
	log.Info("successfully initialised repos")
}

func (a *app) initService() {
	a.accountService = service.NewAccountService(a.accountRepo)
	a.validatorService = service.NewValidatorService(a.accountService)
	log.Info("successfully initialised services")
}

func (a *app) initScheduler() {
	a.scheduler = gocron.NewScheduler(time.UTC)
	a.scheduler.StartAsync()

	log.Info("successfully initialised scheduler")
}

func (a *app) initControllers() {
	a.accountController = controllers.NewAccountController(a.accountService)
	a.validatorController = controllers.NewValidatorController(a.validatorService)
	log.Info("successfully initialised controllers")
}

func (a *app) initServer() {
	e := echo.New()
	e.HideBanner = true
	e.Validator = util.NewRequestValidation(a.validator)
	e.Use(middleware.LoggingMiddleware())
	e.Use(defaultMiddleware.Recover())
	e.Use(defaultMiddleware.CORS())

	// Register routes
	e.GET("/accounts/:accountId", a.accountController.GetAccount)
	e.PUT("/accounts", a.accountController.GetAccounts)

	// Register shardeum equivalent routes
	e.GET("/account/:accountId", a.validatorController.GetEVMAccount)
	e.GET("/eth_getCode", a.validatorController.GetCodeBytes)

	a.appServer = e

	log.Info("starting app server...")
	log.Fatal(e.Start(":9100"))
}

func (a *app) shutdown(ctx context.Context) error {
	log.Info("shutting down scheduler...")
	a.scheduler.Stop()

	log.Info("shutting down servers...")
	err := a.appServer.Shutdown(ctx)
	if err != nil {
		return err
	}

	log.Info("shutting down db...")
	db, err := a.db.DB()
	if err != nil {
		return err
	}
	_ = db.Close()

	return nil
}
