package tree

import (
	"fmt"

	"github.com/yiye93/algo/queue"
	"github.com/yiye93/algo/stack"
)

//前序遍历
// 根节点 左子树 右子树
func PreOrder(tree *Node) {
	if tree == nil {
		return
	}
	//打印根节点值
	fmt.Printf("data:%v\n", tree.Data)
	//遍历左子树
	PreOrder(tree.LChild)
	//遍历右子树
	PreOrder(tree.RChild)
}

// 前序遍历
// 栈实现
func PreOrderV1(tree *Node) {
	if tree == nil {
		return
	}
	var elems = make([]interface{}, 0)
	var s = stack.NewStackOnLinklist()
	// 步骤1: 直到当前结点为空 & 栈空时，循环结束
	for tree != nil || !s.IsEmpty() {
		// 步骤2: 判断当前结点是否为空
		if tree != nil {
			// 步骤3: 保存当前节点值,并将该节点入栈
			elems = append(elems, tree.Data)
			s.Push(tree)
			// 步骤4: 置当前结点的左孩子为当前节点
			tree = tree.LChild
			// 返回步骤1
		} else {
			// 步骤5：出栈栈顶结点
			tree = s.Pop().(*Node)
			// 步骤6：置当前结点的右孩子为当前节点
			tree = tree.RChild
			// 返回步骤1
		}
	}
	for _, elem := range elems {
		fmt.Printf("data:%v\n", elem)
	}
	return
}

//中序遍历
//左子树 根节点 右子树
func MidOrder(tree *Node) {
	if tree == nil {
		return
	}
	//遍历左子树
	MidOrder(tree.LChild)
	//打印根节点值
	fmt.Printf("data:%v\n", tree.Data)
	//遍历右子树
	MidOrder(tree.RChild)
}

// 中序遍历
// 栈实现
func MidOrderV1(tree *Node) {
	if tree == nil {
		return
	}
	var elems = make([]interface{}, 0)
	var s = stack.NewStackOnLinklist()
	// 步骤1: 直到当前结点为空 & 栈空时，循环结束
	for tree != nil || !s.IsEmpty() {
		// 步骤2: 判断当前结点是否为空
		if tree != nil {
			// 步骤3: 将该节点入栈
			s.Push(tree)
			// 步骤4: 置当前结点的左孩子为当前节点
			tree = tree.LChild
			// 返回步骤1
		} else {
			// 步骤5：出栈栈顶结点
			tree = s.Pop().(*Node)
			// 步骤6：保存节点
			elems = append(elems, tree.Data)
			// 步骤7：置当前结点的右孩子为当前节点
			tree = tree.RChild
			// 返回步骤1
		}
	}
	for _, elem := range elems {
		fmt.Printf("data:%v\n", elem)
	}
	return
}

//后序遍历
//左子树 右子树 根节点
func PostOrder(tree *Node) {
	if tree == nil {
		return
	}
	//遍历左子树
	PostOrder(tree.LChild)
	//遍历右子树
	PostOrder(tree.RChild)
	//打印根节点值
	fmt.Printf("data:%v\n", tree.Data)
}

//后序遍历
//栈实现
func PostOrderV1(tree *Node) {
	if tree == nil {
		return
	}
	// 构建普通栈
	var s = stack.NewStackOnLinklist()
	// 构建中间输出栈
	var out = stack.NewStackOnLinklist()
	// 步骤1: 直到当前结点为空 & 栈空时，循环结束
	for tree != nil || !s.IsEmpty() {
		// 步骤2: 判断当前结点是否为空
		if tree != nil {
			// 步骤3: 将该节点入普通栈和中间输出栈
			s.Push(tree)
			out.Push(tree)
			// 步骤4: 置当前结点的右孩子为当前节点
			tree = tree.RChild
			// 返回步骤1
		} else {
			// 步骤5：出栈栈顶结点
			tree = s.Pop().(*Node)
			// 步骤6：置当前结点的左孩子为当前节点
			tree = tree.LChild
			// 返回步骤1
		}
	}
	for !out.IsEmpty() {
		fmt.Printf("data:%v\n", out.Pop().(*Node).Data)
	}
}

// 层次遍历
// 从上到下、从左到右
// 队列实现
func LevelTravel(tree *Node) {
	if tree == nil {
		return
	}
	elems := make([]interface{}, 0)
	// 此处搞了个容量限制...
	q := queue.NewQueueOnLinklist(10000)
	// 头元素入队
	q.EnQueue(tree)

	// 队列非空
	for !q.IsEmpty() {
		n := q.DeQueue().(*Node)
		elems = append(elems, n.Data)
		// 左子树非空,左子树乖乖入队
		if n.LChild != nil {
			q.EnQueue(n.LChild)
		}
		// 右子树非空,左子树乖乖入队
		if n.RChild != nil {
			q.EnQueue(n.RChild)
		}
	}
	for _, elem := range elems {
		fmt.Printf("data:%v\n", elem)
	}
}
