package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rafaph/notte-auth/config"
	. "github.com/rafaph/notte-auth/infrastructure/repositories/http"
	"io"
	"net/http"
	"net/url"
)

type NotteUserClient struct {
	config *config.UserConfig
}

func end(body io.ReadCloser) {
	_ = body.Close()
}

func (n *NotteUserClient) GetUser(request GetUserRequest) (*GetUserResponse, error) {
	requestJson, _ := json.Marshal(request)
	requestBody := bytes.NewBuffer(requestJson)
	endpoint, _ := url.JoinPath(n.config.BaseUrl, "users/verify")

	response, err := http.Post(endpoint, "application/json", requestBody)
	defer end(response.Body)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid credentials")
	}

	responseJson, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	getUserResponse := GetUserResponse{}
	err = json.Unmarshal(responseJson, &getUserResponse)

	if err != nil {
		return nil, err
	}

	return &getUserResponse, nil
}

func NewNotteUserClient(config *config.UserConfig) *NotteUserClient {
	return &NotteUserClient{config}
}
