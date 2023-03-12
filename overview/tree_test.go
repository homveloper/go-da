package overview_test

import (
	"testing"
)

type TreeNode[T any] struct {
	key   int
	value T

	left  *TreeNode[T]
	right *TreeNode[T]
}

func NewTreeNode[T any](key int, value T) *TreeNode[T] {
	return &TreeNode[T]{key: key, value: value}
}

type BTree[T any] struct {
	root *TreeNode[T]
}

func NewBTree[T any]() *BTree[T] {
	return &BTree[T]{}
}

func insertTreeNode[T any](rootNode *TreeNode[T], newTreeNode *TreeNode[T]) {
	if newTreeNode.key < rootNode.key {
		if rootNode.left == nil {
			rootNode.left = newTreeNode
		} else {
			insertTreeNode(rootNode.left, newTreeNode)
		}
	} else {
		if rootNode.right == nil {
			rootNode.right = newTreeNode
		} else {
			insertTreeNode(rootNode.right, newTreeNode)
		}
	}
}

func (t *BTree[T]) Insert(key int, value T) *TreeNode[T] {
	if t.root == nil {
		t.root = NewTreeNode(key, value)
		return t.root
	}

	newTreeNode := NewTreeNode(key, value)
	insertTreeNode(t.root, newTreeNode)

	return newTreeNode
}

func (t *BTree[T]) Begin() *TreeNode[T] {
	if t.root == nil {
		return nil
	}

	node := t.root
	for node.left != nil {
		node = node.left
	}

	return node
}

func (t *BTree[T]) End() *TreeNode[T] {
	if t.root == nil {
		return nil
	}

	node := t.root
	for node.right != nil {
		node = node.right
	}

	return node
}

func searchNode[T any](node *TreeNode[T], key int) *TreeNode[T] {
	if node == nil {
		return nil
	}

	if key == node.key {
		return node
	} else if key < node.key {
		return searchNode(node.left, key)
	} else {
		return searchNode(node.right, key)
	}
}

func (t *BTree[T]) Search(key int) *TreeNode[T] {
	return searchNode(t.root, key)
}

func deleteNode[T any](node *TreeNode[T], key int) *TreeNode[T] {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = deleteNode(node.left, key)
	} else if key > node.key {
		node.right = deleteNode(node.right, key)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		minNode := node.right
		for minNode.left != nil {
			minNode = minNode.left
		}

		node.key = minNode.key
		node.value = minNode.value
		node.right = deleteNode(node.right, minNode.key)
	}

	return node
}

func (t *BTree[T]) Delete(key int) bool {

	if t.root == nil {
		return false
	}

	return deleteNode(t.root, key) != nil
}

func (t *BTree[T]) InOrderIterate() <-chan *TreeNode[T] {
	ch := make(chan *TreeNode[T])

	go func() {
		defer close(ch)

		var iterate func(*TreeNode[T])
		iterate = func(node *TreeNode[T]) {
			if node == nil {
				return
			}

			iterate(node.left)
			ch <- node
			iterate(node.right)
		}

		iterate(t.root)
	}()

	return ch
}

func (t *BTree[T]) PreOrderIterate() <-chan *TreeNode[T] {
	ch := make(chan *TreeNode[T])

	go func() {
		defer close(ch)

		var iterate func(*TreeNode[T])
		iterate = func(node *TreeNode[T]) {
			if node == nil {
				return
			}

			ch <- node
			iterate(node.left)
			iterate(node.right)
		}

		iterate(t.root)
	}()

	return ch
}

func (t *BTree[T]) PostOrderIterate() <-chan *TreeNode[T] {
	ch := make(chan *TreeNode[T])

	go func() {
		defer close(ch)

		var iterate func(*TreeNode[T])
		iterate = func(node *TreeNode[T]) {
			if node == nil {
				return
			}

			iterate(node.left)
			iterate(node.right)
			ch <- node
		}

		iterate(t.root)
	}()

	return ch
}

func TestBTree(t *testing.T) {
	tree := NewBTree[int]()

	tree.Insert(5, 5)
	tree.Insert(3, 3)
	tree.Insert(7, 7)
	tree.Insert(1, 1)
	tree.Insert(4, 4)
	tree.Insert(6, 6)
	tree.Insert(8, 8)

	t.Log("InOrderIterate:")
	for node := range tree.InOrderIterate() {
		t.Log(node.key, node.value)
	}

	t.Log("PreOrderIterate:")
	for node := range tree.PreOrderIterate() {
		t.Log(node.key, node.value)
	}

	t.Log("PostOrderIterate:")
	for node := range tree.PostOrderIterate() {
		t.Log(node.key, node.value)
	}

	if tree.Begin().key != 1 {
		t.Error("Begin key error")
	}

	if tree.End().key != 8 {
		t.Error("End key error")
	}

	if tree.Search(5).value != 5 {
		t.Error("Search key error")
	}

	if tree.Delete(5) == false {
		t.Error("Delete key error")
	}
}
