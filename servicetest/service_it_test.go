package servicetest

import (
	"context"
	"fmt"
	. "github.com/guherbozdogan/causalera-api-search/exec"
	. "github.com/guherbozdogan/causalera-api-search/service"
	util "github.com/guherbozdogan/causalera-api-search/servicetest/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("It testing search rest api", func() {

	Describe("test Search GET method", func() {
		Context("with normal parameters ", func() {
			It("what to write here then:)", func() {
				//Ω(obictRecordIO).NotTo(Equal(nil))

				var recChan chan []byte
				recChan = make(chan []byte, 1000)

				errc := make(chan error)
				go RunServices(errc)

				tmpEndPoints, err := util.MakeClientEndpoints("http://127.0.0.1:8082")

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
				var s []byte

				Eventually(recChan, 1099990000).Should(Receive(&s))

				Ω(err).Should(Equal(nil))

			})

		})
	})
})
