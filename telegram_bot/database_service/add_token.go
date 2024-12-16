package database_service

import (
	"context"
	"telegram_bot/models"
)

func (ds *DatabaseService) AddToken(telegramID int64, token *models.ApiToken) error {
	if token == nil {
		token = new(models.ApiToken)
	}

	query := "INSERT INTO telegram (token, telegram_id) VALUES ($1, $2)"
	_, errInsert := ds.pool.Exec(context.Background(), query, token.Value, telegramID)
	if errInsert != nil {
		return errInsert
	}

	return nil
}
