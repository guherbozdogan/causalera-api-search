package maven

import (
	//"context"
	"github.com/guherbozdogan/kit/endpoint"
	//httptransport "github.com/guherbozdogan/kit/transport/http"
	//"net/url"
	//"strings"
)

type Endpoints struct {
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
}
