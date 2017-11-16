package util

import (
	"github.com/guherbozdogan/kit/endpoint"
)

type Endpoints struct {
	SearchAPIEndpoint                     endpoint.Endpoint
	SearchAPIEndpointWithWrongJSON        endpoint.Endpoint
	SearchAPIEndpointWithMissingJSONParam endpoint.Endpoint
}
