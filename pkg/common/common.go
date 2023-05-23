package common

import (
	"encoding/base64"
)

const (
	Get     = "GET"

	APIContextPattern = "/{context:(?:%s|)}%s"

	APIHealthStatusUrl = "aGVhbHRoL2VtcGVyb3I="

	ContextPath = "/"
)

func GetAPIHealthStatusUrl() string {
	return decode(APIHealthStatusUrl)
}

func decode(input string) (string) {
	result, err := base64.StdEncoding.DecodeString(input)
    if err != nil {
        panic(err)
    }
    
	return string(result)
}