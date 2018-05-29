package chapter2

func (ll *DoublyLinkedList) RemoveDuplicates() {
	current := ll.head

	for current != nil {
		runner := current
		for runner.next != nil {
			if current.value == runner.next.value {
				runner.next = runner.next.next
			} else {
				runner = runner.next
			}
		}
		current = current.next
	}
}
