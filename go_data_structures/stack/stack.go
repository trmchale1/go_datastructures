package stack

import "fmt"

type Stack struct {
    Head *Node
    int size = 0
}

type Node struct {
    Next *Node
    Data interface{}
}

func New() *Stack {
    emptyNode := &Node{
        Next: nil,
        Data: nil,
    }

    return &Stack{
        Head: emptyNode,
        size: 0,
    }
}

func (s *Stack) IsEmpty() bool {
        if s.Head == nil {
            return true
        }
        return false
}

func (s *Stack) Push(d interface{}) *Stack {
    nextNode := &Node{
        Next: nil,
        Data: d,
    }
    if s.Head.Data == nil {
        s.Head = nextNode
        s.size += 1
    } else {
       s.Head.Next = s.Head
       s.Head = nextNode
       s.size += 1
    }
    return s
}

func (s *Stack) Pop(v interface{}) *Stack {
    if s.IsEmpty() == true {
        fmt.Println("Stack is empty")
        return s
    }
    var node = s.Head
    s.Head = s.Head.Next
    s.size -= 1
    return s
}

func (ll *LinkedList) PrintAll() {
    var node = ll.Head
    for {
        fmt.Println(node.Data)
        if node.Next == nil {
            return
        }
        node = node.Next
    }
}

