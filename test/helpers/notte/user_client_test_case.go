package notte

import (
	"encoding/json"
	. "github.com/rafaph/notte-auth/config"
	. "github.com/rafaph/notte-auth/external/notte"
	. "github.com/rafaph/notte-auth/test/helpers/http"
	"log"
)

const endpoint = "/users/verify"
const method = "POST"

type UserClientTestCase struct {
	Response *MockResponse
}

func (t *UserClientTestCase) Run(callback func(client *UserClient)) {
	server := NewMockServer()

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

	server.When(method, endpoint).Return(response)

	server.Run(func(baseUrl string) {
		conf := &UserConfig{BaseUrl: baseUrl}
		client := NewUserClient(conf)
		callback(client)
	})
}

func NewUserClientTestCase(response MockResponse) *UserClientTestCase {
	return &UserClientTestCase{
		Response: &response,
	}
}
