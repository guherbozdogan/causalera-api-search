package common

import (
	//"bytes"
	//"encoding/json"
	"fmt"
)

//import "encoding/json"
type APIError struct {
	Id               string
	StatusCodeVal    int
	Version          string
	ErrorCode        string
	ErrorDesc        string
	ErrorPage        string
	SupportErrorCode string
}

func (e APIError) Error() string {
	return fmt.Sprintf("%s - %s", e.ErrorCode, e.ErrorDesc)
}

//func (e APIError) MarshalJSON() ([]byte, error) {
//
//	var buf bytes.Buffer
//	//json.Marshal(e)
//	err := json.NewEncoder(&buf).Encode(e)
//	if err != nil {
//		return nil, err
//	}
//	return buf.Bytes(), nil
//
//	//fmt.Print("hey whts happnien")
//}

func (e APIError) StatusCode() int {
	return e.StatusCodeVal
}
