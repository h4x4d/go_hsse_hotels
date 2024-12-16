package api_service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"telegram_bot/models"
)

func (s *Service) Login(user *models.User) (bool, error) {
	login := new(models.Login)
	login.Login = *user.Login
	login.Password = *user.Password
	loginJSON, errJson := json.Marshal(login)
	if errJson != nil {
		return false, errJson
	}
	responseLogin, errLogin := http.Post(s.authUrl+"login", "application/json", bytes.NewBuffer(loginJSON))
	if errLogin != nil {
		return false, errLogin
	}
	defer responseLogin.Body.Close()

	if responseLogin.StatusCode != http.StatusOK {
		return false, nil
	}

	apiTokenJson, errRead := ioutil.ReadAll(responseLogin.Body)
	if errRead != nil {
		return false, errRead
	}

	apiToken := new(models.ApiToken)
	errDecode := json.Unmarshal(apiTokenJson, apiToken)
	if errDecode != nil {
		return false, errDecode
	}

	currToken, errToken := s.GetToken(user.TelegramID)
	if errToken != nil {
		return false, errToken
	}
	// wrong condition
	if currToken == nil {
		errAddToken := s.AddToken(user.TelegramID, apiToken)
		if errAddToken != nil {
			return false, errAddToken
		}
	} else {
		errSetToken := s.SetToken(user.TelegramID, apiToken)
		if errSetToken != nil {
			return false, errSetToken
		}
	}
	return true, nil
}
