package strings

import "testing"

func TestIsEmpty(t *testing.T) {
	if !IsEmpty("") {
		t.Log("incorrect IsEmpty")
		t.Fail()
	}
	if IsEmpty(" ") {
		t.Log("incorrect IsEmpty")
		t.Fail()
	}
}

func TestIsBlank(t *testing.T) {
	if !IsBlank("") {
		t.Log("incorrect IsBlank")
		t.Fail()
	}
	if !IsBlank(" ") {
		t.Log("incorrect IsBlank")
		t.Fail()
	}
	if !IsBlank("\t\f") {
		t.Log("incorrect IsBlank")
		t.Fail()
	}
	if IsBlank("\t\fa") {
		t.Log("incorrect IsBlank")
		t.Fail()
	}
}
