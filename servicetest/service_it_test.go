package servicetest

import (
	"context"
	//	"fmt"
	//	"fmt"
	//	. "github.com/guherbozdogan/causalera-api-search/exec"
	. "github.com/guherbozdogan/causalera-api-search/service"
	util "github.com/guherbozdogan/causalera-api-search/servicetest/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//	"runtime"
	"fmt"
	"time"
)

func recoverFunc() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in f", r)
		cmd := fmt.Sprintf("cd $GOPATH/bin; cat outtmp | xargs kill -9 ")
		fmt.Println(cmd)
		execCmd(cmd)
		time.Sleep(1400 * time.Millisecond)
	}
}

var _ = Describe("It testing search rest api", func() {

	//BeforeEach(func() {
	//	server = ghttp.NewServer()
	//	client = sprockets.NewClient(server.URL())
	//})
	//
	//AfterEach(func() {
	//	//shut down the server between tests
	//	server.Close()
	//})

	Context("test Search GET method", func() {
		It("with normal json parameters and path", func() {
			//			runtime.GOMAXPROCS(2)

			defer recoverFunc()

			time.Sleep(1000 * time.Millisecond)
			tmpEndPoints, errMakeClients := util.MakeClientEndpoints("http://127.0.0.1:8083")
			Ω(errMakeClients).Should(BeNil())

			inp := SearchAPIRequest{
				Id:        "12132332",
				App:       "search.app",
				Version:   "0.1",
				Keyword:   "guice",
				StartId:   "0",
				Direction: "asc",
			}
			inp.MetaData.RowCount = 5

			resp, _ := tmpEndPoints.SearchAPIEndpoint(context.Background(),
				inp)

			dataStruct, isOk := resp.(SearchAPIResponse)

			Ω(isOk).Should(Equal(true))
			Ω(dataStruct.MetaData.RowCount).Should(Equal(5))
			Ω(len(dataStruct.ResultSet)).Should(Equal(5))

		})
		//mssing rowCount param
		It("with missing rowcount param", func() {

			defer recoverFunc()

			tmpEndPoints, errMakeClients := util.MakeClientEndpoints("http://127.0.0.1:8083")
			Ω(errMakeClients).Should(BeNil())

			inp := util.SearchAPIRequestWithBytes{
				Bytes: []byte(ReqWithMissingRowCountJSONStr),
			}

			resp, errR := tmpEndPoints.SearchAPIEndpointWithWrongJSON(context.Background(),
				inp)

			Ω(errR).Should(MatchError(util.ErrAPIError))
			Ω(resp).Should(BeNil())
		})

		It("with non-correct json", func() {

			defer recoverFunc()

			//time.Sleep(1000 * time.Millisecond)

			//			errc := make(chan error)
			//			go RunServices(errc)
			tmpEndPoints, errMakeClients := util.MakeClientEndpoints("http://127.0.0.1:8083")
			Ω(errMakeClients).Should(BeNil())

			inp := util.SearchAPIRequestWithBytes{
				Bytes: []byte(StrReqWithIncorrectJSON),
			}

			resp, errR := tmpEndPoints.SearchAPIEndpointWithWrongJSON(context.Background(),
				inp)

			Ω(errR).Should(MatchError(util.ErrAPIError))
			Ω(resp).Should(BeNil())
		})

	})
})
