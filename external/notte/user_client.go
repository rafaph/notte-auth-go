package notte

import (
	"errors"
	"github.com/imroc/req/v3"
	"github.com/rafaph/notte-auth/config"
	"github.com/rafaph/notte-auth/infrastructure/clients"
	"github.com/rafaph/notte-auth/lib/validator"
	"log"
	"net/url"
)

type UserClient struct {
	config *config.UserConfig
}

func (n *UserClient) GetUser(request clients.GetUserRequest) (*clients.GetUserResponse, error) {
	endpoint, _ := url.JoinPath(n.config.BaseUrl, "users", "verify")

	getUserResponse := clients.GetUserResponse{}
	var apiError map[string]interface{}

	res, _ := req.
		R().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetHeader("Accept", "application/json").
		SetResult(&getUserResponse).
		SetError(&apiError).
		SetBody(request).
		Post(endpoint)

	if res.IsError() {
		log.Println("Fail to get resource", apiError)
		return nil, errors.New("fail to get resource")
	}

	err := validator.Validate(getUserResponse)

	if err != nil {
		return nil, err
	}

	return &getUserResponse, nil
}

func NewUserClient(config *config.UserConfig) *UserClient {
	return &UserClient{config}
}
