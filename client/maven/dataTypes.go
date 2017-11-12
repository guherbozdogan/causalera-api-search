package maven

import (
//"context"
//"github.com/guherbozdogan/kit/endpoint"
//httptransport "github.com/guherbozdogan/kit/transport/http"
//"net/url"
//"strings"
)

type ErrorResponse struct {
	StatusCode int
	ErrorMsg   string
}
type SimpleSearchReturnAllVersionsofFullySpecifiedGroupIdAndArtifactIDRequest struct {
	SearchKeyWord string
	RowCount      int
}

type SearchResponseParams struct {
	Spellcheck      string `json:"spellcheck"`
	Fl              string `json:"fl"`
	Sort            string `json:"sort"`
	Indent          string `json:"indent"`
	Q               string `json:"q"`
	Qf              string `json:"qf"`
	SpellcheckCount string `json:"spellcheck.count"`
	Wt              string `json:"wt"`
	RowCount        string `json:"rowCount"`
	Version         string `json:"version"`
	DefType         string `json:"defType"`
}

type SimpleSearchResponseHeader struct {
	Status int                  `json:"status"`
	QTime  int                  `json:"QTime"`
	Params SearchResponseParams `json:"params"`
}

type SimpleSearchResponseDoc struct {
	Id            string   `json:"id"`
	G             string   `json:"g"`
	A             string   `json:"a"`
	LatestVersion string   `json:"latestVersion"`
	RepositoryId  string   `json:"repositoryId"`
	P             string   `json:"p"`
	Timestamp     int64    `json:"timestamp"`
	VersionCount  int      `json:"versionCount"`
	Text          []string `json:"text"`
	Ec            []string `json:"ec"`
}

type SimpleSearchResponse struct {
	NumFound int                       `json:"numFound"`
	Start    int                       `json:"start"`
	Docs     []SimpleSearchResponseDoc `json:"docs"`
}

type SimpleSearchReturnAllVersionsofFullySpecifiedGroupIdAndArtifactIDResponse struct {
	ResponseHeader SimpleSearchResponseHeader `json:"responseHeader"`
	Response       SimpleSearchResponse       `json:"response"`
	Spellcheck     struct {
		Suggestions []string `json:"suggestions"`
	} `json:"spellcheck"`
}

type SimpleSearchReturnAllVersionsofFullySpecifiedGroupIdAndArtifactIDResponseV2 struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime,omitempty"`
		Params struct {
			Spellcheck      string `json:"spellcheck,omitempty"`
			Fl              string `json:"fl,omitempty"`
			Sort            string `json:"sort,omitempty"`
			Indent          string `json:"indent,omitempty"`
			Q               string `json:"q,omitempty"`
			Qf              string `json:"qf,omitempty"`
			RowCount        string `json:"rowCount,omitempty"`
			SpellcheckCount string `json:"spellcheck.count,omitempty"`
			Wt              string `json:"wt,omitempty"`
			Version         string `json:"version,omitempty"`
			DefType         string `json:"defType,omitempty"`
		} `json:"params,omitempty"`
	} `json:"responseHeader,omitempty"`
	response struct {
		NumFound int `json:"numFound,omitempty"`
		Start    int `json:"start,omitempty"`
		Docs     []struct {
			Id            string   `json:"id,omitempty"`
			G             string   `json:"g,omitempty"`
			A             string   `json:"a,omitempty"`
			LatestVersion string   `json:"latestVersion,omitempty"`
			RepositoryId  string   `json:"repositoryId,omitempty"`
			P             string   `json:"p,omitempty"`
			Timestamp     int64    `json:"timestamp,omitempty"`
			VersionCount  string   `json:"versionCount,omitempty"`
			Text          []string `json:"text,omitempty"`
			Ec            []string `json:"ec,omitempty"`
		} `json:"docs,omitempty"`

		Spellcheck struct {
			Suggestions []string `json:"suggestions,omitempty"`
		} `json:"spellcheck,omitempty"`
	} `json:"response,omitempty"`
}
