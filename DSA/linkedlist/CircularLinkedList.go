package linkedlist

import (
	"fmt"
	"strings"
)

type CircularLinkedList struct {
	// 头结点
	head *LinkedListNode
	// 循环节点
	marker *LinkedListNode
	// 前置节点
	previous *LinkedListNode
	size     int
}

func newCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{}
}

func (c *CircularLinkedList) insert(value interface{}) {
	node := &LinkedListNode{Value: value}
	if c.isEmpty() {
		c.marker = node
		c.head = c.marker
		c.marker.Next = c.marker // 成环
	} else {
		c.marker.Next = node
		node.Next = c.head // 成环
		c.marker = node
	}
	c.size++
}

func (c *CircularLinkedList) setMarkerPosition(item interface{}) {
	if c.marker.Value == item {
		return
	}
	for c.marker.Value != item {
		c.marker = c.marker.Next
	}
}

func (c *CircularLinkedList) move() {
	if c.marker != nil {
		c.previous = c.marker
		c.marker = c.marker.Next
	}
}

func (c *CircularLinkedList) delete() {
	if c.isEmpty() {
		return
	}
	// 当前节点要删除了，故也需要移动 previous
	c.previous.Next = c.marker.Next
	c.marker = c.marker.Next
	c.size--
}

func (c *CircularLinkedList) isEmpty() bool {
	return c.marker == nil
}

func (c *CircularLinkedList) display() {
	if c.isEmpty() {
		fmt.Println("empty list.")
	}
	fmt.Printf("display circular Linked List: ")
	head := c.marker.Next
	for head != c.marker {
		fmt.Printf("%v->", head.Value)
		head = head.Next
	}
	fmt.Printf("%v\n", head.Value)
}

func (c *CircularLinkedList) string() string {
	if c.isEmpty() {
		fmt.Println("empty list.")
	}
	var output strings.Builder
	output.WriteString("display circular Linked List: ")
	head := c.marker.Next
	for head != c.marker {
		output.WriteString(fmt.Sprintf("%v->", head.Value))
		head = head.Next
	}
	output.WriteString(fmt.Sprintf("%v\n", head.Value))

	return output.String()
}
