package main

import (
	"container/list"
	"fmt"
)

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Printf("%v\n", minStack.GetMin() == -3)
	minStack.Pop()
	fmt.Printf("%v\n", minStack.Top() == 0)
	fmt.Printf("%v\n", minStack.GetMin() == -2)
}

type StackElement struct {
	val      int
	smallest int
}

type MinStack struct {
	stack *list.List
}

func Constructor() MinStack {
	return MinStack{
		list.New(),
	}
}

func (this *MinStack) Push(val int) {
	el := StackElement{val, val}
	if this.stack.Len() > 0 && this.stack.Front().Value.(StackElement).smallest < val {
		el.smallest = this.stack.Front().Value.(StackElement).smallest
	}
	this.stack.PushFront(el)
}

func (this *MinStack) Pop() {
	topEl := this.stack.Front()
	this.stack.Remove(topEl)
}

func (this *MinStack) Top() int {
	return this.stack.Front().Value.(StackElement).val
}

func (this *MinStack) GetMin() int {
	return this.stack.Front().Value.(StackElement).smallest
}
