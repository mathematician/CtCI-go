package chapter1

import "testing"

func TestIsPalindromePermutation(t *testing.T) {
	trueInput := "Tact Coa"
	falseInput := "Tact Coae"

	if !isPalindromePermutation(trueInput) {
		t.Errorf("failure, expecting 'true' and got 'false' for input: %s", trueInput)
	}

	if isPalindromePermutation(falseInput) {
		t.Errorf("failure, expecting 'false' and got 'true' for input: %s", falseInput)
	}
}
