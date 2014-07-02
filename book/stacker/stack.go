package stack

type Stack []interface{}

func (stack *Stack) Push(x interface{}) {
    *stack = append(*stack, x)
}
