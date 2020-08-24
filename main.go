package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	key int
	left *Node
	right *Node
	parent *Node
}

func (n *Node) Find(key int) *Node {
	if key == n.key {
		return n
	}

	if n.left != nil && key < n.key {
		return n.left.Find(key)
	}

	if n.right != nil && key > n.key {
		return n.right.Find(key)
	}
	return nil
}

func (t *Node) findParentNode(key int) *Node {
	node := t
	for {
		if node.key == key {
			break
		}
		if key < node.key {
			if node.left != nil {
				node = node.left
				continue
			}
			break
		}

		if node.right != nil {
			node = node.right
			continue
		}
		break
	}
	return node
}

func (n *Node) Insert(key int) {
	newNode := &Node{
		key: key,
	}
	lastNode := n.findParentNode(key)
	newNode.parent = lastNode
	if key < lastNode.key {
		lastNode.left = newNode
		return
	}
	lastNode.right = newNode
}

func (n *Node) Remove(key int) {
	if key < n.key {
		n.left.Remove(key)
		return
	}
	if key > n.key {
		n.right.Remove(key)
		return
	}

	if key == n.key {
		// Первый случай, у узла нет потомков
		if n.left == nil && n.right == nil {
			if key > n.parent.key {
				n.parent.right = nil
				return
			}
			n.parent.left = nil
			return
		}

		// Второй случай, у зла 1 потомок
		if n.left == nil {
			if key > n.parent.key {
				n.parent.right = n.right
				return
			}
			return
		}

		if n.right == nil {
			if key > n.parent.key {
				n.parent.right = n.left
				return
			}
			n.parent.left = n.left
			return
		}

		// 3 случай, у удаляемого узла 2 потомка
		maxLeftNode := n.right
		for {
			if maxLeftNode.left == nil {
				break
			}
			maxLeftNode = n.right.left
		}

		maxLeftNode.left = n.left
		n.left.parent = maxLeftNode
		if (key > n.parent.key) {
			n.parent.right = n.right
			return
		}
		n.parent.left = n.right
		return
	}
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}
	text := ""
	parentNode := n
	for parentNode != nil {
		parentNode = parentNode.parent
		text += "-"
	}
	text += " " + strconv.Itoa(n.key) + "\n"

	text += n.left.String()
	text += n.right.String()

	return text
}

func main() {
	node := &Node{
		key:44,
	}
	node.left = &Node{
		key: 22,
		parent: node,
	}

	node.Insert(11)
	node.Insert(99)
	node.Insert(33)
	node.Insert(9)
	node.Insert(31)
	node.Insert(41)
	node.Insert(98)
	node.Insert(120)
	node.Insert(124)
	node.Insert(101)


	fmt.Println(node)
	fmt.Println("---")
	node.Remove(31)
	node.Remove(22)
	//fmt.Println(node.find(55))
	//fmt.Println(node.findLastNode(33))
	fmt.Println(node)
}
