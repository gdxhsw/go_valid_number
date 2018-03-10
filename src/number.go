package number

import (
    "regexp"
)

var matchFunc []func(string) (bool, error)

func IsNumber(s string) bool {
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
    pattern := "^(((-|[0-9]+)?\\.[0-9]+)|((-|[0-9]+)+\\.[0-9]?))$"

    return regexp.MatchString(pattern, s)
}

func IsExponentialFunction(s string) (bool, error) {
    pattern := "^([0-9]+e[0-9]+)$"

    return regexp.MatchString(pattern, s)
}
