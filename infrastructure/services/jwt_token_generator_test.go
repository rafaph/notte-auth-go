package services_test

import (
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rafaph/notte-auth/infrastructure/services"
	"github.com/rafaph/notte-auth/test/factories"
)

var _ = Describe("infrastructure/servies/jwt_token_generator", func() {
	Context("Generate", func() {
		It("should generate a token successfully", func() {
			// given
			jwtConfig := factories.MakeJwtConfig()
			tokenGenerator := services.NewJwtTokenGenerator(jwtConfig)
			user := factories.MakeUser()
			// when
			token, err := tokenGenerator.Generate(user)
			// then
			Expect(err).To(BeNil())
			Expect(token).ToNot(BeEmpty())
		})

		It("should return an error if private key is invalid", func() {
			// given
			jwtConfig := factories.MakeJwtConfig()
			jwtConfig.PrivateKey = faker.Word()
			tokenGenerator := services.NewJwtTokenGenerator(jwtConfig)
			user := factories.MakeUser()
			// when
			token, err := tokenGenerator.Generate(user)
			// then
			Expect(err).ToNot(BeNil())
			Expect(token).To(BeEmpty())
		})
	})
})
