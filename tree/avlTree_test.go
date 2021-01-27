package tree

import (
	"testing"
)

func TestAvlTree(t *testing.T) {
	var elems = []int{4, 2, 6, 1, 3, 5, 8, 7, 9, 10}
	tree := NewAvlTree()
	// 插入元素
	for _, e := range elems {
		tree.Insert(e)
	}
	// 打印平衡二叉树
	println("tree Print...")
	tree.Print()
	println()
	// // 中序遍历元素
	// println("tree PreOrder...")
	// tree.PreOrder()
	// println()

	// 查询
	// node := tree.Search(2)
	// if node == nil {
	// 	t.Fatalf("node should be exist")
	// }

	// node = tree.Search(20)
	// if node != nil {
	// 	t.Fatalf("node should not be exist")
	// }
	// 查询最大值和最小值
	// minNode, maxNode := tree.FindMin(), tree.FindMax()
	// fmt.Printf("getMin:%v, getMax:%v\n", minNode.Value, maxNode.Value)
	// 节点删除
	tree.Delete(8)
}

// 测试LL情况
func TestAvlTreeLL(t *testing.T) {
	var elems = []int{10, 8, 11, 7, 9, 5}
	tree := NewAvlTree()
	// 插入元素
	for _, e := range elems {
		tree.Insert(e)
	}
	//上述元素插入5后,失衡,将做一次右旋
	//                       10
	//
	//              8                    11
	//
	//        7           9
	//
	//    5
	//                  ||失衡调整后
	//                        8
	//
	//                 7             10
	//
	//          5                9         11
	// 打印平衡二叉树
	println("tree Print...")
	tree.Print()
	println()
	// 中序遍历元素
	println("tree PreOrder...")
	tree.PreOrder()
	println()
}

// 测试RR情况
func TestAvlTreeRR(t *testing.T) {
	var elems = []int{6, 5, 8, 7, 9, 10}
	tree := NewAvlTree()
	// 插入元素
	for _, e := range elems {
		tree.Insert(e)
	}
	//上述元素插入10后,失衡,将做一次左旋
	//                       6
	//
	//              5                    8
	//
	//                             7            9
	//
	//                                              10
	//                  ||失衡调整后
	//                        8
	//
	//                 6             9
	//
	//          5           7              10
	// 打印平衡二叉树
	println("tree Print...")
	tree.Print()
	println()
	// 中序遍历元素
	println("tree PreOrder...")
	tree.PreOrder()
	println()
}

// 测试LR情况
func TestAvlTreeLR(t *testing.T) {
	var elems = []int{10, 6, 11, 5, 7, 9}
	tree := NewAvlTree()
	// 插入元素
	for _, e := range elems {
		tree.Insert(e)
	}
	//上述元素插入9后,失衡,将做一次右旋
	//                       10
	//
	//              6                    11
	//
	//        5           7
	//
	//                         9
	//                  ||失衡调整后
	//                        7
	//
	//                 6             10
	//
	//          5                9         11
	// 打印平衡二叉树
	println("tree Print...")
	tree.Print()
	println()
	// 中序遍历元素
	println("tree PreOrder...")
	tree.PreOrder()
	println()
}

// 测试RL情况
func TestAvlTreeRL(t *testing.T) {
	var elems = []int{6, 5, 12, 10, 15, 11}
	tree := NewAvlTree()
	// 插入元素
	for _, e := range elems {
		tree.Insert(e)
	}
	//上述元素插入9后,失衡,将做一次右旋
	//                       6
	//
	//              5                    12
	//
	//                               10           15
	//
	//                                   11
	//                  ||失衡调整后
	//                        10
	//
	//                 6             12
	//
	//          5                11         15
	// 打印平衡二叉树
	println("tree Print...")
	tree.Print()
	println()
	// 中序遍历元素
	println("tree PreOrder...")
	tree.PreOrder()
	println()
}
