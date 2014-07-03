package stack

type Stack []interface{}

func (stack *Stack) Push(x interface{}) {
    *stack = append(*stack, x)
}

func (stack Stack) Len() int {
    return len(stack)
}

func (stack Stack) Cap() int {
    return cap(stack)
}

func (stack Stack) Top() interface{} {
    return stack[len(stack)-1]
}

func (stack *Stack) Pop() interface{} {
    last := stack.Top()
    new_stack := *stack
    *stack = new_stack[:len(*stack)-1]
    return last
}
