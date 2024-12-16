package database_service

import (
	"context"
	"telegram_bot/models"
)

func (ds *DatabaseService) GetToken(telegramID int64) (*models.ApiToken, error) {
	tokenRow, errGet := ds.pool.Query(context.Background(),
		"SELECT token FROM telegram WHERE telegram_id = $1", telegramID)
	if errGet != nil {
		return nil, errGet
	}

	defer tokenRow.Close()

	if !tokenRow.Next() {
		return nil, nil
	}

	apiToken := new(models.ApiToken)

	// scaning token
	errScan := tokenRow.Scan(&apiToken.Value)
	if errScan != nil {
		return nil, errScan
	}
	return apiToken, nil
}
