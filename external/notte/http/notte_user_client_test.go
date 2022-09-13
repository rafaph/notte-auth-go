package http_test

import (
	"encoding/json"
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rafaph/notte-auth/config"
	. "github.com/rafaph/notte-auth/external/notte/http"
	. "github.com/rafaph/notte-auth/infrastructure/repositories/http"
	"github.com/rafaph/notte-auth/test/helpers"
	"net/http"
	"testing"
)

func TestNotteHttpSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Notte Http Suite")
}

var _ = Describe("external/notte/http/notte_user_client", func() {
	It("should return an user response successfully", func() {
		// given
		fakeServer := helpers.NewFakeServer()
		userId := faker.UUIDHyphenated()
		body := map[string]string{"id": userId}
		responseBody, _ := json.Marshal(body)
		context := helpers.Context{
			StatusCode:   http.StatusOK,
			ResponseBody: string(responseBody),
		}
		fakeServer.When("POST", "/users/verify").Returns(context)

		fakeServer.Run(func(baseUrl string) {
			conf := &config.UserConfig{BaseUrl: baseUrl}
			client := NewNotteUserClient(conf)
			// when
			response, err := client.GetUser(GetUserRequest{Email: faker.Email(), Password: faker.Password()})
			// then
			Expect(response).ToNot(BeNil())
			Expect(response.Id).To(Equal(userId))
			Expect(err).To(BeNil())
		})
	})
})
