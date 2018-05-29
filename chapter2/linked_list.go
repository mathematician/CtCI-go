package chapter2

type DoublyLinkedList struct {
	head  *node
	tail  *node
	count int
}

type node struct {
	value int
	next  *node
	prev  *node
}

//GetLinkedList, GetLinkedListFromValues, Insert, insertNode, getNode, Slice, Get, Len
func GetLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil, 0}
}

func GetLinkedListFromValues(values []int) *DoublyLinkedList {
	ll := GetLinkedList()
	for _, val := range values {
		ll.Insert(val)
	}

	return ll
}

func (ll *DoublyLinkedList) Get(index int) int {
	return ll.GetNode(index).value
}

func (ll *DoublyLinkedList) GetNode(index int) *node {
	node := ll.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (ll *DoublyLinkedList) Insert(val int) {
	newNode := &node{value: val}
	ll.InsertNode(newNode)
}

func (ll *DoublyLinkedList) InsertNode(newNode *node) {
	ll.count++
	if ll.tail == nil {
		ll.head, ll.tail = newNode, newNode
	} else {
		ll.tail.next = newNode
		newNode.prev = ll.tail
		ll.tail = newNode
	}
}

func (ll *DoublyLinkedList) Slice() []int {
	node := ll.head
	slice := []int{node.value}
	for node.next != nil {
		node = node.next
		slice = append(slice, node.value)
	}

	return slice
}
