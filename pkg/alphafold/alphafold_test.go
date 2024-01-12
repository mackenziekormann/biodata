package alphafold_test

import (
	"testing"

	"github.com/mackenziekormann/biodata/pkg/alphafold"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSetupAndExecution(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AlphaFold Fetching Tests")
}

const basicProteinID = "P00520"
const emptyString = ""

var _ = Describe("AlphaFold Tests", func() {

	Describe("Fetch Tests", func() {

		Specify("Fetch Tests: Basic Test", func() {
			data, err := alphafold.FetchAlphaFoldData(basicProteinID)

			Expect(err).To(BeNil())
			Expect(data).ToNot(BeEmpty())
		})

		Specify("Fetch Tests: Empty Input", func() {
			data, err := alphafold.FetchAlphaFoldData(emptyString)

			Expect(err).ToNot(BeNil())
			Expect(data).To(BeEmpty())
		})

	})

})
