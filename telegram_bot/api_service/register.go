package api_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"telegram_bot/models"
)

func (s *Service) Register(user *models.User) (*models.ApiToken, error) {
	register := new(models.Register)
	register.Login.Login = *user.Login
	register.Login.Password = *user.Password
	register.Role = *user.Role
	register.Email = *user.Mail
	register.TelegramID = user.TelegramID

	registerJSON, errJson := json.Marshal(register)
	if errJson != nil {
		return nil, errJson
	}

	fmt.Print(string(registerJSON))

	responseRegister, errRegister := http.Post(s.authUrl+"register/", "application/json", bytes.NewBuffer(registerJSON))
	if errRegister != nil {
		return nil, errRegister
	}
	defer responseRegister.Body.Close()

	if responseRegister.StatusCode != http.StatusOK {
		return nil, nil
	}

	apiTokenJson, errRead := ioutil.ReadAll(responseRegister.Body)
	if errRead != nil {
		return nil, errRead
	}

	apiToken := new(models.ApiToken)
	errDecode := json.Unmarshal(apiTokenJson, apiToken)
	if errDecode != nil {
		return nil, errDecode
	}
	return apiToken, nil
}
