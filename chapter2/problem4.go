package chapter2

func (ll *DoublyLinkedList) Partition(partition int) {
	higher := GetLinkedList()
	lower := GetLinkedList()
	current := ll.head

	for current != nil {
		next := current.next
		if current.value < partition {
			lower.Insert(current.value)
		} else {
			higher.Insert(current.value)
		}
		current = next
	}

	lower.tail.next = higher.head
	higher.head.prev = lower.tail
	ll.tail = higher.tail
	ll.head = lower.head
}
