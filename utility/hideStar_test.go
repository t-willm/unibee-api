package utility

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	result := HideStar("1511X2456@qq.com") // 151***@qq.com
	fmt.Println(result)
	result = HideStar("13077881053") // 130****1053
	fmt.Println(result)
	result = HideStar("362201200005302565") // 36***15
	fmt.Println(result)
	result = HideStar("***REMOVED***")
	fmt.Println(result)
}
