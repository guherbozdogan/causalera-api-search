package maven

import (
	//"context"
	//"github.com/guherbozdogan/kit/endpoint"
	httptransport "github.com/guherbozdogan/kit/transport/http"
	//"net/url"
	//"strings"
	"net/url"
	"strings"
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

		SimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactID: httptransport.NewClient("GET", tgt,
			encodeSimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactID,
			decodeSimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactID, options...).Endpoint(),

		SimpleSearchReturnAllVersionsofFullySpecifiedGroupIdAndArtifactID:          nil,
		SimpleSearchReturnLatestVersionsOfAllArtifactsWithFullySpecifiedGroupId:    nil,
		SimpleSearchReturnLatestVersionsOfAllArtifactsWithFullySpecifiedArtifactID: nil,

		DownloadArtifact: nil,

		AdvancedSearchReturnAllVersionsWithUnionOfFullySpecifiedFilters:              nil,
		AdvancedSearchReturnAllVersionsWithIntersectionOfFullySpecifiedFilters:       nil,
		AdvancedSearchReturnAllVersionsArtifactsContainingTheClassName:               nil,
		AdvancedSearchReturnAllVersionsArtifactsContainingTheFullyQualifiedClassName: nil,
		AdvancedSearchReturnAllVersionsArtifactsContainingTheFileWithSHA1CheckSum:    nil,

		SimpleSearchReturnAllVersionsContainingFullySpecifiedTags: nil,
	}, nil
}
