package repository

import (
	"app/internal/domain"
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(dns string) (*Repository, error) {
	db, err := gorm.Open(sqlite.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm open: %w", err)
	}

	err = db.AutoMigrate(
		&domain.User{},
		&domain.Button{},
	)
	if err != nil {
		return nil, fmt.Errorf("gorm automigrate: %w", err)
	}

	return &Repository{
		db: db,
	}, nil
}
