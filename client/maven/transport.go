package maven

//import "github.com/gorilla/mux"
import (
	"context"
	"encoding/json"
	common "github.com/guherbozdogan/causalera-api-search/common"
	//"github.com/dgrijalva/jwt-go/request"
	//"github.com/guherbozdogan/kit/endpoint"
	//httptransport "github.com/guherbozdogan/kit/transport/http"
	////"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	//"strings"
	"fmt"
	//	"net/url"
	//	"io/ioutil"
	//	"io/ioutil"
	//	"errors"
	"io/ioutil"
	"log"
)

/*

SimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactID       endpoint.Endpoint
	SimpleSearchReturnAllVersionsofFullySpecifiedGroupIdAndArtifactID          endpoint.Endpoint
	SimpleSearchReturnLatestVersionsOfAllArtifactsWithFullySpecifiedGroupId    endpoint.Endpoint
	SimpleSearchReturnLatestVersionsOfAllArtifactsWithFullySpecifiedArtifactID endpoint.Endpoint

	DownloadArtifact endpoint.Endpoint

	AdvancedSearchReturnAllVersionsWithUnionOfFullySpecifiedFilters              endpoint.Endpoint
	AdvancedSearchReturnAllVersionsWithIntersectionOfFullySpecifiedFilters       endpoint.Endpoint
	AdvancedSearchReturnAllVersionsArtifactsContainingTheClassName               endpoint.Endpoint
	AdvancedSearchReturnAllVersionsArtifactsContainingTheFullyQualifiedClassName endpoint.Endpoint
	AdvancedSearchReturnAllVersionsArtifactsContainingTheFileWithSHA1CheckSum    endpoint.Endpoint

	SimpleSearchReturnAllVersionsContainingFullySpecifiedTags endpoint.Endpoint

*/

func encodeSimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactID(ctx context.Context, req *http.Request, request interface{}) error {

	tmpReq := request.(SimpleSearchReturnAllVersionsofFullySpecifiedGroupIdAndArtifactIDRequest)
	req.Method, req.URL.Path = "GET", "/solrsearch/select"

	var tmp = url.Values{}
	//q := req.URL.Query()
	tmp.Add("q", tmpReq.SearchKeyWord)
	i := strconv.Itoa(tmpReq.RowCount)
	tmp.Add("rowCount", i)
	tmp.Add("wt", "json")

	req.URL.RawQuery = tmp.Encode()

	fmt.Println(req.URL.String())

	return nil

}

func errorDecoder(resp *http.Response) error {
	//log things here later

	tmpErr := common.APIError{
		StatusCodeVal: resp.StatusCode,
		ErrorPage:     "", Version: common.StrVersion,
	}

	bodyBytes, errReadRsp := ioutil.ReadAll(resp.Body)
	if errReadRsp != nil {

		tmpErr.ErrorDesc = StrJSONParseErrorDesc
		tmpErr.ErrorCode = SystemCode + SubsystemCode + StrJSONParseErrorCode
		tmpErr.SupportErrorCode = SystemCode + SubsystemCode + StrInternalServerErrorCode

		return tmpErr
	}

	tmpErr.ErrorCode = SystemCode + SubsystemCode + StrMAVENRESTAPIErrorCode
	tmpErr.SupportErrorCode = SystemCode + SubsystemCode + StrMAVENRESTAPIErrorCode
	tmpErr.ErrorDesc = StrMAVENRESTAPIErrorDesc + string(bodyBytes)

	return tmpErr

}
func decodeSimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactID(_ context.Context, resp *http.Response) (interface{}, error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, errorDecoder(resp)
	}
	var response SimpleSearchReturnAllVersionsofFullySpecifiedGroupIdAndArtifactIDResponse
	bodyBytes, errReadRsp := ioutil.ReadAll(resp.Body)

	if errReadRsp == nil {

		if err := json.Unmarshal(bodyBytes, &response); err != nil {

			tmpErr := common.APIError{
				StatusCodeVal: 500,
				ErrorDesc:     StrJSONParseErrorDesc,
				ErrorCode:     SystemCode + SubsystemCode + StrInternalServerErrorCode,
				ErrorPage:     "", Version: common.StrVersion,
				SupportErrorCode: SystemCode + SubsystemCode + StrInternalServerErrorCode,
			}
			log.Fatal(err)
			return nil, tmpErr
		}

		bodyString := string(bodyBytes)

		fmt.Println(bodyString)
		return response, nil

	}

	tmpErr := common.APIError{
		StatusCodeVal: 500,
		ErrorDesc:     StrMAVENRESTAPIErrorDesc + errReadRsp.Error(),
		ErrorCode:     SystemCode + SubsystemCode + StrMAVENRESTAPIErrorCode,
		ErrorPage:     "", Version: common.StrVersion,
		SupportErrorCode: SystemCode + SubsystemCode + StrMAVENRESTAPIErrorCode,
	}

	return nil, tmpErr
}
