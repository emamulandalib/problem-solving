package main

import (
	"container/list"
	"fmt"
)

func TestLinkedLIst() {
	l := list.New()
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	fmt.Println(l)

	//for e := l.Front(); e != nil; e = e.Next() {
	//	fmt.Println(e.Value)
	//}

	el := l.Front()
	l.Remove(el)
	fmt.Println(l.Front().Value)
}

func main() {
	TestLinkedLIst()
}
