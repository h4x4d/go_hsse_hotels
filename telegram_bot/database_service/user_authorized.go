package database_service

func (ds *DatabaseService) UserAuthorized(telegramID int64) (bool, error) {
	token, errToken := ds.GetToken(telegramID)
	if errToken != nil {
		return false, errToken
	}
	if token == nil {
		return false, nil
	}
	return (token.Value != ""), nil
}
