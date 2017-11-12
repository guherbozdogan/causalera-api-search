package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	service "github.com/guherbozdogan/causalera-api-search/service"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func encodeSearchAPIRequestWithBytes(ctx context.Context, r *http.Request, request interface{}) error {

	//tmpReq := request.(service.SearchAPIRequest)
	r.Method, r.URL.Path = "GET", "/search/libs"

	tmpReq := request.(SearchAPIRequestWithBytes)

	var buf bytes.Buffer = bytes.NewBuffer(tmpReq.bytes)

	r.Body = ioutil.NopCloser(&buf)
	return nil

}

func encodeSearchAPIReq(ctx context.Context, r *http.Request, request interface{}) error {

	//tmpReq := request.(service.SearchAPIRequest)
	r.Method, r.URL.Path = "GET", "/search/libs"

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil

}

func decodeSearchAPIReq(_ context.Context, resp *http.Response) (interface{}, error) {

	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	var response service.SearchAPIResponse
	bodyBytes, errReadRsp := ioutil.ReadAll(resp.Body)

	if errReadRsp == nil {

		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			log.Fatal(err)
			return response, err
		}

		bodyString := string(bodyBytes)

		fmt.Println(bodyString)
		return response, nil

	}

	return response, errReadRsp
}
