package storage

import (
	"URL_Shortener/internal/config"
	"URL_Shortener/internal/models"
	"URL_Shortener/internal/utils"
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBStorage struct {
	db *gorm.DB
}

func NewDatabaseConnection(cfg *config.DB) (db *DBStorage, err error) {
	db = &DBStorage{}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Etc/UTC",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
	db.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.Migrate()

	return db, nil
}

func (db *DBStorage) Migrate() error {
	return db.db.AutoMigrate(&models.DBData{})
}

func (db *DBStorage) SaveData(ctx context.Context, link, token string) error {
	data := &models.DBData{
		Link:  link,
		Token: token,
	}
	return db.db.WithContext(ctx).Create(data).Error
}

func (db *DBStorage) GetLinkByToken(ctx context.Context, token string) (string, error) {
	data := &models.DBData{}
	err := db.db.WithContext(ctx).Where(&models.DBData{Token: token}).Take(&data).Error
	if err == gorm.ErrRecordNotFound {
		return "", utils.ErrNotFound
	}
	if err != nil {
		return "", err
	}

	return data.Link, nil
}

func (db *DBStorage) TryGetTokenByLink(ctx context.Context, link string) (string, error) {
	data := &models.DBData{}
	err := db.db.WithContext(ctx).Where(&models.DBData{Link: link}).Take(&data).Error
	if err == gorm.ErrRecordNotFound {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	return data.Token, nil
}
