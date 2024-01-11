package AlphaFold_test

import (
	"testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSetupAndExecution(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AlphaFold Fetching Tests")
}

const basicProteinID = "P00520"

Describe("Fetch Tests", func() {

	Context("When fetching data with a valid ID,", func() {
		It("should return valid protein data without error", func() {
			data, err := AlphaFold.FetchAlphaFoldData(basicProteinID)

			Expect(err).To(BeNil())
			Expect(data).ToNot(BeNil())
		})
	})
})