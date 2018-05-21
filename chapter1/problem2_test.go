package chapter1

import "testing"

func TestIsPermutation(t *testing.T) {
	isPermInput1 := "ABC"
	isPermInput2 := "BCA"
	notPermInput1 := "BCD"
	notPermInput2 := "DCA"

	if !isPermutation(isPermInput1, isPermInput2) {
		t.Errorf("failure, expected %s to be a permutation of %s", isPermInput1, isPermInput2)
	}

	if isPermutation(notPermInput1, notPermInput2) {
		t.Errorf("failure, expected %s to not be a permutation of %s", notPermInput1, notPermInput2)
	}
}
