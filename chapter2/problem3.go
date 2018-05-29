package chapter2

func DeleteNode(node *node) {
	for node.next != nil {
		node.value = node.next.value
		if node.next.next == nil {
			node.next = nil
		} else {
			node = node.next
		}
	}
}
