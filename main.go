package main

import "fmt"

const (
	RED   = true
	BLACK = false
)

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Parent *Node
	Color  bool
}

type RedBlackTree struct {
	Root *Node
}

func (t *RedBlackTree) LeftRotate(x *Node) {
	y := x.Right
	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

func (t *RedBlackTree) RightRotate(x *Node) {
	y := x.Left
	x.Left = y.Right
	if y.Right != nil {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Right {
		x.Parent.Right = y
	} else {
		x.Parent.Left = y
	}
	y.Right = x
	x.Parent = y
}

func (t *RedBlackTree) InsertFixup(z *Node) {
	for z.Parent != nil && z.Parent.Color == RED {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			if y != nil && y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right {
					z = z.Parent
					t.LeftRotate(z)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				t.RightRotate(z.Parent.Parent)
			}
		} else {
			y := z.Parent.Parent.Left
			if y != nil && y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					t.RightRotate(z)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				t.LeftRotate(z.Parent.Parent)
			}
		}
	}
	t.Root.Color = BLACK
}

func (t *RedBlackTree) Insert(key int) {
	z := &Node{Key: key, Color: RED}
	var y *Node
	x := t.Root
	for x != nil {
		y = x
		if z.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	z.Parent = y
	if y == nil {
		t.Root = z
	} else if z.Key < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}
	t.InsertFixup(z)
}

func inOrderTraversal(node *Node) {
	if node != nil {
		inOrderTraversal(node.Left)
		fmt.Printf("Key: %d, Color: %v\n", node.Key, node.Color)
		inOrderTraversal(node.Right)
	}
}

func (t *RedBlackTree) PrintInOrder() {
	inOrderTraversal(t.Root)
}

func main() {
	tree := RedBlackTree{}
	// 插入示例数据
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(5)
	tree.Insert(15)

	// 打印红黑树节点信息
	tree.PrintInOrder()
}
