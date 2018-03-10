package number

import (
    "regexp"
    "strings"
)

var matchFunc []func(string) (bool, error)

func IsNumber(s string) bool {
    s = strings.TrimSpace(s)
    if matchFunc == nil {
        matchFunc = []func(string) (bool, error){
            IsInteger,
            IsFloat,
            IsExponentialFunction,
        }
    }

    for _, f := range matchFunc {
        is_matched, _ := f(s)

        if is_matched {
            return true
        }
    }

    return false
}

func IsInteger(s string) (bool, error) {
    pattern := "^(-|\\+|[0-9])[0-9]*$"

    return regexp.MatchString(pattern, s)
}

func IsFloat(s string) (bool, error) {
    pattern1 := "^((-|\\+)|[0-9]+)?\\.[0-9]+$"
    pattern2 := "^(-|\\+)?[0-9]+\\.[0-9]*$"

    isMatched, err := regexp.MatchString(pattern1, s)
    if !isMatched {
        return regexp.MatchString(pattern2, s)
    }

    return isMatched, err
}

func IsExponentialFunction(s string) (bool, error) {
    pattern := "^-?((([0-9]+\\.?)|(\\.[0-9]+))e[0-9]+)$"

    return regexp.MatchString(pattern, s)
}
