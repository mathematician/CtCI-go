package chapter2

import "fmt"

func GetIntersection(list1 *DoublyLinkedList, list2 *DoublyLinkedList) *DoublyLinkedList {
	node1 := list1.head
	node2 := list2.head
	intersection := GetLinkedList()

	fmt.Println(node1.value)
	fmt.Println(node2.value)

	for node1 != nil {
		for node2 != nil {
			node2 = node2.next
			if node11 
			intersection.InsertNode(node1)
		}
		node1 = node1.next
	}

	return intersection
}
