package use_cases_test

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rafaph/notte-auth/test/factories"
	"github.com/rafaph/notte-auth/test/services/doubles"
)

var _ = Describe("application/use_cases/check_use_case", func() {
	It("should return an user Id when token is valid", func() {
		// given
		token := faker.Word()
		sut := factories.MakeCheckUseCase()
		useCase := sut.UseCase
		tokenVerifier := sut.TokenVerifier.(*doubles.MockTokenVerifier)
		// when
		output, err := useCase.Execute(token)
		// then
		Expect(err).To(BeNil())
		Expect(output.Id).To(Equal(tokenVerifier.Returneduser.Id))
	})

	It("should return an error when token is invalid", func() {
		// given
		token := faker.Word()
		tokenVerifier := doubles.NewMockTokenVerifier(nil, fmt.Errorf("invalid token"))
		sut := factories.MakeCheckUseCase(&factories.CheckUseCaseTypes{TokenVerifier: tokenVerifier})
		useCase := sut.UseCase
		// when
		output, err := useCase.Execute(token)
		// then
		Expect(output).To(BeNil())
		Expect(err).ToNot(BeNil())
	})
})
