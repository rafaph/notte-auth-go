package http

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/rafaph/notte-auth/config"
	. "github.com/rafaph/notte-auth/infrastructure/repositories/http"
	"io"
	"log"
	"net/url"
)

type NotteUserClient struct {
	config *config.UserConfig
}

func end(body io.ReadCloser) {
	_ = body.Close()
}

func (n *NotteUserClient) GetUser(request GetUserRequest) (*GetUserResponse, error) {
	endpoint, _ := url.JoinPath(n.config.BaseUrl, "users", "verify")

	getUserResponse := GetUserResponse{}
	var apiError map[string]interface{}

	res, err := req.
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&getUserResponse).
		SetError(&apiError).
		SetBody(request).
		Post(endpoint)

	if err != nil {
		return nil, err
	}

	if res.IsError() {
		log.Println("Fail to get resource", apiError)
		return nil, fmt.Errorf("fail to get resource")
	}

	return &getUserResponse, nil
}

func NewNotteUserClient(config *config.UserConfig) *NotteUserClient {
	return &NotteUserClient{config}
}
