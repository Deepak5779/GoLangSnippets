package testing

import "testing"

func Test_Hello(t *testing.T) {
	expected := "Hello Go"
	actual := hello()
	if actual != expected {
		t.Error("Test failed")
	}
}
