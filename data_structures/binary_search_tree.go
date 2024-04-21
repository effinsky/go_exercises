package datastructures

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// BST property: For any node N, the value of any node in N’s left
// subtree is less than N’s value, and the value of any node in N’s
// right subtree is greater than N’s value. Unique values to start with.
// Consider dealing with duplicates later..

// The power of binary search trees, and dynamic data structures in general,
// arises in cases where the data is changing. If we are doing many searches
// over a dynamic data set, this combination of efficiencies (of dealing with
// a sorted array) becomes critical.
type (
	BST[I item] struct{ Root *node[I] }

	item interface{ constraints.Ordered }

	node[I item] struct {
		v  I        // value
		p  *node[I] // parent
		lc *node[I] // left child
		rc *node[I] // right child
		// consider storing aux data, for instance counter for duplicate items if
		// such occur
	}
)

func NewBST[I item](v I) BST[I] {
	return BST[I]{
		Root: &node[I]{v: v},
	}
}

func (t *BST[I]) FindValue(tgt I) *node[I] {
	return findValueAux(t.Root, tgt)
}

// params: current, target
func findValueAux[I item](curr *node[I], tgt I) *node[I] {
	fmt.Printf("curr: %+v\n", curr)
	if curr == nil {
		return nil
	}
	if curr.v == tgt {
		return curr
	}
	if curr.v < tgt {
		return findValueAux(curr.lc, tgt)
	}
	if curr.v > tgt { // simplify?
		return findValueAux(curr.rc, tgt)
	}
	return nil
}

// params: value to be inserted
func (t *BST[I]) Insert(v I) {
	if t.Root == nil {
		t.Root = &node[I]{v: v}
	} else {
		insertAux(t.Root, v)
	}
}

// params: ptr to current node, value to be inserted
func insertAux[I item](curr *node[I], v I) {
	switch {
	case v == curr.v:
		fmt.Printf("node with value %v already in tree", v)
		return

	case v < curr.v:
		if curr.lc != nil {
			insertAux(curr.lc, v)
		} else {
			curr.lc = &node[I]{v: v}
			curr.lc.p = curr
		}

	default:
		if curr.rc != nil {
			insertAux(curr.rc, v)
		} else {
			curr.rc = &node[I]{v: v}
			curr.rc.p = curr
		}
	}
}

func (t *BST[I]) TraversePreOrder() {
	curr := t.Root
	traversePreorderAux(curr)
}

// params: ptr to current node
func traversePreorderAux[I item](curr *node[I]) {
	if curr == nil {
		return
	}

	fmt.Printf("%v\n", curr.v)

	traversePreorderAux(curr.lc)
	traversePreorderAux(curr.rc)
}

// params: value to be deleted, stored in the tree under some node to be located
func (t *BST[I]) Delete(v I) bool {
	if t.Root == nil {
		return false
	}

	curr := t.FindValue(v)
	if curr == nil {
		return false
	}

	// Case A: Deleting a leaf node.
	if curr.lc == nil && curr.rc == nil {
		t.deleteLeafNode(curr)
		return true // removed
	}

	// Case B: Deleting a node with one child
	// (don't know which one is there..)
	if curr.lc == nil || curr.rc == nil {
		t.deleteWithOneChild(curr)
		return true // removed
	}

	// else Case C: Deleting a node with two children.
	t.deleteWithTwoChildren(curr) // need v itself for recursion
	return true
}

func (t *BST[I]) deleteLeafNode(curr *node[I]) {
	if curr.p == nil { // you gotta have a parent, no?
		t.Root = nil
	} else if curr.p.lc == curr {
		curr.p.lc = nil
	} else {
		curr.p.rc = nil
	}
}

func (t *BST[I]) deleteWithOneChild(curr *node[I]) {
	child := curr.lc
	if curr.rc != nil {
		child = curr.rc
	}
	// Now we have the child, need to assign its parent to the removed node's
	// parent.
	child.p = curr.p
	// the child may become the root..
	if curr.p == nil {
		t.Root = child
		// if the node to be removed is its parent's left child..
	} else if curr.p.lc == curr {
		curr.p.lc = child
		// or if the node to be removed is its parent's right child..
	} else {
		curr.p.rc = child
	}
}

func (t *BST[I]) deleteWithTwoChildren(curr *node[I]) {
	succ := findSuccessor(curr)

	if succ.p != curr {
		succ.lc = succ.rc
		if succ.rc != nil {
			succ.rc.p = succ.p
		}
		succ.rc = curr.rc
		curr.rc.p = succ
	}

	// Set the parent of the successor to the parent of the node to be
	// deleted..
	succ.p = curr.p
	if curr.p == nil {
		t.Root = succ
	} else if curr.p.lc == curr { // if the node to be deleted is a lc..
		curr.p.lc = succ
	} else { // if the node to be deleted is a rc..
		curr.p.rc = succ
	}

	// Set the left child of the successor to the left child of the node
	// to be deleted..
	succ.lc = curr.lc // make child relation on succ as parent
	curr.lc.p = succ  // make parent relation on deleted node's lc
}

// Make the leftmost node in the right subtree of the deleted node the successor,
func findSuccessor[I item](curr *node[I]) *node[I] {
	s := curr.rc      // go right
	for s.lc != nil { // keep going left until you've found a leaf..
		s = s.lc
	}
	return s
}
