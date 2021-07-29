package stack

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Infix struct {
	value string
}

func newInfix(value string) *Infix {
	return &Infix{value: value}
}

/*
infix 表达式转换为机器已解析的 postfix

步骤：
	- 从左到右遍历 infix
	- 如果是操作数，加入到 postfix
	- 如果是 left parentheses `(`， push in stack
	- 如果是 operator(+, -, *, /)
		- pop all operators
		- push new token on stack
	- 如果是 right parentheses `)`
		- pop all operators from stack and append them to output string
		- pop the left parenthesis
	- 当所有 infix 遍历完后，pop all the elements from stack and append them to the output string
*/

// TODO 暂未考虑空格、非法字符问题
func (infix *Infix) convertToPostfix() string {
	var postfix strings.Builder

	cacheStack := NewStackBasedLinkedList(20)
	for _, elem := range []byte(infix.value) {
		if isDigit(elem) {
			postfix.WriteByte(elem)
		} else if elem == '(' {
			cacheStack.Push(elem)
		} else if isOperator(elem) {
			for !cacheStack.IsEmpty() && getPrecedence(cacheStack.Peek().(byte)) >= getPrecedence(elem) {
				postfix.WriteByte(cacheStack.Pop().(byte))
			}
			cacheStack.Push(elem)
		} else if elem == ')' {
			for cacheStack.Peek().(byte) != '(' {
				postfix.WriteByte(cacheStack.Pop().(byte))
			}
			cacheStack.Pop() // Opening bracket
		}
	}

	for !cacheStack.IsEmpty() {
		postfix.WriteByte(cacheStack.Pop().(byte))
	}

	fmt.Println("postfix: ", postfix.String())
	return postfix.String()
}

func isDigit(elem byte) bool {
	// if elem >= 'a' && elem <= 'z' {
	// 	return true
	// }
	if elem >= '0' && elem <= '9' {
		return true
	}
	return false
}

func isOperator(str byte) bool {
	operators := "+-*/"
	for _, op := range []byte(operators) {
		if str == op {
			return true
		}
	}
	return false
}

func getPrecedence(operator byte) int {
	switch operator {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '^':
		return 3
	default:
		return 0
	}
}

func postfixEvaluation(items string) int {
	stack := NewStackBasedLinkedList(20)
	var result int
	for _, item := range []byte(items) {
		if isDigit(item) {
			stack.Push(item)
		} else {
			operand1 := byteValueConvertToInt(stack.Pop())
			operand2 := byteValueConvertToInt(stack.Pop())

			switch item {
			case '+':
				result = operand1 + operand2
				stack.Push(result)
			case '-':
				result = operand2 - operand1
				stack.Push(result)
			case '*':
				result = operand1 * operand2
				stack.Push(result)
			case '/':
				// TODO 注意 / ^ - 时 operand 的位置
				result = operand2 / operand1
				stack.Push(result)
			case '^':
				result = int(math.Pow(float64(operand2), float64(operand1)))
				stack.Push(result)
			}
		}
	}

	return result
}

func byteValueConvertToInt(item interface{}) int {
	switch value := item.(type) {
	case byte:
		return byteToInt(value)
	case int:
		return value
	default:
		panic(fmt.Sprintf("unsupport type %T", value))
	}
}

func byteToInt(item byte) int {
	value, _ := strconv.Atoi(string(item))
	return value
}
