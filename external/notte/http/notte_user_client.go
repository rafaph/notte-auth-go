package http

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/rafaph/notte-auth/config"
	. "github.com/rafaph/notte-auth/infrastructure/repositories/http"
	"github.com/rafaph/notte-auth/lib/validator"
	"log"
	"net/url"
)

type NotteUserClient struct {
	config *config.UserConfig
}

func (n *NotteUserClient) GetUser(request GetUserRequest) (*GetUserResponse, error) {
	endpoint, _ := url.JoinPath(n.config.BaseUrl, "users", "verify")

	getUserResponse := GetUserResponse{}
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
		return nil, fmt.Errorf("fail to get resource")
	}

	err := validator.Validate(getUserResponse)

	if err != nil {
		return nil, err
	}

	return &getUserResponse, nil
}

func NewNotteUserClient(config *config.UserConfig) *NotteUserClient {
	return &NotteUserClient{config}
}
