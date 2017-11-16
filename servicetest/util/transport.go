package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	common "github.com/guherbozdogan/causalera-api-search/common"
	service "github.com/guherbozdogan/causalera-api-search/service"
	"io/ioutil"
	"log"
	"net/http"
	//"net/url"
	//	"strconv"
)

func encodeSearchAPIRequestWithBytes(ctx context.Context, r *http.Request, request interface{}) error {

	//tmpReq := request.(service.SearchAPIRequest)
	r.Method, r.URL.Path = "GET", "/search/libs"

	tmpReq := request.(SearchAPIRequestWithBytes)

	var buf *bytes.Buffer = bytes.NewBuffer(tmpReq.Bytes)

	r.Body = ioutil.NopCloser(buf)
	return nil

}
func encodeSearchAPIReqWithBytes(ctx context.Context, r *http.Request, request interface{}) error {

	//tmpReq := request.(service.SearchAPIRequest)
	r.Method, r.URL.Path = "GET", "/search/libs"

	req, isOk := request.(SearchAPIRequestWithBytes)

	if !isOk {
		return ErrInternalServerError
	}
	buf := bytes.NewBuffer(req.Bytes)

	r.Body = ioutil.NopCloser(buf)
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

func errorDecoder(resp *http.Response) error {

	if resp.StatusCode == 404 {

		return ErrHttp404Error
	}
	bodyBytes, errReadRsp := ioutil.ReadAll(resp.Body)
	var response common.APIError
	if errReadRsp == nil {
		bodyString := string(bodyBytes)

		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			log.Println(ErrJSONParseError)

			log.Println(bodyString)
			return ErrJSONParseError
		}

		log.Println(bodyString)
		return ErrAPIError
	}
	log.Println(ErrInternalServerError)
	return ErrInternalServerError

}
func decodeSearchAPIReq(_ context.Context, resp *http.Response) (interface{}, error) {

	if resp.StatusCode != http.StatusOK {
		return nil, errorDecoder(resp)
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	var response service.SearchAPIResponse
	bodyBytes, errReadRsp := ioutil.ReadAll(resp.Body)

	if errReadRsp == nil {

		bodyString := string(bodyBytes)

		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			log.Println(ErrJSONParseError)
			log.Println(bodyString)

			return nil, ErrJSONParseError
		}

		return response, nil

	}

	log.Println(ErrInternalServerError)
	return nil, ErrInternalServerError
}
