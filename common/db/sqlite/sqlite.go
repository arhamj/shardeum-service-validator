package sqlite

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	DbPath   string              `mapstructure:"db_path"`
	LogLevel gormLogger.LogLevel `mapstructure:"log_level"`
}

func NewGormDB(databaseConfig DatabaseConfig) (*gorm.DB, error) {
	dbLogger := gormLogger.New(
		log.StandardLogger(),
		gormLogger.Config{
			SlowThreshold:             time.Second,             // Slow SQL threshold
			LogLevel:                  databaseConfig.LogLevel, // Log level
			IgnoreRecordNotFoundError: true,                    // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,                   // Disable color
		},
	)
	var db *gorm.DB
	var err error
	db, err = gorm.Open(sqlite.Open("file:"+databaseConfig.DbPath+"?cache=shared&mode=ro"), &gorm.Config{Logger: dbLogger, PrepareStmt: true})
	if err != nil {
		log.Fatalf("error in connecting to database %v", err)
		return nil, err
	}
	return db, nil
}
