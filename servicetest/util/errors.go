package util

import "errors"

var StrJSONParseErrorDesc = "JSON Parse Error"
var StrInternalServerErrorDesc = "Internal Server Error"
var StrAPIErrorDesc = "API Error"

var StrJSONParseErrorCode = "001"
var StrInternalServerErrorCode = "002"
var StrAPIErrorCode = "003"

var ErrJSONParseError = errors.New(StrJSONParseErrorDesc)
var ErrInternalServerError = errors.New(StrInternalServerErrorDesc)

var ErrAPIError = errors.New(StrAPIErrorDesc)
