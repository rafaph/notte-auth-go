package notte_test

import (
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/rafaph/notte-auth/external/notte"
	. "github.com/rafaph/notte-auth/test/factories"
	. "github.com/rafaph/notte-auth/test/helpers/http"
	"github.com/rafaph/notte-auth/test/helpers/notte"
	"net/http"
	"testing"
)

func TestExternalNotte(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "external/notte tests")
}

var _ = Describe("external/notte/http/notte_user_client", func() {
	It("should return an user response successfully", func() {
		userId := faker.UUIDHyphenated()
		response := MockResponse{
			StatusCode: http.StatusOK,
			Body:       map[string]string{"id": userId},
		}
		notte.NewUserClientTestCase(response).Run(func(client *UserClient) {
			// given
			request := MakeGetUserRequest()

			// when
			response, err := client.GetUser(request)

			// then
			Expect(response).ToNot(BeNil())
			Expect(response.Id).To(Equal(userId))
			Expect(err).To(BeNil())
		})
	})

	It("should return an error when response status is 400", func() {
		response := MockResponse{
			StatusCode: http.StatusBadRequest,
			Body:       map[string]string{},
		}
		notte.NewUserClientTestCase(response).Run(func(client *UserClient) {
			// given
			request := MakeGetUserRequest()

			// when
			response, err := client.GetUser(request)

			// then
			Expect(response).To(BeNil())
			Expect(err).ToNot(BeNil())
		})
	})

	It("should return an error when response is different from expected", func() {
		response := MockResponse{
			StatusCode: http.StatusOK,
			Body:       map[string]string{},
		}
		notte.NewUserClientTestCase(response).Run(func(client *UserClient) {
			// given
			request := MakeGetUserRequest()

			// when
			response, err := client.GetUser(request)

			// then
			Expect(response).To(BeNil())
			Expect(err).ToNot(BeNil())
		})
	})
})
