package chapter2

func SumLinkedLists(list1 *DoublyLinkedList, list2 *DoublyLinkedList) *DoublyLinkedList {
	node1 := list1.head
	node2 := list2.head
	sumList := GetLinkedList()
	extra := 0

	for node1 != nil && node2 != nil {
		sum := node1.value + node2.value + extra

		if sum >= 10 {
			sumList.Insert(sum - 10)
			extra = 1
		} else {
			extra = 0
			sumList.Insert(sum)
		}

		node1 = node1.next
		node2 = node2.next
	}
	return sumList
}

func ReverseLinkedList(ll *DoublyLinkedList) *DoublyLinkedList {
	vals := ll.Slice()
	reversedVals := []int{}

	for i := len(vals) - 1; i >= 0; i-- {
		reversedVals = append(reversedVals, vals[i])
	}

	return GetLinkedListFromValues(reversedVals)
}
