package vatsense

import (
	"fmt"
	"testing"
)

var apiKey = "***REMOVED***"

func Test(t *testing.T) {
	fmt.Println(ListAllCountries(apiKey))
}
