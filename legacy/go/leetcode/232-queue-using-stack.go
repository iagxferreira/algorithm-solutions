package leetcode

type MyQueue struct {
	insertStack  []int
	consultStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		insertStack:  make([]int, 0),
		consultStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.insertStack = append(this.insertStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.consultStack) == 0 {
		for len(this.insertStack) > 0 {
			this.consultStack = append(this.consultStack, this.insertStack[len(this.insertStack)-1])
			this.insertStack = this.insertStack[:len(this.insertStack)-1]
		}
	}
	result := this.consultStack[len(this.consultStack)-1]
	this.consultStack = this.consultStack[:len(this.consultStack)-1]
	return result
}

func (this *MyQueue) Peek() int {
	if len(this.consultStack) == 0 {
		for len(this.insertStack) > 0 {
			this.consultStack = append(this.consultStack, this.insertStack[len(this.insertStack)-1])
			this.insertStack = this.insertStack[:len(this.insertStack)-1]
		}
	}

	return this.consultStack[len(this.consultStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.insertStack) == 0 && len(this.consultStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
