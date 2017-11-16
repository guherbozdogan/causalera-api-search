package servicetest

var StrReqWithIncorrectJSON = `{		"Id" : "12dsljkwq9233809",
"App": "search.libs",
"Version": "v0.1",
"StartId": "347289fsdkjhdf",
"Direction": "up",
"Keyword": "b",
"MetaData: {
"RowCount": 100
}`

var ReqWithMissingRowCountJSONStr = `{		"Id" : "12dsljkwq9233809",
"App1": "search.libs",
"Version": "v0.1",
"StartId": "347289fsdkjhdf",
"Direction": "up",
"Keyword": "b",
"MetaData": {
"RowCount": 100
}`

var ReqWithMissingKeywordJSONStr = `{		"Id" : "12dsljkwq9233809",
"App1": "search.libs",
"Version": "v0.1",
"StartId": "347289fsdkjhdf",
"Direction": "up",
"MetaData": {
"RowCount": 100
}`
