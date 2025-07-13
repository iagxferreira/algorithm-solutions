type MyStack struct {
	items []int
	size  int
}

func Constructor() MyStack {
	return MyStack{
		items: make([]int, 0),
		size:  0,
	}
}

func (this *MyStack) Push(x int) {
	this.items = append(this.items, x)
	this.size++
}

func (this *MyStack) Pop() int {
	this.size--
	item := this.items[this.size]
	this.items = this.items[:this.size]
	return item
}

func (this *MyStack) Top() int {
	return this.items[this.size-1]
}

func (this *MyStack) Empty() bool {
	return this.size == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
