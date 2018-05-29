package chapter2

import (
	"reflect"
	"testing"
)

func TestSumLinkedLists(t *testing.T) {
	vals1 := []int{7, 1, 6}
	vals2 := []int{5, 9, 2}
	expected := []int{2, 1, 9}

	list1 := GetLinkedListFromValues(vals1)
	list2 := GetLinkedListFromValues(vals2)

	sum := SumLinkedLists(list1, list2)
	actual := sum.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}
}

func TestReverseLinkedList(t *testing.T) {
	vals := []int{6, 1, 7}
	expected := []int{7, 1, 6}
	ll := GetLinkedListFromValues(vals)

	reversedList := ReverseLinkedList(ll)
	actual := reversedList.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}

}
