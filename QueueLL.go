package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	back *Node
	size int
}

func (queue *Queue) Enqueue(value int)  {
	temp := &Node{value,nil}
	if queue.front == nil {
		queue.front = temp
		queue.back = temp
	}
	queue.back.next = temp
	queue.back = queue.back.next
	queue.size++
}

func (queue *Queue) Pop() (error,int) {
	if queue.front == nil {
		return errors.New("QUEUE IS EMPTY"),0
	}
	val := queue.front.data
	queue.front = queue.front.next
	queue.size--
	return nil,val
}

func main()  {
	var queue Queue
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(7)
	queue.Enqueue(67)

	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}
