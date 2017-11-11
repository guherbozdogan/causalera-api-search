package json

import (
//"encoding/json"
//"errors"
//"fmt"
//"os"
//"time"
//"github.com/guherbozdogan/kit/log"
//"github.com/guherbozdogan/kit/metrics"
)

//Search API JSON data structs
type MetaData struct {
	rowCount int `json:"rowCount"`
}

type SearchRequestBody struct {
	id        string   `json:"id"`
	app       string   `json:"app"`
	version   string   `json:"version"`
	startId   string   `json:"startdId"`
	direction string   `json:"direction"`
	metaData  MetaData `json:"metaData"`
}
type ResultSetItemParams interface {
}
type ResultSetItemMavenParams struct {
	groupId    string `json:"groupId"`
	artifactId string `json:"artifactId"`
}

type ResultSetItem struct {
	id          string              `json:"id"`
	repoType    string              `json:"repoType"`
	params      ResultSetItemParams `json:"params"`
	imgURI      string              `json:"imgURI"`
	description string              `json:"description"`
}

type SearchResponseBody struct {
	id        string          `json:"id"`
	app       string          `json:"app"`
	version   string          `json:"version"`
	metaData  MetaData        `json:"metaData"`
	resultSet []ResultSetItem `json:"resultSet"`
}

type SearchErrorResponseBody struct {
	id               string `json:"id"`
	statusCode       string `json:"statusCode"`
	app              string `json:"app"`
	version          string `json:"version"`
	errorCode        string `json:"errorCode"`
	error            string `json:"error"`
	errorPage        string `json:"errorPage"`
	supportErrorCode string `json:"supportErrorCode"`
}
