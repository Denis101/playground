package binary_tree

import (
  "testing"
)

func TestBinaryTree(t *testing.T) {
  bst := NewBinaryTree()
  bst.Insert(3)
  bst.Insert(8)
  bst.Insert(1)
  bst.Insert(4)
  bst.Insert(6)
  bst.Insert(2)
  bst.Insert(10)
  bst.Insert(9)
  bst.Insert(20)
  bst.Insert(25)
  bst.Insert(15)
  bst.Insert(16)

  bst.Display(Infix)

  if !bst.Contains(4) {
    t.Errorf("Expected node of value 4 to exist in tree")
  }

  bst.Delete(2)
  if bst.Contains(2) {
    t.Errorf("Expected node of value 2 to have been removed from tree")
  }

  bst.Display(Infix)

  bst.Delete(4)
  if bst.Contains(4) {
    t.Errorf("Expected node of value 4 to have been removed from tree")
  }

  bst.Display(Infix)

  bst.Delete(10)
  if bst.Contains(10) {
    t.Errorf("Expected node of value 10 to have been removed from the tree")
  }

  bst.Display(Prefix)
  bst.Display(Infix)
  bst.Display(Postfix)

  // BinarySearchTree b = new BinarySearchTree();
	// 	b.insert(3);b.insert(8);
	// 	b.insert(1);b.insert(4);b.insert(6);b.insert(2);b.insert(10);b.insert(9);
	// 	b.insert(20);b.insert(25);b.insert(15);b.insert(16);
	// 	System.out.println("Original Tree : ");
	// 	b.display(b.root);
	// 	System.out.println("");
	// 	System.out.println("Check whether Node with value 4 exists : " + b.find(4));
	// 	System.out.println("Delete Node with no children (2) : " + b.delete(2));
	// 	b.display(root);
	// 	System.out.println("\n Delete Node with one child (4) : " + b.delete(4));
	// 	b.display(root);
	// 	System.out.println("\n Delete Node with Two children (10) : " + b.delete(10));
	// 	b.display(root);
}
