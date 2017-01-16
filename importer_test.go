package main

import (
	"testing"
)

func TestTest(t *testing.T) {

	if 1 != 1 {
		t.Error(
			"expected", 1,
			"got", 1,
		)
	}
}
