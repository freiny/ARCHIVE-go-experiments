package main

import "fmt"

func main() {
	l := list{}.new()
	l.insert(1)
	l.insert(2)
	l.insert(3)

	l.iter()
	l.iter()
	l.iter()
	l.iter()
	l.iter()
	l.iter()
	l.iter()
	l.iter()
	l.iter()
	l.iter()
	l.iter()
	l.iter()
	// fmt.Println(l.head, l.head.next, l.head.next.next)
}

type list struct {
	head *node
	tail *node
	size int
	iter fnBool
}

func (l list) new() list {
	l.iter = l.titer()
	return l
}

func (l *list) insert(data int) {
	newNode := node{data, nil}
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
	} else {
		l.tail.next = &newNode
		l.tail = &newNode
	}
}

func (l *list) delete() {

}

type fnBool func() bool

func (l *list) titer() fnBool {
	i := 0
	return func() bool {
		if i < 4 {
			fmt.Println("hello", i)
			i++
			return true
		} else {
			i = 0
			return false
		}
	}

	// i := 0
	// return fn
}

type node struct {
	data int
	next *node
}
