package binary_tree

import (
  "fmt"
)

type treeNode struct {
  data int
  left *treeNode
  right *treeNode
}

// This is a BST really.
type BinaryTree struct {
  root *treeNode
}

func NewBinaryTree() *BinaryTree {
  return &BinaryTree{}
}

func (b *BinaryTree) Insert(value int) error {
  node := &treeNode{
    data: value,
  }

  if b.root == nil {
    b.root = node
    return nil
  }

  current := b.root
  var parent *treeNode = nil

  for true {
    parent = current

    if (value < current.data) {
      current = current.left

      if current == nil {
        parent.left = node
        return nil
      }
    } else {
      current = current.right

      if current == nil {
        parent.right = node
        return nil
      }
    }
  }

  return fmt.Errorf("something went wrong inserting the node")
}

func (b *BinaryTree) Contains(value int) bool {
  current := b.root

  for current != nil {
    if current.data == value {
      return true
    } else if value < current.data {
      current = current.left
    } else {
      current = current.right
    }
  }

  return false
}

func (b *BinaryTree) Delete(value int) error {
  parent := b.root
  current := b.root
  var isLeftChild bool = false

  for current.data != value {
    parent = current

    if value < current.data {
      isLeftChild = true
      current = current.left
    } else {
      isLeftChild = false
      current = current.right
    }

    if current == nil {
      return fmt.Errorf("value doesn't exist in tree. can't delete.")
    }
  }

  // No children
  if current.left == nil && current.right == nil {
    if current == b.root {
      b.root = nil
    }

    if isLeftChild {
      parent.left = nil
    } else {
      parent.right = nil
    }

    return nil
  }

  // If node only has one child
  if current.right == nil {
    if current == b.root {
      b.root = current.left
    } else if isLeftChild {
      parent.left = current.left
    } else {
      parent.right = current.left
    }

    return nil
  }

  if current.left == nil {
    if current == b.root {
      b.root = current.right
    } else if isLeftChild {
      parent.left = current.right
    } else {
      parent.right = current.right
    }

    return nil
  }

  // If node has all children

  successor := getSuccessor(current)

  if current == b.root {
    b.root = successor
  } else if isLeftChild {
    parent.left = successor
  } else {
    parent.right = successor
  }

  successor.left = current.left
  return nil
}

type DisplayType int

const (
  Prefix DisplayType = iota
  Infix
  Postfix
)

func (b *BinaryTree) Display(displayType DisplayType) {
  if b.root != nil {
    display(b.root, displayType)
  }

  fmt.Println()
}

func display(node *treeNode, displayType DisplayType) {
  if node != nil {
    if displayType == Prefix {
      fmt.Printf(" %d", node.data)
    }
    display(node.left, displayType)
    if displayType == Infix {
      fmt.Printf(" %d", node.data)
    }
    display(node.right, displayType)
    if displayType == Postfix {
      fmt.Printf(" %d", node.data)
    }
  }
}

func getSuccessor(deleteNode *treeNode) *treeNode {
  var successor *treeNode = nil
  var successorParent *treeNode = nil
  current := deleteNode.right

  for current != nil {
    successorParent = successor
    successor = current
    current = current.left
  }

  if successor != deleteNode.right {
    successorParent.left = successor.right
    successor.right = deleteNode.right
  }

  return successor
}
