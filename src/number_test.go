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
    }
    for _, value := range intNumbers {
        assert.True(t, IsNumber(value), fmt.Sprintf("%s should be integer", value))
    }
}
