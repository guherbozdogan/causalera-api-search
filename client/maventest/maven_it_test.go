package maventest

import (
	"context"
	"fmt"
	. "github.com/guherbozdogan/causalera-api-search/client/maven"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("It testing with maven rest api", func() {

	Describe("dummy test", func() {
		Context("dummy test ", func() {
			It("dummy test", func() {
				//Ω(obictRecordIO).NotTo(Equal(nil))
				tmpEndPoints, err := MakeClientEndpoints("http://search.maven.org")

				inp := SimpleSearchReturnAllVersionsofFullySpecifiedGroupIdAndArtifactIDRequest{
					SearchKeyWord: "guice", RowCount: 20,
				}

				resp, err := tmpEndPoints.SimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactID(context.Background(),
					inp)

				if resp == nil {
					fmt.Println("hey")
				}
				//iii := resp.(SimpleSearchReturnAllVersionsofFullySpecifiedGroupIdAndArtifactIDResponse)
				//
				//Ω(iii).Should(Not(Equal(nil)))

				Ω(err).Should(Equal(nil))

			})

		})
	})
})
