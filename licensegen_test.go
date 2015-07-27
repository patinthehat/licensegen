package main

import "testing"

func TestSuccessful(t *testing.T) {
	fn := successful

	if fn(1) == true {
		t.Fail()
	}
	if fn(nil) != true {
		t.Fail()
	}
	if fn(false) == true {
		t.Fail()
	}
	if fn("abc") == true {
		t.Fail()
	}
}
