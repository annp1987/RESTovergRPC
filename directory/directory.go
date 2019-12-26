package directory

import "encoding/json"

// ExtractQuery extract a query string into a map[string]interface{}
func ExtractQuery(query string) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(query), &result)
	return result
}
