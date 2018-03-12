package number

import (
    "regexp"
    "strings"
)

var matchFunc []func(string) bool

func IsNumber(s string) bool {
    s = strings.TrimSpace(s)
    if matchFunc == nil {
        matchFunc = []func(string) bool {
            IsInteger,
            IsFloat,
            IsExponentialFunction,
        }
    }

    for _, f := range matchFunc {
        is_matched := f(s)

        if is_matched {
            return true
        }
    }

    return false
}

var matchInteger *regexp.Regexp

func IsInteger(s string) bool {
    if matchInteger == nil {
        pattern := "^(-|\\+|[0-9])[0-9]*$"
        matchInteger, _ = regexp.Compile(pattern)
    }

    return matchInteger.MatchString(s)
}

var matchFloat1 *regexp.Regexp
var matchFloat2 *regexp.Regexp

func IsFloat(s string) bool {
    if matchFloat1 == nil {
        pattern1 := "^((-|\\+)|[0-9]+)?\\.[0-9]+$"
        matchFloat1, _ = regexp.Compile(pattern1)
    }

    if matchFloat2 == nil {
        pattern2 := "^(-|\\+)?[0-9]+\\.[0-9]*$"
        matchFloat2, _ = regexp.Compile(pattern2)
    }

    isMatched := matchFloat1.MatchString(s)
    if !isMatched {
        return matchFloat2.MatchString(s)
    }

    return isMatched
}

var matchExponentialFunction *regexp.Regexp

func IsExponentialFunction(s string) bool {
    if matchExponentialFunction == nil {
        pattern := "^(-|\\+)?((([0-9]+\\.?)|([0-9]*\\.[0-9]+))e(-|\\+)?[0-9]+)$"
        matchExponentialFunction, _ = regexp.Compile(pattern)
    }

    return matchExponentialFunction.MatchString(s)
}
