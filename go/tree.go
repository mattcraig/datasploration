package main

import (
	"fmt"
	"math"
)

type Tree struct {
	Left *Tree
	Value float32
	Depth float64
	Right *Tree
}


func (t *Tree) Add(item float32) bool {
	rdepth := 0.0
	ldepth := 0.0

	if t.Value > item {
		if t.Left == nil {
			t.Left = new(Tree)
			t.Left.Value = item
		} else {
			t.Left.Add(item)
		}

		ldepth = t.Left.Depth
		if t.Right != nil {
			rdepth = t.Right.Depth
		}
	} else {
		if t.Right == nil {
			t.Right = new(Tree)
			t.Right.Value = item
		} else {
			t.Right.Add(item)
		}

		rdepth = t.Right.Depth
		if t.Left != nil {
			ldepth = t.Left.Depth
		}
	}

	t.Rebalance()

	maxdepth := math.Max(rdepth, ldepth)
	t.Depth = maxdepth + 1

	return true
}

func (t *Tree) Rebalance() {
	if t.Left == nil && t.Right == nil {
		return
	}
	if t.Left == nil || t.Right == nil && t.Depth == 1 {
		return
	}

	switch {
		case t.Left == nil && t.Depth > 1, t.Right != nil && t.Right.Depth > (t.Left.Depth+1) :
			// right is too deep, rebalance it
				rebalval := t.Right.Min()
				readdval := t.Value
				if t.Right.RemoveLeaf(rebalval) {
					t.Value = rebalval
					t.Add(readdval)
				}
		case t.Right == nil && t.Depth > 1, t.Left != nil && t.Left.Depth > (t.Right.Depth + 1) :
			// left is too deep, rebalance it
				rebalval := t.Left.Max()
				readdval := t.Value
				if t.Left.RemoveLeaf(rebalval) {
					t.Value = rebalval
					t.Add(readdval)
				}
	}
}

func (t *Tree) RemoveLeaf(item float32) bool {
	if t.Value == item && t.Left == nil && t.Right == nil {
		t = nil
		return true
	}
	return false
}

func (t *Tree) Max() float32 {
	if t.Right == nil {
		return t.Value
	}
	return t.Right.Max()
}

func (t *Tree) Min() float32 {
	if t.Left == nil {
		return t.Value
	}
	return t.Left.Min()
}

func (t *Tree) Preorder() []float32 {

	var left []float32
	var right []float32

	if t.Left != nil {
		left = t.Left.Preorder()
	}

	if t.Right != nil {
		right = t.Right.Preorder()
	}

	treeorder := make([]float32, 1)

	treeorder[0] = t.Value
	if len(left) > 0 {
		treeorder = append(treeorder, left...)
	}
	if len(right) > 0 {
		treeorder = append(treeorder, right...)
	}

	return treeorder
}

func (t *Tree) String() string {

	if t.Left != nil && t.Right != nil {
		return fmt.Sprintf("( %s ) %f ( %s )", t.Left, t.Value, t.Right)
	}
	if t.Left == nil && t.Right != nil {
		return fmt.Sprintf("() %f ( %s )", t.Value, t.Right)
	}
	if t.Right == nil && t.Left != nil {
		return fmt.Sprintf("( %s ) %f ()", t.Left, t.Value)
	}

	return fmt.Sprintf("%f", t.Value)
}

func main() {
	//fmt.Println("go trees")

	t := Tree{Value: 6}
	t.Add(9)
	t.Add(3)
	t.Add(29)
	t.Add(39)
	t.Add(8)
	t.Add(19)
	//glug := t.Max()
	//blub := t.Preorder()
	//fmt.Printf("%#v with max %f\n", t, glug)
	fmt.Printf("%s\n", t.String())
	//fmt.Printf("%#v left\n", t.Left)
	//fmt.Printf("%#v right\n", t.Right)
	//fmt.Printf("%#v right right\n", t.Right.Right)
	//fmt.Printf("%#s\n", blub)

	shee := Tree{Value: 7}
	fmt.Printf("%s\n", shee.String())
	//fmt.Printf("%f\n", 7.8)
}
