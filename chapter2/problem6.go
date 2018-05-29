package chapter2

func IsPalindrome(ll *DoublyLinkedList) bool {
	current := ll.head
	m := map[int]int{}

	for current != nil {
		m[current.value] = (m[current.value] + 1) % 2
		current = current.next
	}

	sum := 0
	for _, val := range m {
		sum = sum + val
		if sum > 1 {
			return false
		}
	}

	return true
}
