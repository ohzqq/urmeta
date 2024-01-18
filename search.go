package audible

import (
	"net/url"
)

var queryFields = []string{
	"keywords",
	"author",
	"narrator",
	"title",
}

func GetSearchParams(vals url.Values) url.Values {
	params := make(url.Values)
	for _, f := range queryFields {
		if v, ok := vals[f]; ok {
			params[f] = v
		}
	}
	return params
}
