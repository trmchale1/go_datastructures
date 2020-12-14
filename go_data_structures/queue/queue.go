/*
Package queue creates a queue data structure
*/
package queue

import "fmt"

// Queue : holds 2 nodes
type Queue struct {
	Head *Node
	Tail *Node
}

// Node : A Queue node
type Node struct {
	Next *Node
	Data interface{}
}

// New : Create a new Queue
func New() *Queue {
	emptyNode := &Node{
		Next: nil,
		Data: nil,
	}
	return &Queue {
		Head: emptyNode,
		Tail: emptyNode,
	}
}

// Check: empty?
func (q *Queue) IsEmpty() bool {
    if q.Head == nil {
        return true
    }
    return false
}

// Appends to the end of the list
func (q *Queue) Enqueue(d interface {}) *Queue {
    nextNode := &Node {
        Next: nil,
        Data: d,
    }
    if q.Head.Data == nil {
        q.Head = nextNode
    } else {
        q.Tail = nextNode
    }
    q.Tail = nextNode
    return q
}

func (q *Queue) Dequeue() *Queue {
    var node = q.Head
    q.Head = q.Head.Next
    
    return q
}

func (q *Queue) PrintAll() {
    var node = q.Head
    for {
        fmt.Println(node.Data)
        if node.Next == nil {
            return
        }
        node = node.Next
    }
}
