package main

import (
	"fmt"
	//"math"
)

type List struct {
	Prev *List
	Value float32
	Next *List
}


func (l *List) Insert(item float32) bool {
	if l == nil {
		l = &List { nil, item, nil }
		return true
	}

	if l.Value >= item {
		newl := &List{ l.Prev, item, l }
		if l.Prev != nil {
			l.Prev.Next = newl
		}
		l.Prev = newl
		return true
	}

	if l.Value < item {
		if l.Next == nil {
			newl := &List{ l, item, nil }
			l.Next = newl
		} else {
			return l.Next.Insert(item)
		}
	}

	return true
}

func (l *List) Remove(item float32) bool {
	if l != nil {
		if l.Value == item {
			l.Prev.Next = l.Next
			l.Next.Prev = l.Prev
			l = nil
		}
	}
	return false
}

func (l *List) Max() float32 {
	if l == nil {
		return 0
	}

	if l.Next == nil {
		return l.Value
	}

	return l.Next.Max()
}

func (l *List) Min() float32 {
	if l == nil {
		return 0
	}

	if l.Prev == nil {
		return l.Value
	}

	return l.Prev.Min()
}

func (l *List) String() string {

	str := ""
	node := l
	if l.Prev != nil {
		for node = l.Prev; node.Prev != nil; node = node.Prev {
		}
	}

	for ; node != nil; node = node.Next {
		str += fmt.Sprintf(" %f ", node.Value)
	}

	return str
}

func main() {

	l := List{Value: 6}
	l.Insert(70)
	l.Insert(40)
	l.Insert(2)
	l.Insert(18)
	l.Insert(190)
	l.Insert(44)
	l.Insert(40)

	fmt.Printf("%#v\n", l)
	fmt.Printf("%s\n", l.String())
	fmt.Printf("%s max\n", l.Max())
	fmt.Printf("%s min\n", l.Min())
	//fmt.Printf("%#s\n", blub)
}
