package bitget

import (
	"net/url"
	"sort"
	"strings"
)

func BuildGetParams(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}
	// Sort keys for consistency (optional, but good practice).
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	sb.WriteString("?")
	for i, k := range keys {
		if i > 0 {
			sb.WriteString("&")
		}
		sb.WriteString(url.QueryEscape(k))
		sb.WriteString("=")
		sb.WriteString(url.QueryEscape(params[k]))
	}
	return sb.String()
}
