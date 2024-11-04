package core

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDatabase interface {
	DatabaseConnect() error
	DatabasePing() error
	DatabaseDisconnect() error
}

type Database struct {
	Handler *gorm.DB
	Logger  *Log
}

func NewDatabase(logger *Log) *Database {
	db := &Database{
		Logger: logger,
	}

	if err := db.DatabaseConnect(); err != nil {
		db.Logger.Error(ErrConnectDatabase + err.Error())
	}

	if err := db.DatabasePing(); err != nil {
		db.Logger.Error(ErrPingDatabase + err.Error())
	}

	return db
}

func (db *Database) DatabaseConnect() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_SSL"))

	handler, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.Handler = handler
	return nil
}

func (db *Database) DatabasePing() error {
	newDb, _ := db.Handler.DB()

	if err := newDb.Ping(); err != nil {
		return err
	}

	db.Logger.Info(InfoPingDatabase)
	return nil
}

func (db *Database) ModelMigrate(model ...interface{}) error {
	if err := db.Handler.AutoMigrate(model...); err != nil {
		return err
	}
	return nil
}
