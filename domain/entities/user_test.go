package entities_test

import (
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rafaph/notte-auth/domain/entities"
	"testing"
)

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}

var _ = Describe("domain/entities/user", func() {
	Context("NewUser", func() {
		It("should return an error when id is invalid", func() {
			// given/when
			user, err := entities.NewUser(faker.Word())
			// then
			Expect(user).To(BeNil())
			Expect(err).ToNot(BeNil())
		})

		It("should return an user when id is valid", func() {
			// given/when
			user, err := entities.NewUser(faker.UUIDHyphenated())
			// then
			Expect(user).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
	})
})
