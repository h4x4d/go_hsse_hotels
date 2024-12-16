package database_service

import (
	"context"
	"telegram_bot/models"
)

func (ds *DatabaseService) SetToken(telegramID int64, token *models.ApiToken) error {
	if token == nil {
		token = new(models.ApiToken)
	}

	query := "UPDATE telegram SET token = $1 WHERE telegram_id = $2"
	_, errUpdate := ds.pool.Exec(context.Background(), query, token.Value, telegramID)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
