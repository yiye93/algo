package tree

import (
	"fmt"

	"github.com/yiye93/algo/queue"
)

// 二叉查找树节点定义
type BinarySearchNode struct {
	Value int               //元素值
	Nums  int               // 值重复次数
	Left  *BinarySearchNode //左子树
	Right *BinarySearchNode //右子树
}

type BinarySearchTree struct {
	Root *BinarySearchNode //根节点
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// 插入操作
func (tree *BinarySearchTree) Insert(value int) {
	if tree.Root == nil {
		tree.Root = &BinarySearchNode{Value: value, Nums: 1}
		return
	}
	// 步骤1: 找到要插入值的父节点
	currentNode := tree.Root
	for currentNode != nil {
		// 小于子树根节点的值
		if value < currentNode.Value && currentNode.Left != nil {
			currentNode = currentNode.Left
			// 大于子树根节点的值
		} else if value > currentNode.Value && currentNode.Right != nil {
			currentNode = currentNode.Right
			// 等于子树根节点的值
		} else {
			break
		}
	}
	// 步骤2: 把插入值插入到步骤1中找到的父节点的左子树或者右子树或者只增加重复次数
	if value < currentNode.Value {
		currentNode.Left = &BinarySearchNode{Value: value, Nums: 1}
	} else if value > currentNode.Value {
		currentNode.Right = &BinarySearchNode{Value: value, Nums: 1}
	} else {
		currentNode.Nums += 1
	}
	return
}

// 二叉查找树打印
// 为了直观感受二叉查找树的构造结构,我们实现一个简单的打印功能
func (tree *BinarySearchTree) Print() {
	if tree.Root == nil {
		fmt.Printf("tree is empty!\n")
		return
	}
	// 采用层次遍历
	// 此处搞了个容量限制...
	q := queue.NewQueueOnLinklist(10000)
	// 头元素入队
	q.EnQueue(tree.Root)
	// 队列非空
	for !q.IsEmpty() {
		n := q.DeQueue().(*BinarySearchNode)
		// 左子树非空,左子树乖乖入队
		if n.Left != nil {
			q.EnQueue(n.Left)
			fmt.Printf("[value/nums/leftNode] is [%v/%v/%v]\n", n.Value, n.Nums, n.Left.Value)
		}
		// 右子树非空,左子树乖乖入队
		if n.Right != nil {
			q.EnQueue(n.Right)
			fmt.Printf("[value/nums/rightNode] is [%v/%v/%v]\n", n.Value, n.Nums, n.Right.Value)
		}
		if n.Left == nil && n.Right == nil {
			fmt.Printf("[value/nums/isLeafNode] is [%v/%v]\n", n.Value, n.Nums)
		}
	}
}

// 二叉查找树的搜索
func (tree *BinarySearchTree) Search(value int) *BinarySearchNode {
	if tree.Root == nil {
		return nil
	}
	currentNode := tree.Root
	for currentNode != nil {
		// 小于子树根节点的值
		if value < currentNode.Value {
			currentNode = currentNode.Left
			// 大于子树根节点的值
		} else if value > currentNode.Value {
			currentNode = currentNode.Right
			// 找到该值
		} else {
			return currentNode
		}
	}
	return nil
}

// 查找当前节点和父节点
func (tree *BinarySearchTree) SearchAndParent(value int) (*BinarySearchNode, *BinarySearchNode) {
	if tree.Root == nil {
		return nil, nil
	}
	currentNode := tree.Root
	var parentNode *BinarySearchNode = nil
	for currentNode != nil {
		// 小于子树根节点的值
		if value < currentNode.Value {
			parentNode = currentNode
			currentNode = currentNode.Left
			// 大于子树根节点的值
		} else if value > currentNode.Value {
			parentNode = currentNode
			currentNode = currentNode.Right
			// 找到该值
		} else {
			return currentNode, parentNode
		}
	}
	return currentNode, parentNode
}

// 查询最小值
func (tree *BinarySearchTree) FindMin() *BinarySearchNode {
	if tree.Root == nil {
		return nil
	}
	currentNode := tree.Root
	// 一直往左子树遍历
	for currentNode != nil {
		if currentNode.Left != nil {
			currentNode = currentNode.Left
		} else {
			break
		}
	}
	return currentNode
}

// 查询最大值
func (tree *BinarySearchTree) FindMax() *BinarySearchNode {
	if tree.Root == nil {
		return nil
	}
	currentNode := tree.Root
	// 一直往右子树遍历
	for currentNode != nil {
		if currentNode.Right != nil {
			currentNode = currentNode.Right
		} else {
			break
		}
	}
	return currentNode
}

// 中序遍历实现排序
func (tree *BinarySearchTree) PreOrder() {
	if tree.Root == nil {
		return
	}
	tree.preOrder(tree.Root)
	return
}

// 中序遍历实现排序
func (tree *BinarySearchTree) preOrder(node *BinarySearchNode) {
	if node == nil {
		return
	}
	tree.preOrder(node.Left)
	fmt.Printf("travel value is:%v\n", node.Value)
	tree.preOrder(node.Right)
	return
}

// 删除操作
func (tree *BinarySearchTree) Delete(value int) {
	if tree.Root == nil {
		return
	}
	// 判断值是否存在
	currentNode, parentNode := tree.SearchAndParent(value)
	if currentNode == nil {
		return
	}
	// 依照三种情况分类处理
	// 情况1: 该节点即为叶子节点
	if currentNode.Left == nil && currentNode.Right == nil {
		// 特殊情况,只有一个根节点的情况,树置空即可
		if parentNode == nil {
			tree.Root = nil
			return
		}
		println("ggg")
		// 如果删除的是父节点的左儿子,父节点左孩子置空即直接删除
		if parentNode.Left != nil && currentNode.Value == value {
			parentNode.Left = nil
			return
		} else {
			// 删除的是父节点的右儿子,父节点右孩子置空即直接删除
			parentNode.Right = nil
			return
		}
	}
	//情况2: 存在左子树或者右子树
	// 左子树非空,右子树为空
	if currentNode.Left != nil && currentNode.Right == nil {
		// 删除的是父节点的左儿子,那么直接让当前节点的左儿子变为父节点的左儿子
		if parentNode.Left != nil && currentNode.Value == value {
			parentNode.Left = currentNode.Left
			return
		} else {
			// 删除的是父节点的右儿子,那么直接让当前节点的左儿子变为父节点的右儿子
			parentNode.Right = currentNode.Left
		}
		return
	}
	// 左子树为空,右子树非空
	if currentNode.Left == nil && currentNode.Right != nil {
		// 删除的是父节点的左儿子,那么直接让当前节点的儿子变为父节点的左儿子
		if parentNode.Left != nil && currentNode.Value == value {
			parentNode.Left = currentNode.Right
			return
		} else {
			// 删除的是父节点的右儿子,那么直接让当前节点的左儿子变为父节点的右儿子
			parentNode.Right = currentNode.Right
		}
		return
	}

	//情况3: 左子树、右子树都非空
	if currentNode.Left != nil && currentNode.Right != nil {
		// 删除的节点下有两个子树，因为右子树的值都比左子树大，那么用右子树中的最小元素来替换删除的节点。
		// 右子树的最小元素，只要一直往右子树的左边一直找一直找就可以找到，替换后二叉查找树的性质又满足了。
		minNode := currentNode.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		//特殊情况: 删除节点的右子树只有一个节点时,此时最小值即为该节点
		if currentNode.Right.Value == minNode.Value {
			currentNode.Value = minNode.Value
			currentNode.Nums = minNode.Nums
			currentNode.Right = nil
			return
		}
		// 删除最小值
		tree.Delete(minNode.Value)
		currentNode.Value = minNode.Value
		currentNode.Nums = minNode.Nums
		return
	}
	return
}
