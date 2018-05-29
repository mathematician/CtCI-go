package chapter2

func (ll *DoublyLinkedList) GetFromTail(k int) int {
	current := ll.head

	for i := 0; i < k; i++ {
		current = current.next
	}

	kthNode := ll.head

	for current != nil {
		current = current.next
		kthNode = kthNode.next
	}

	return kthNode.value
}
