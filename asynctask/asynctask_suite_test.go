package asynctask_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAsynctask(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Asynctask Suite")
}
