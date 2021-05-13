package commons

type SqlFilter map[string]interface{}

type FormatQuery struct {
	Query  string
	Params []interface{}
}

func (paramsMap SqlFilter) ToFormattedQuery() FormatQuery {
	formatQuery := FormatQuery{}
	formatQuery.Params = make([]interface{}, 0, len(paramsMap))

	for key, value := range paramsMap {
		formatQuery.Query = formatQuery.Query + " " + key
		formatQuery.Params = append(formatQuery.Params, value)
	}

	return formatQuery
}
