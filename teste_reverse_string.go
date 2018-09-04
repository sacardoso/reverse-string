package main

import (
	"testing"
)

// TestReverse blablabla TestReverse
func TestReverse(t *testing.T) {

	result, _ := reverse("word"), ""

	expected := "drow"

	if result != expected {
		t.Errorf("'%s' != '%s'", result, expected)
	}
}
