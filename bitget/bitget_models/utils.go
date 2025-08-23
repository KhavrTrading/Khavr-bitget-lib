package bitget_models

import (
	"strconv"
	"strings"
)

// StringOrNumber can handle JSON values like "20" or 20 and store them as a string.
type StringOrNumber string

func (sn *StringOrNumber) UnmarshalJSON(b []byte) error {
	s := strings.TrimSpace(string(b))

	// If the JSON value is quoted, remove the quotes.
	if len(s) > 1 && s[0] == '"' && s[len(s)-1] == '"' {
		unquoted, err := strconv.Unquote(s)
		if err != nil {
			return err
		}
		*sn = StringOrNumber(unquoted)
	} else {
		// It's likely a raw number. Convert it to a string.
		*sn = StringOrNumber(s)
	}
	return nil
}

// Float64 is a helper method to parse the internal string as a float64.
func (sn StringOrNumber) Float64() (float64, error) {
	return strconv.ParseFloat(string(sn), 64)
}
