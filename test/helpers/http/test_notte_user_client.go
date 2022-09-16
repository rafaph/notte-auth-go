package http

import (
	"encoding/json"
	. "github.com/rafaph/notte-auth/config"
	. "github.com/rafaph/notte-auth/external/notte/http"
	"log"
)

const endpoint = "/users/verify"
const method = "POST"

type TestNotteUserClient struct {
	Response *MockResponse
}

func (t *TestNotteUserClient) Run(callback func(client *NotteUserClient)) {
	fakeServer := NewMockServer()

	body, isString := t.Response.Body.(string)
	if !isString {
		responseBody, err := json.Marshal(t.Response.Body)
		if err != nil {
			log.Panicln(err)
		}
		body = string(responseBody)
	}

	response := MockResponse{
		StatusCode: t.Response.StatusCode,
		Body:       body,
	}

	fakeServer.When(method, endpoint).Return(response)

	fakeServer.Run(func(baseUrl string) {
		conf := &UserConfig{BaseUrl: baseUrl}
		client := NewNotteUserClient(conf)
		callback(client)
	})
}

func NewTestNotteUserClient(response MockResponse) *TestNotteUserClient {
	return &TestNotteUserClient{
		Response: &response,
	}
}
