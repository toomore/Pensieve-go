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
