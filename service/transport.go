package maven

//import "github.com/gorilla/mux"
import (
	"context"
	"encoding/json"
	//"github.com/dgrijalva/jwt-go/request"
	//"github.com/guherbozdogan/kit/endpoint"
	//httptransport "github.com/guherbozdogan/kit/transport/http"
	////"io/ioutil"
	"github.com/guherbozdogan/kit/log"
	"net/http"
	//	"net/url"
	//	"strconv"
	//"strings"
	//	"fmt"
	//	"net/url"
	//	"io/ioutil"
	//	"io/ioutil"
	//	"bytes"
	//	"context"
	//	"encoding/json"
	"errors"
	//	"io/ioutil"

	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/guherbozdogan/kit/tracing/opentracing"
	httptransport "github.com/guherbozdogan/kit/transport/http"
)

type errorWrapper struct {
	Error string `json:"error"`
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	code := http.StatusInternalServerError
	msg := err.Error()

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errorWrapper{Error: msg})
}

func errorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

func encodeSimpleSearchAPIResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeSimpleSearchAPIReq(_ context.Context, r *http.Request) (interface{}, error) {

	var req SearchAPIRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func MakeHTTPHandler(endpoints Endpoints, tracer stdopentracing.Tracer, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorLogger(logger),
	}
	m := http.NewServeMux()
	m.Handle("/search/libs", httptransport.NewServer(
		endpoints.SimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactIDEndPoint,
		decodeSimpleSearchAPIReq,
		encodeSimpleSearchAPIResponse,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "SimpleSearchAPICall",
			logger)))...,
	))
	return m
}
