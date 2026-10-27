package rules

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	fmt.Println("Test is running")
	if 1 != 1 {
		t.Error("Test failed")
	} else {
		fmt.Println("Test passed")
	}
}
