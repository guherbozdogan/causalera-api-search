package servicetest

import (
	"context"
	"fmt"
	. "github.com/guherbozdogan/causalera-api-search/main"
	. "github.com/guherbozdogan/causalera-api-search/service"
	util "github.com/guherbozdogan/causalera-api-search/servicetest/util"
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

				tmpEndPoints, err := util.MakeClientEndpoints("http://search.maven.org")

				inp := SearchAPIRequest{
					Id:        "12132332",
					App:       "search.app",
					Version:   "0.1",
					Keyword:   "guice",
					StartId:   "0",
					Direction: "asc",
				}
				inp.MetaData.RowCount = 20

				resp, err := tmpEndPoints.SearchAPIEndpoint(context.Background(),
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
