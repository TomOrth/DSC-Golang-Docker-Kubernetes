package main

import "fmt"

// Node represents a node in the Singly Linked List
type Node struct {
	Value int
	Next  *Node
}

// SLL is an instance of a Singly Linked List
type SLL struct {
	Head *Node
}

// Add a node to the SLL
func (l *SLL) add(node *Node) {
	tmp := l.Head
	if tmp == nil {
		l.Head = node
	} else {
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = node
	}
}

// AddValue wil add an int to the
func (l *SLL) AddValue(val int) {
	node := &Node{
		Value: val,
		Next:  nil,
	}
	l.add(node)
}

// DeleteValue deletes an entry in the linked list if a Value exists
func (l *SLL) DeleteValue(val int) bool {
	tmp := l.Head
	if tmp == nil {
		return false
	}
	if tmp.Value == val {
		if tmp.Next == nil {
			l.Head = nil
			return true
		} else {
			l.Head = l.Head.Next
			return true
		}
	}
	if tmp.Next == nil {
		if tmp.Value == val {
			l.Head = nil
			return true
		}
	} else {
		Next := tmp.Next
		for Next != nil {
			if Next.Value == val {
				tmp.Next = Next.Next
				return true
			}
			tmp = Next
			Next = Next.Next
		}
	}
	return false
}

// Delete will remove a node from the list with a given value
func (l *SLL) Delete(node *Node) bool {
	return l.DeleteValue(node.Value)
}

// RetrieveValue will retrieve a node by a given value
func (l *SLL) RetrieveValue(val int) (*Node, bool) {
	tmp := l.Head
	for tmp != nil {
		if tmp.Value == val {
			return tmp, true
		}
		tmp = tmp.Next
	}
	return nil, false
}

// Print will print a given linked list
func (l *SLL) Print() {
	tmp := l.Head
	for i := 0; i < l.Len(); i++ {
		fmt.Printf("%d ", tmp.Value)
		tmp = tmp.Next
	}
	fmt.Println()
}

// Len returns the length of the list
func (l *SLL) Len() int {
	tmp := l.Head
	res := 0
	if tmp == nil {
		return 0
	}
	for tmp != nil {
		res++
		tmp = tmp.Next
	}
	return res
}

func main() {
	list := &SLL{
		&Node{
			10,
			nil,
		},
	}
	list.AddValue(11)
	list.AddValue(12)

	fmt.Printf("Len: %d\n", list.Len())
	fmt.Println("List:")
	list.Print()

	list.DeleteValue(11)
	fmt.Println("List:")
	list.Print()

	_, present := list.RetrieveValue(10)
	_, notPresent := list.RetrieveValue(100)
	fmt.Printf("Is 10 in the list? %t, Is 100? %t\n", present, notPresent)

}
