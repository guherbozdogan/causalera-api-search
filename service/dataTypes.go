package maven

import (
//"bytes"
//"encoding/json"
//"fmt"
)

//
//func (e SearchAPIError) Headers() http.Header {
//
//	//return e.errorJsonBytes, nil
//
//}

type StopResponse struct {
}
type SearchAPIRequest struct {
	Id        string `json:"id,omitempty"`
	Keyword   string `json:"keyword,omitempty"`
	App       string `json:"app,omitempty"`
	Version   string `json:"version,omitempty"`
	StartId   string `json:"startId,omitempty"`
	Direction string `json:"direction,omitempty"`
	MetaData  struct {
		RowCount int `json:"rowCount,omitempty"`
	} `json:"metaData,omitempty"`
}
type SearchAPIResponseDoc struct {
	Id       string `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	RepoType string `json:"repoType,omitempty"`
	Params   struct {
		GroupId    string `json:"groupId,omitempty"`
		ArtifactId string `json:"artifactId,omitempty"`
	} `json:"params,omitempty"`

	ImgURI      string `json:"imgURI,omitempty"`
	Description string `json:"description,omitempty"`
}
type SearchAPIResponse struct {
	Id        string `json:"id,omitempty"`
	App       string `json:"app,omitempty"`
	Version   string `json:"version,omitempty"`
	StartId   string `json:"startId,omitempty"`
	Direction string `json:"direction,omitempty"`
	MetaData  struct {
		RowCount int `json:"rowCount,omitempty"`
	} `json:"metaData,omitempty"`

	ResultSet []SearchAPIResponseDoc `json:"resultSet,omitempty"`
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
