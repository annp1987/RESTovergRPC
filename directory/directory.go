package directory

import "net/url"

// ExtractQuery extract a query string into a map {'name': 'AAA', last_name: 'BBB'}
func ExtractQuery(query string) url.Values {
	result, err := url.ParseQuery(query)
	if err != nil {
		panic(err)
	}
	return result
}
