package tree

import "fmt"

type Node struct {
	Data   interface{} //节点数据
	LChild *Node       //左子树
	RChild *Node       //右子树
}

func NewNode(data interface{}) *Node {
	return &Node{Data: data}
}

func (this *Node) String() string {
	return fmt.Sprintf("value:%+v, left child:%+v,right child:%+v", this.Data, this.LChild, this.RChild)
}
