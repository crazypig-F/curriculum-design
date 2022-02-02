package tree

import (
	"HuffmanCoding/data"
)

type HuffmanTreeNode struct {
	Val    string
	Number int
	Left   *HuffmanTreeNode
	Right  *HuffmanTreeNode
}

func ChangeNode(queue []data.Node) []*HuffmanTreeNode {
	HuffmanTreeQueue := make([]*HuffmanTreeNode, len(queue))
	for index, node := range queue {
		var val string
		var num int
		for key, v := range node {
			val = key
			num = v
		}
		n := new(HuffmanTreeNode)
		n.Val = val
		n.Number =  num
		HuffmanTreeQueue[index] = n
	}
	return HuffmanTreeQueue
}

func InsertHuffmanTreeQueue(huffmanTreeQueue []*HuffmanTreeNode, node *HuffmanTreeNode) []*HuffmanTreeNode{
	for index, n := range huffmanTreeQueue {
		if node.Number <= n.Number {
			front := make([]*HuffmanTreeNode, index)
			behind := make([]*HuffmanTreeNode, len(huffmanTreeQueue)-index)
			copy(front, huffmanTreeQueue[:index])
			copy(behind, huffmanTreeQueue[index:])
			front = append(front, node)
			huffmanTreeQueue = append(front, behind...)
			break
		}
		if index == len(huffmanTreeQueue)-1{
			huffmanTreeQueue = append(huffmanTreeQueue, node)
		}
	}
	return huffmanTreeQueue
}

func CreateTree(huffmanTreeQueue []*HuffmanTreeNode) *HuffmanTreeNode{
	node := new(HuffmanTreeNode)
	if len(huffmanTreeQueue) > 1{
		node.Left = huffmanTreeQueue[0]
		node.Right = huffmanTreeQueue[1]
		node.Val = "node"
		node.Number = huffmanTreeQueue[0].Number + huffmanTreeQueue[1].Number
		huffmanTreeQueue = InsertHuffmanTreeQueue(huffmanTreeQueue[2:], node)
	}
	if len(huffmanTreeQueue) == 1{
		node.Left = huffmanTreeQueue[0]
		node.Val = "node"
		return node
	}
	if len(huffmanTreeQueue) == 0 {
		return node
	}
	return CreateTree(huffmanTreeQueue)
}
