package number

import (
    "testing"
    "fmt"
    "github.com/stretchr/testify/assert"
)

func TestIsInteger(t *testing.T) {
    intNumbers := []string{
        "0",
        "1",
        "22",
        "234",
        "-234",
        "999999999",
        "-999999999",
        "+888",
    }
    for _, value := range intNumbers {
        assert.True(t, IsNumber(value), fmt.Sprintf("%s should be integer", value))
    }
}

func TestIsFloat(t *testing.T) {
    intNumbers := []string{
        ".1",
        "0.11",
        "2.123456",
        ".0091",
        "3.",
        "123.",
        "-.8",
    }
    for _, value := range intNumbers {
        assert.True(t, IsNumber(value), fmt.Sprintf("%s should be float", value))
    }
}

func TestIsExponentialFunction(t *testing.T) {
    intNumbers := []string{
        "2e0",
        "3e100",
    }
    for _, value := range intNumbers {
        assert.True(t, IsNumber(value), fmt.Sprintf("%s should be number", value))
    }
}

func TestIsNotNumeric(t *testing.T) {
    notNumbers := []string {
        "e9",
    }
    for _, value := range notNumbers {
        assert.False(t, IsNumber(value), fmt.Sprintf("%s is not numeric", value))
    }
}
