package maven

import (
	"context"
	"github.com/guherbozdogan/kit/endpoint"
	//httptransport "github.com/guherbozdogan/kit/transport/http"
	//"net/url"
	//"strings"
)

type Endpoints struct {
	SimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactIDEndPoint endpoint.Endpoint
}

func MakeSimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactIDEndPoint(s SearchService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		searchAPIReq := request.(SearchAPIRequest)

		v, err := s.SimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactID(ctx,
			searchAPIReq)

		return v, err
	}
}
