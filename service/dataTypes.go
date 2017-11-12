package maven

import (
	"bytes"
	"encoding/json"
	"fmt"
)

//
//func (e SearchAPIError) Headers() http.Header {
//
//	//return e.errorJsonBytes, nil
//
//}

type SearchAPIRequest struct {
	id        string `json:"id,omitempty"`
	keyword   string `json:"keyword,omitempty"`
	app       string `json:"app,omitempty"`
	version   string `json:"version,omitempty"`
	startId   string `json:"startId,omitempty"`
	direction string `json:"direction,omitempty"`
	metaData  struct {
		rowCount int `json:"rowCount,omitempty"`
	} `json:"metaData,omitempty"`
}
type SearchAPIResponseDoc struct {
	id       string `json:"id,omitempty"`
	title    string `json:"title,omitempty"`
	repoType string `json:"repoType,omitempty"`
	params   struct {
		groupId    string `json:"groupId,omitempty"`
		artifactId string `json:"artifactId,omitempty"`
	} `json:"params,omitempty"`

	imgURI      string `json:"imgURI,omitempty"`
	description string `json:"description,omitempty"`
}
type SearchAPIResponse struct {
	id        string `json:"id,omitempty"`
	app       string `json:"app,omitempty"`
	version   string `json:"version,omitempty"`
	startId   string `json:"startId,omitempty"`
	direction string `json:"direction,omitempty"`
	metaData  struct {
		rowCount int `json:"rowCount,omitempty"`
	} `json:"metaData,omitempty"`

	resultSet []SearchAPIResponseDoc `json:"resultSet,omitempty"`
}

/*
type SimpleSearchReturnAllVersionsofFullySpecifiedGroupIdAndArtifactIDResponse struct {
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
*/
