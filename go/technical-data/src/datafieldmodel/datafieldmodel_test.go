package datafieldmodel_test

import (
	"bitbucket.sprinteins.com/users/trusz/repos/code-dojo/go/technical-data/src/datafieldmodel"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Datafieldmodel", func() {

	It("creates DF Model", func() {
		var dfModel datafieldModel
		dfModel = datafieldmodel.NewDatafieldModel("_", "_")

		// to satisfy error "declared and not used"
		dfModel.GetArkiID()
	})

	It("sets Arki ID", func() {
		const arkiID string = "arki1"

		dfModel := datafieldmodel.NewDatafieldModel(arkiID, "_")

		actualArkiID := dfModel.GetArkiID()

		Expect(actualArkiID).To(Equal(arkiID), "Could not set arki id")

	})

	It("sets SSP ID", func() {
		const sspID string = "sspid1"

		dfModel := datafieldmodel.NewDatafieldModel(sspID, "_")

		actualSspID := dfModel.GetArkiID()

		Expect(actualSspID).To(Equal(sspID), "Could not set ssp id")
	})

})

type datafieldModel interface {
	GetArkiID() string
	GetSspID() string
}

type attribute interface{}
