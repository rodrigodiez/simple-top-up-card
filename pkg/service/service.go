package service

import (
	"database/sql"
	"errors"
)

type CardService interface {
	TopUp(string, float32) error
}

type cardService struct {
	db *sql.DB
}

func (svc *cardService) TopUp(cardID string, amount float32) error {

	if amount <= 0 {
		return errors.New("Top-up amount must be greater than 0")
	}
}

func New(db *sql.DB) CardService {
	return &cardService{
		db: db,
	}
}
