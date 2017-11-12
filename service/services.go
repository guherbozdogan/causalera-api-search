package maven

// This file contains the Service definition, and a basic service
// implementation. It also includes service middlewares.

import (
	"context"
	//	"errors"
	//"time"
	//
	mavenCli "github.com/guherbozdogan/causalera-api-search/client/maven"
	"github.com/guherbozdogan/kit/log"

	//"log"
	"time"
	//	"fmt"
	"errors"
	"github.com/guherbozdogan/causalera-api-search/common"
	"strconv"
)

// Service describes a service that adds things together.
type SearchService interface {
	SimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactID(
		ctx context.Context, req SearchAPIRequest) (interface{}, error)
}

type searchServiceImp struct {
	endPoints mavenCli.Endpoints
}

func NewSearchServiceImp(endPoints mavenCli.Endpoints) SearchService {
	tmp := searchServiceImp{endPoints: endPoints}
	return tmp

}

var systemErr = errors.New("System Error")

func (s searchServiceImp) SimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactID(
	ctx context.Context, req SearchAPIRequest) (interface{}, error) {

	rowCount := -1 //meaning there is no row limit

	rowCount = req.metaData.rowCount
	inp := mavenCli.SimpleSearchReturnAllVersionsofFullySpecifiedGroupIdAndArtifactIDRequest{
		SearchKeyWord: req.keyword, RowCount: rowCount,
	}

	//add circuit breaker and rate limiter and open tracing later
	resp, err := s.endPoints.SimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactID(context.Background(),
		inp)

	if err != nil {
		errTmp, isOk := err.(common.APIError)
		if isOk {
			errTmp.Id = req.id
			return nil, errTmp
		}
		return nil, err
	}

	response := resp.(mavenCli.SimpleSearchReturnAllVersionsofFullySpecifiedGroupIdAndArtifactIDResponse)

	var searchAPIResponse SearchAPIResponse

	searchAPIResponse.app = req.app
	searchAPIResponse.direction = req.direction
	searchAPIResponse.id = req.id
	searchAPIResponse.startId = req.startId
	//searchAPIResponse.resultSet[0].

	searchAPIResponse.metaData.rowCount = len(response.Response.Docs)

	searchAPIResponse.resultSet = make([]SearchAPIResponseDoc, searchAPIResponse.metaData.rowCount, searchAPIResponse.metaData.rowCount)
	for i, el := range response.Response.Docs {
		var doc SearchAPIResponseDoc
		//fallback image to maven //change to your webserver at sometime:)
		doc.imgURI = "https://upload.wikimedia.org/wikipedia/commons/thumb/0/0b/Maven_logo.svg/320px-Maven_logo.svg.png"

		doc.id = strconv.Itoa(i)
		doc.params.artifactId = el.A
		doc.params.groupId = el.G
		doc.repoType = "maven" //make it const
		doc.description = ""
		doc.title = el.A
		searchAPIResponse.resultSet[i] = doc

	}
	return searchAPIResponse, nil
}

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(SearchService) SearchService

type serviceLoggingMiddleware struct {
	logger log.Logger
	next   SearchService
}

// ServiceLoggingMiddleware returns a service middleware that logs the
// parameters and result of each method invocation.
func ServiceLoggingMiddleware(logger log.Logger) Middleware {
	return func(next SearchService) SearchService {
		return serviceLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}
func (mw serviceLoggingMiddleware) SimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactID(
	ctx context.Context, req SearchAPIRequest) (interface{}, error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "search",
			//	"a", a, "b", b, "result", v, "error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.SimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactID(ctx,
		req)

}

//func (mw serviceLoggingMiddleware) Concat(ctx context.Context, a, b string) (v string, err error) {
//	defer func(begin time.Time) {
//		mw.logger.Log(
//			"method", "Concat",
//			"a", a, "b", b, "result", v, "error", err,
//			"took", time.Since(begin),
//		)
//	}(time.Now())
//	return mw.next.Concat(ctx, a, b)
//}

// Business-domain errors like these may be served in two ways: returned
// directly by endpoints, or bundled into the response struct. Both methods can
// be made to work, but errors returned directly by endpoints are counted by
// middlewares that check errors, like circuit breakers.
//
// If you don't want that behavior -- and you probably don't -- then it's better
// to bundle errors into the response struct.

var (
//// ErrTwoZeroes is an arbitrary business rule for the Add method.
//ErrTwoZeroes = errors.New("can't sum two zeroes")
//
//// ErrIntOverflow protects the Add method. We've decided that this error
//// indicates a misbehaving service and should count against e.g. circuit
//// breakers. So, we return it directly in endpoints, to illustrate the
//// difference. In a real service, this probably wouldn't be the case.
//ErrIntOverflow = errors.New("integer overflow")
//
//// ErrMaxSizeExceeded protects the Concat method.
//ErrMaxSizeExceeded = errors.New("result exceeds maximum size")
)

// These annoying helper functions are required to translate Go error types to
// and from strings, which is the type we use in our IDLs to represent errors.
// There is special casing to treat empty strings as nil errors.

func str2err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

// NewBasicService returns a naïve, stateless implementation of Service.
//func NewBasicService() Service {
//	return basicService{}
//}
//
