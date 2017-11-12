package maven

import "errors"

var SystemCode = "100"
var SubsystemCode = "001"
var StrJSONParseErrorDesc = "JSON Parse Error"
var StrMAVENRESTAPIErrorDesc = "Maven Rest Error:"
var StrInternalServerErrorDesc = "Internal Server Error"

var StrJSONParseErrorCode = "001"
var StrMAVENRESTAPIErrorCode = "002"
var StrInternalServerErrorCode = "003"

var ErrJSONParseError = errors.New(StrJSONParseErrorDesc)
var ErrMAVENRESTAPIError = errors.New(StrMAVENRESTAPIErrorDesc)
var ErrInternalServerError = errors.New(StrInternalServerErrorDesc)
