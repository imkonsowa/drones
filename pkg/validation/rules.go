package validation

import (
	"fmt"
	"strings"
)

var validationMessages = map[string]string{
	"lte":   "exceeded the limit, max: :lte",
	"oneof": "not allowed value, one of: :oneof",
}

func GetRuleMessage(rule string, params map[string]string) string {
	if r, ok := validationMessages[rule]; ok {
		for key, val := range params {
			r = strings.ReplaceAll(r, fmt.Sprintf(":%s", key), val)
		}

		return r
	}

	return rule
}
