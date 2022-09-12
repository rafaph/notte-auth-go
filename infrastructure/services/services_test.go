package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestInfrastructureServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Infrastructure Services Suite")
}
