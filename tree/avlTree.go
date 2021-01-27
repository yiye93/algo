package tree

import (
	"fmt"

	"github.com/yiye93/algo/queue"
)

type AvlTreeNode struct {
	Value  int // 值
	Height int // 以该节点为根节点, 对应的树的高度, 便于计算平衡因子
	Nums   int // 出现的次数
	Left   *AvlTreeNode
	Right  *AvlTreeNode
}

type AvlTree struct {
	Root *AvlTreeNode //根节点,因不停的调整可能会变化
}

func NewAvlTree() *AvlTree {
	return new(AvlTree)
}

func (n *AvlTreeNode) GetHeight() int {
	return n.Height
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 获取树的平衡因子
// 左子树高度减去右子树高度
func (n *AvlTreeNode) GetBalanceFactor() int {
	var leftTreeHeight, rigthTreeHeight = 0, 0
	if n.Left != nil {
		leftTreeHeight = n.Left.GetHeight()
	}
	if n.Right != nil {
		rigthTreeHeight = n.Right.GetHeight()
	}
	return leftTreeHeight - rigthTreeHeight
}

// 更新树高度
// 树高=max(左子树高, 右子树高)+1
func (n *AvlTreeNode) UpdateHeight() {
	if n == nil {
		return
	}
	var leftTreeHeight, rigthTreeHeight = 0, 0
	if n.Left != nil {
		leftTreeHeight = n.Left.GetHeight()
	}
	if n.Right != nil {
		rigthTreeHeight = n.Right.GetHeight()
	}
	n.Height = Max(leftTreeHeight, rigthTreeHeight) + 1
	return
}

// 左旋调整步骤:
// 1. 失衡点的右孩子节点代替失衡点位置
// 2. 右孩子节点的左子树变为失衡点的右子树
// 3. 失衡点本身变为右孩子的左子树
func (n *AvlTreeNode) LeftRotation() *AvlTreeNode {
	newRootNode := n.Right
	n.Right = n.Right.Left
	newRootNode.Left = n
	// 更新失衡点的高度
	n.UpdateHeight()
	// 更新失衡点右孩子节点的高度
	n.Right.UpdateHeight()
	return newRootNode
}

// 右旋调整步骤:
// 1. 失衡点的左孩子节点代替失衡点
// 2. 左孩子节点的右子树变为失衡点的左子树
// 3. 失衡点本身变为左孩子的右子树
func (n *AvlTreeNode) RigthRotation() *AvlTreeNode {
	newRootNode := n.Left
	n.Left = n.Left.Right
	newRootNode.Right = n
	// 更新失衡点的高度
	n.UpdateHeight()
	// 更新失衡点左孩子节点的高度
	n.Left.UpdateHeight()
	return newRootNode
}

// 左旋然后右旋
// 调整策略:
// 1. 先绕失衡点的左孩子节点左旋做局部调整
// 2. 再绕失衡点右旋即可完成整体调整
func (n *AvlTreeNode) LeftRightRotation() *AvlTreeNode {
	n.Left = n.Left.LeftRotation()
	return n.RigthRotation()
}

// 右旋然后左旋
// 调整策略:
// 1. 先绕失衡点的右孩子节点右旋做局部调整
// 2. 再绕失衡点左旋即可完成整体调整
func (n *AvlTreeNode) RightLeftRotation() *AvlTreeNode {
	n.Right = n.Right.RigthRotation()
	return n.LeftRotation()
}

func (tree *AvlTree) Insert(value int) {
	// 因插入新的元素,可能造成树失衡调整,而使根节点变更
	tree.Root = tree.Root.Insert(value)
	return
}

// 处理节点树高度问题
func (n *AvlTreeNode) HandleBF(value int) *AvlTreeNode {
	// 平衡因子(左子树和右子树高度之差)
	factor := n.GetBalanceFactor()
	newNode := new(AvlTreeNode)
	// 左子树的高度变高了，导致左子树-右子树的高度从1变成了2。
	if factor == 2 {
		if value < n.Left.Value {
			// 表示在左子树上插上左儿子导致失衡，需要单右旋：
			newNode = n.RigthRotation()
		} else {
			//表示在左子树上插上右儿子导致失衡，先左后右旋：
			newNode = n.LeftRightRotation()
		}
	} else if factor == -2 { // 右子树的高度变高了，导致左子树-右子树的高度从-1变成了-2。
		if value > n.Right.Value {
			// 表示在右子树上插上右儿子导致失衡，需要单左旋：
			newNode = n.LeftRotation()
		} else {
			//表示在右子树上插上左儿子导致失衡，先由后左旋：
			newNode = n.RightLeftRotation()
		}
	} else {
		newNode = n
	}
	return newNode
}

func (n *AvlTreeNode) Insert(value int) *AvlTreeNode {
	if n == nil {
		return &AvlTreeNode{Height: 1, Nums: 1, Value: value}
	}
	// 插入的值小于节点值，要从左子树继续插入
	if value < n.Value {
		n.Left = n.Left.Insert(value)
		n = n.HandleBF(value)
		// 插入的值大于节点值，要从右子树继续插入
	} else if value > n.Value {
		n.Right = n.Right.Insert(value)
		n = n.HandleBF(value)
		// 插入值相同,重复次数加1即可
	} else {
		n.Nums += 1
	}
	// 刷新新树根高度
	n.UpdateHeight()
	return n
}

// 平衡二叉树打印
// 为了直观感受平衡二叉树的构造结构,我们实现一个简单的打印功能
func (tree *AvlTree) Print() {
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
		n := q.DeQueue().(*AvlTreeNode)
		// 左子树非空,左子树乖乖入队
		if n.Left != nil {
			q.EnQueue(n.Left)
			fmt.Printf("[value/nums/height/leftNode] is [%v/%v/%v/%v]\n", n.Value, n.Nums, n.Height, n.Left.Value)
		}
		// 右子树非空,左子树乖乖入队
		if n.Right != nil {
			q.EnQueue(n.Right)
			fmt.Printf("[value/nums/height/rightNode] is [%v/%v/%v/%v]\n", n.Value, n.Nums, n.Height, n.Right.Value)
		}
		if n.Left == nil && n.Right == nil {
			fmt.Printf("[value/nums/height/isLeafNode] is [%v/%v/%v]\n", n.Value, n.Nums, n.Height)
		}
	}
}

