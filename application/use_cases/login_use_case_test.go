package use_cases_test

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rafaph/notte-auth/test/factories"
	. "github.com/rafaph/notte-auth/test/repositories/doubles"
	. "github.com/rafaph/notte-auth/test/services/doubles"
	"testing"
)

func TestUseCases(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "application/use_cases tests")
}

var _ = Describe("application/use_cases/login_use_case", func() {
	It("should return a token if credentials are valid", func() {
		input := factories.MakeLoginUseCaseInput()
		sut := factories.MakeLoginUseCase()
		token := sut.ReturnedToken
		useCase := sut.UseCase

		output, err := useCase.Execute(input)

		Expect(err).To(BeNil())
		Expect(output.Token).To(Equal(token))
	})

	It("should return an error if input is invalid", func() {
		input := factories.MakeLoginUseCaseInput()
		input.Email = faker.Word()
		sut := factories.MakeLoginUseCase()
		useCase := sut.UseCase

		output, err := useCase.Execute(input)

		Expect(err).ToNot(BeNil())
		Expect(output).To(BeNil())
		validationErrors := err.(validator.ValidationErrors)
		Expect(validationErrors).To(HaveLen(1))
	})

	It("should return an error if getUserRepository fails", func() {
		getUserRepository := NewMockGetUserRepository(nil, fmt.Errorf("fail to get user"))
		sut := &factories.LoginUseCaseTypes{
			GetUserRepository: getUserRepository,
		}
		input := factories.MakeLoginUseCaseInput()
		sut = factories.MakeLoginUseCase(sut)
		useCase := sut.UseCase

		output, err := useCase.Execute(input)

		Expect(err).ToNot(BeNil())
		Expect(output).To(BeNil())
	})

	It("should return an error if tokenGenerator fails", func() {
		tokenGenerator := NewMockTokenGenerator("", fmt.Errorf("fail to generate token"))
		sut := &factories.LoginUseCaseTypes{
			TokenGenerator: tokenGenerator,
		}
		input := factories.MakeLoginUseCaseInput()
		sut = factories.MakeLoginUseCase(sut)
		useCase := sut.UseCase

		output, err := useCase.Execute(input)

		Expect(err).ToNot(BeNil())
		Expect(output).To(BeNil())
	})
})
