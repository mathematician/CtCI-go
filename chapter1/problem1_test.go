package chapter1

import "testing"

func TestIsUniq(t *testing.T) {
	uniqStr := "abcdefg"
	notUniqStr := "abcdefcg"

	if !isUniq(uniqStr) {
		t.Errorf("failure, expected %s to have only unique characters", uniqStr)
	}

	if isUniq(notUniqStr) {
		t.Errorf("failure, expected %s to have only unique characters", notUniqStr)
	}
}
