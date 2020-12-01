package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
	size int
}

func (stack *Stack) Push(value int)  {
	stack.top = &Node{value,stack.top}
	stack.size++
}

func (stack *Stack) Pop() (error,int) {
	if stack.top == nil {
		return errors.New("STACK IS EMPTY"),0
	}
	value := stack.top.data
	stack.top = stack.top.next
	stack.size--
	return nil,value
}

func (stack *Stack) Top() (error,int) {
	if stack.top == nil {
		return errors.New("STACK IS EMPTY"),0
	}
	return nil,stack.top.data
}

func main()  {
	var stack Stack
	stack.Push(1)
	stack.Push(3)
	stack.Push(6)

	for {
		err, el := stack.Pop()
		if err != nil{
			fmt.Println(err)
			break
		}else{
			fmt.Println(el)
		}
	}
}
