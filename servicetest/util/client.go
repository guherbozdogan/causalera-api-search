package util

import (
	httptransport "github.com/guherbozdogan/kit/transport/http"
	"net/url"
	"strings"
)

func MakeClientEndpoints(instance string) (Endpoints, error) {

	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}
	tgt.Path = ""

	options := []httptransport.ClientOption{}

	// Note that the request encoders need to modify the request URL, changing
	// the path and method. That's fine: we simply need to provide specific
	// encoders for each endpoint.

	return Endpoints{

		SearchAPIEndpoint: httptransport.NewClient("GET", tgt,
			encodeSearchAPIReq,
			decodeSearchAPIReq, options...).Endpoint(),
		SearchAPIEndpointWithWrongJSON: httptransport.NewClient("GET", tgt,
			encodeSearchAPIReqWithBytes,
			decodeSearchAPIReq, options...).Endpoint(),
		SearchAPIEndpointWithMissingJSONParam: httptransport.NewClient("GET", tgt,
			encodeSearchAPIReqWithBytes,
			decodeSearchAPIReq, options...).Endpoint(),
	}, nil
}
