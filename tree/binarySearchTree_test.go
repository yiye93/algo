package tree

import (
	"testing"
)

func TestBinarySearchTree1(t *testing.T) {
	// **测试二叉查找树结构如下**
	//            4
	//       3         7
	//    1         6     8
	//
	//      2     5
	tree := NewBinarySearchTree()
	//fmt.Printf("test Insert...\n")
	//tree.Print()
	//插入元素
	insertElems := []int{4, 3, 7, 1, 6, 8, 2, 5}
	for _, elem := range insertElems {
		tree.Insert(elem)
	}
	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(8)
	//插入元素后打印树的结构
	//tree.Print()
	// 查找
	// fmt.Printf("test Search...\n")
	// for i := 0; i <= 10; i++ {
	// 	if tree.Search(i) == nil {
	// 		fmt.Printf("search elem:%v not exist\n", i)
	// 	} else {
	// 		fmt.Printf("search elem:%v exist\n", i)
	// 	}
	// }
	//fmt.Printf("test SearchAndParent...\n")
	// for i := 0; i <= 10; i++ {
	// 	node, parentNode := tree.SearchAndParent(i)
	// 	if node == nil {
	// 		fmt.Printf("search elem:%v not exist\n", i)
	// 	} else {
	// 		if parentNode != nil {
	// 			fmt.Printf("search elem:%v exist,parentElem is:%v\n", i, parentNode.Value)
	// 		} else {
	// 			fmt.Printf("search elem:%v exist\n", i)
	// 		}
	// 	}
	// }
	// 查找最小值和最大值
	// fmt.Printf("test FindMin and FindMax...\n")
	// fmt.Printf("find [MIN/MAX] is:[%v/%v]\n", tree.FindMin().Value, tree.FindMax().Value)
	// 中序遍历
	// fmt.Printf("test PreOrder...\n")
	// tree.PreOrder()
	// 删除操作
	// fmt.Printf("test Delete...\n")
	// tree.Delete(4)
	// tree.Print()
}

func TestBinarySearchTree2(t *testing.T) {
	// **测试二叉查找树结构如下**
	//  1
	//     2
	//        3
	//           4
	//              5
	//                 6
	//                    7
	//                      8
	tree := NewBinarySearchTree()
	//tree.Print()
	//插入元素
	insertElems := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, elem := range insertElems {
		tree.Insert(elem)
	}
	tree.Insert(1)
	tree.Insert(7)
	tree.Insert(7)
	//插入元素后打印树的结构
	//tree.Print()
	// 查找最小值和最大值
	// fmt.Printf("test FindMin and FindMax...\n")
	// fmt.Printf("find [MIN/MAX] is:[%v/%v]\n", tree.FindMin().Value, tree.FindMax().Value)
	// 中序遍历
	tree.PreOrder()
}
