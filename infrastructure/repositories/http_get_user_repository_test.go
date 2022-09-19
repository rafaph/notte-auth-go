package repositories_test

import (
	"errors"
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rafaph/notte-auth/infrastructure/clients"
	"github.com/rafaph/notte-auth/test/clients/doubles"
	"github.com/rafaph/notte-auth/test/factories"
	"testing"
)

func TestInfraestructureRepositories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "infrastructure/repositories tests")
}

var _ = Describe("infrastructure/repositories/http_get_user_repository", func() {
	It("should return an error when UserClient returns an error", func() {
		// given
		client := doubles.NewMockUserClient(nil, errors.New("error"))
		sut := factories.MakeHttpGetUserRepository(&factories.HttpGetUserRepositoryTypes{
			Client: client,
		})
		repository := sut.Repository

		// when
		input := factories.MakeGetUserInput()
		user, err := repository.GetUser(input)

		// then
		Expect(user).To(BeNil())
		Expect(err).ToNot(BeNil())
	})

	It("should return an error when UserClient returns an invalid id as response", func() {
		// given
		client := doubles.NewMockUserClient(&clients.GetUserResponse{Id: faker.Word()}, nil)
		sut := factories.MakeHttpGetUserRepository(&factories.HttpGetUserRepositoryTypes{
			Client: client,
		})
		repository := sut.Repository

		// when
		input := factories.MakeGetUserInput()
		user, err := repository.GetUser(input)

		// then
		Expect(user).To(BeNil())
		Expect(err).ToNot(BeNil())
	})

	It("should return an user when client returns", func() {
		// given
		sut := factories.MakeHttpGetUserRepository()
		repository := sut.Repository

		// when
		input := factories.MakeGetUserInput()
		user, err := repository.GetUser(input)

		// then
		Expect(user).ToNot(BeNil())
		Expect(err).To(BeNil())
	})
})
