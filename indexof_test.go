package indexof

import "testing"

func TestStringsFound(t *testing.T) {

	var v = "foo"
	var a = []string{"bar", "foo", "bazz"}

	var i = IndexOf(v, a)
	if i != 1 {
		t.Error("not found")
	}
}

func TestStringsNotFound(t *testing.T) {

	var v = "quux"
	var a = []string{"bar", "foo", "bazz"}

	var i = IndexOf(v, a)
	if i != -1 {
		t.Error("found")
	}
}

func TestIntsFound(t *testing.T) {

	var v = 1
	var a = []int{0, 1, 2}

	var i = IndexOf(v, a)
	if i != 1 {
		t.Error("not found")
	}
}

func TestIntsNotFound(t *testing.T) {

	var v = 100
	var a = []int{0, 1, 2}

	var i = IndexOf(v, a)
	if i != -1 {
		t.Error("found")
	}
}