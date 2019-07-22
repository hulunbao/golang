package main

import (
	"fmt"
	"golang.org/x/tools/container/intsets" // 导入第三方库 go get golang.org/x/tools/container/intsets
)

// 定义结构体
type treeNode struct{
	value int
	left, right *treeNode
}

// 打印树
func (node treeNode) print()  {
	fmt.Print(node.value ," ")
}

// 设置值
func (node *treeNode) setValue(value int)  {
	if node == nil {
		fmt.Println("Setting value to nil " +
			"node. Ignored.")
		return
	}
	node.value = value
}


// 遍历数
func (node *treeNode) traverse()  {
	node.TraverseFunc(func(node *treeNode) {
		node.print()
	})
	fmt.Println()
}

func (node *treeNode) TraverseFunc(f func(node *treeNode)){
	if node == nil{
		return
	}

	node.left.TraverseFunc(f)
	f(node)
	node.right.TraverseFunc(f)
}

// 使用组合扩展已有类型   另一种方法为：定义别名
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func testSparse()  {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(100)
	s.Insert(100000)
	fmt.Println(s.Has(100))
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	root.right.left.setValue(4)

	root.traverse()

	testSparse()

	nodeCount := 0

	root.TraverseFunc(func(node *treeNode) {
		nodeCount++
	})
	fmt.Println("nodeCount:",nodeCount)


}