// 中序遍历实现排序
func (tree *AvlTree) PreOrder() {
	if tree.Root == nil {
		return
	}
	tree.preOrder(tree.Root)
	return
}

// 中序遍历实现排序
func (tree *AvlTree) preOrder(node *AvlTreeNode) {
	if node == nil {
		return
	}
	tree.preOrder(node.Left)
	fmt.Printf("travel value is:%v\n", node.Value)
	tree.preOrder(node.Right)
	return
}

// 二叉查找树的搜索
func (tree *AvlTree) Search(value int) *AvlTreeNode {
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

// 查询最小值
func (tree *AvlTree) FindMin() *AvlTreeNode {
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
func (tree *AvlTree) FindMax() *AvlTreeNode {
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

func (n *AvlTreeNode) FindMin() *AvlTreeNode {
	currentNode := n
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

func (tree *AvlTree) Delete(value int) {
	// 空树
	if tree.Root == nil {
		return
	}
	tree.Root = tree.Root.Delete(value)
	return
}

/*删除元素
1. 如果被删除结点A有两个子结点，将该结点右子树内的最小结点取代A, 并删除最小节点。
2. 如果被删除结点A只有一个子结点(左孩子或者右孩子)，就直接将A的子结点连至A的父结点上，并将A删除(根据平衡二叉树的定义, 一个节点只有一个子节点的时候，该子节点比为该节点的孩子节点,否则就失衡)
3. 删除的节点是叶子节点，没有儿子，直接删除后看离它最近的父亲节点是否失衡，做调整操作。
*/
func (n *AvlTreeNode) Delete(value int) *AvlTreeNode {
	if n == nil {
		return nil
	}
	if value < n.Value {
		// 从左子树开始删除
		n.Left = n.Left.Delete(value)
	} else if value > n.Value {
		// 从右子树开始删除
		n.Right = n.Right.Delete(value)
	} else {
		// 找到待删除的节点
		if n.Left != nil && n.Right != nil {
			// 有两个节点
			// 步骤1: 找到右孩子的最小值替换该待删除节点
			rightMinNode := n.Right.FindMin()
			n.Value = rightMinNode.Value
			n.Nums = rightMinNode.Nums
			// 步骤2: 删除右孩子的最小值
			n.Right = n.Right.Delete(n.Value)
		} else if n.Left != nil {
			// 只有左孩子
			n = n.Left
		} else if n.Right != nil {
			// 只有右孩子
			n = n.Right
		} else {
			// 左右孩子均无,为叶子节点
			n = nil
			return n
		}
	}
	// 调整高度
	n.UpdateHeight()
	return n
}
