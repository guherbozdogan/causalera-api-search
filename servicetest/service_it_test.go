package servicetest

import (
	"context"
	"fmt"
	. "github.com/guherbozdogan/causalera-api-search/main"
	. "github.com/guherbozdogan/causalera-api-search/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("It testing search rest api", func() {

	Describe("test Search GET method", func() {
		Context("with normal parameters ", func() {
			var service = It("what to write here then:)", func() {
				//Ω(obictRecordIO).NotTo(Equal(nil))

				errc := make(chan error)
				RunServices(errc)

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
