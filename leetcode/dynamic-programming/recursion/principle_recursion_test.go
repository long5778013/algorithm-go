package recursion

import (
	"strings"
	"testing"
	_ "testing"
)

func Test_ReverseString(t *testing.T) {
	var strArr = []byte("Gopher!")
	reverseString(strArr)
	print(string(strArr))

	if strings.EqualFold(string(strArr[3]), string('h')) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
