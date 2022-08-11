package toolkit

import "testing"

func TestTools_Increment(t *testing.T) {
	var testTools Tools

	expected := 20

	testTools.Increment(10)
	result := testTools.Increment(10)

	if expected != result {
		t.Error("wrong result: result = ", result)
	}
}

func TestTools_Decrement(t *testing.T) {
	var testTools Tools

	expected := -20

	testTools.Decrement(10)
	result := testTools.Decrement(10)

	if expected != result {
		t.Error("wrong result: result = ", result)
	}
}
