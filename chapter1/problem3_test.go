package chapter1

import "testing"

func TestURLify(t *testing.T) {
	str := "Mr John Smith    "
	length := 13
	expectedUrl := "Mr%20John%20Smith"

	url := URLify(str, length)

	if url != expectedUrl {
		t.Errorf("failure, expected %s and got %s", expectedUrl, url)
	}
}
