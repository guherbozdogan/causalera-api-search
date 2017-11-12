package maven

import "errors"

var SystemCode = "100"
var SubsystemCode = "002"
var StrJSONParseErrorDesc = "JSON Parse Error"
var StrInternalServerErrorDesc = "Internal Server Error"

var StrJSONParseErrorCode = "001"
var StrInternalServerErrorCode = "002"

var ErrJSONParseError = errors.New(StrJSONParseErrorDesc)
var ErrInternalServerError = errors.New(StrInternalServerErrorDesc)
