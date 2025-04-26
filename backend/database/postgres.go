package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB
}

func NewDatabase(host, port, user, password, dbName string) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{Connection: db}, nil
}

func (db *Database) Create(value interface{}) error {
	return db.Connection.Create(value).Error
}

func (db *Database) Migrate(models ...interface{}) error {
	return db.Connection.AutoMigrate(models...)
}
