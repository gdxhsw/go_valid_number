package number

import (
    "regexp"
)

var numeric_patterns []string

func IsNumber(s string) bool {
    if numeric_patterns == nil {
        numeric_patterns = []string{
            "^(-|\\+|[0-9])[0-9]*$",
            "^(((-|\\+|[0-9]+)?\\.[0-9]+)|((-|\\+|[0-9]+)+\\.[0-9]?))$",
            "^([0-9]+e[0-9]+)$",
        }
    }

    for _, pattern := range numeric_patterns {
        is_matched, _ := regexp.MatchString(pattern, s)

        if is_matched {
            return true
        }
    }

    return false
}
