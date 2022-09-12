package services_test

import (
	"github.com/go-faker/faker/v4"
	"github.com/kataras/jwt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rafaph/notte-auth/infrastructure/services"
	"github.com/rafaph/notte-auth/test/factories"
	"time"
)

var _ = Describe("infrastructure/servies/jwt_token_verifier", func() {
	Context("Verify", func() {
		It("should verify a token successfully", func() {
			// given
			jwtConfig := factories.MakeJwtConfig()
			tokenGenerator := services.NewJwtTokenGenerator(jwtConfig)
			user := factories.MakeUser()
			token, _ := tokenGenerator.Generate(user)
			tokenVerifier := services.NewJwtTokenVerifier(jwtConfig)
			// when
			userFromToken, err := tokenVerifier.Verify(token)
			// then
			Expect(err).To(BeNil())
			Expect(userFromToken).ToNot(BeNil())
			Expect(userFromToken.Id).To(Equal(user.Id))
		})

		It("should return an error when public key is invalid", func() {
			// given
			jwtConfig := factories.MakeJwtConfig()
			tokenGenerator := services.NewJwtTokenGenerator(jwtConfig)
			user := factories.MakeUser()
			token, _ := tokenGenerator.Generate(user)
			jwtConfig.PublicKey = faker.Word()
			tokenVerifier := services.NewJwtTokenVerifier(jwtConfig)
			// when
			userFromToken, err := tokenVerifier.Verify(token)
			// then
			Expect(err).ToNot(BeNil())
			Expect(userFromToken).To(BeNil())
		})

		It("should return an error if cannot parse token claims", func() {
			// given
			jwtConfig := factories.MakeJwtConfig()
			privateKey, _ := jwt.ParsePrivateKeyEdDSA([]byte(jwtConfig.PrivateKey))
			maxAge := jwt.MaxAge(time.Duration(jwtConfig.ExpirationInMinutes) * time.Minute)
			claims := jwt.Map{"word": faker.Word()} // invalid claim
			tokenBytes, _ := jwt.Sign(jwt.EdDSA, privateKey, claims, maxAge)
			token := jwt.BytesToString(tokenBytes)
			// given
			tokenVerifier := services.NewJwtTokenVerifier(jwtConfig)
			// when
			userFromToken, err := tokenVerifier.Verify(token)
			// then
			Expect(err).ToNot(BeNil())
			Expect(userFromToken).To(BeNil())
		})
	})
})
