package tree

import "testing"

func TestBinaryTree(t *testing.T) {
	// **测试二叉树结构如下**
	//            1
	//       2         3
	//    4    5     6     7
	//  8
	//     9

	// **前序遍历结果**
	// 1 -> 2 -> 4 -> 8 -> 9 -> 5 -> 3 -> 6 -> 7

	// **中序遍历结果**
	// 8 -> 9 -> 4 -> 2 -> 5 -> 1 -> 6 -> 3 -> 7

	// **后序遍历结果**
	// 9 -> 8 -> 4 -> 5 -> 2 -> 6 -> 7 -> 3 -> 1

	// **层次遍历结果**
	// 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9
	var root = NewNode(1)
	root.LChild = NewNode(2)
	root.RChild = NewNode(3)
	root.LChild.LChild = NewNode(4)
	root.LChild.RChild = NewNode(5)
	root.RChild.LChild = NewNode(6)
	root.RChild.RChild = NewNode(7)
	root.LChild.LChild.LChild = NewNode(8)
	root.LChild.LChild.LChild.RChild = NewNode(9)
	// PreOrder(root)
	// PreOrderV1(root)
	// MidOrder(root)
	// MidOrderV1(root)
	// PostOrder(root)
	//PostOrderV1(root)
	LevelTravel(root)
}
