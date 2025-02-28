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

func Gz() {
	// 初始欠款
	initialDebt := 5500.0

	// 已支付金额
	payments := []float64{
		3100, 200, 500, 400, 200, 800, 800, // 9月
		1000, 2500, 500, // 10月
		1000, 100, 300, 1000, 1100, 200, 200, 1100, // 11月
		1300, 200, 500, 450, 500, 2000, 1000, // 12月
		8000, 200, 500, 500, 200, 100, // 1月
		2000, 80, 200, 1300, 300, 800, // 2月
	}

	// 计算总已支付金额
	totalPayments := 0.0
	for _, payment := range payments {
		totalPayments += payment
	}

	fmt.Println(totalPayments)

	// 每月工资
	monthlySalary := 7500.0

	// 计算应得工资
	months := 5 // 9月到1月共5个月
	totalSalary := float64(months) * monthlySalary

	// 剩余工资
	remainingSalary := totalSalary - totalPayments

	repairSalary := float64(2500)

	fmt.Printf("总剩余工资: %.2f元\n", remainingSalary + initialDebt + repairSalary)
}

func main() {
	// tree := RedBlackTree{}
	// // 插入示例数据
	// tree.Insert(10)
	// tree.Insert(20)
	// tree.Insert(30)
	// tree.Insert(5)
	// tree.Insert(15)
	//
	// // 打印红黑树节点信息
	// tree.PrintInOrder()
	Gz()
}
