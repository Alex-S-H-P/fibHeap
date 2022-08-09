package heap

// A Node stores a value and a priority such that getting the minimal priority of a Heap is easy
type Node[P Number, T any] struct {
	// all of the elements
	priority P
	element  T
	// the list of all children
	leftMostChild  *Node[P, T]
	rightMostChild *Node[P, T]

	degree int

	parent *Node[P, T]

	leftSib  *Node[P, T]
	rightSib *Node[P, T]

	marked bool
}

func makeNode[P Number, T any](el T, priority P) *Node[P, T] {
	var n *Node[P, T] = new(Node[P, T])

	n.priority = priority
	n.element = el

	return n
}

func (n *Node[P, T]) Degree() int {
	return n.degree
}

func (n *Node[P, T]) merge(m *Node[P, T]) *Node[P, T] {
	// if n is not the root of the new tree
	if n.priority < m.priority {
		// we flip the calls
		return m.merge(n)
	}
	// m is the subtree
	if n.degree != 0 {
		n.rightMostChild.rightSib = m
		m.leftSib = n.rightMostChild
		n.rightMostChild = m
		m.parent = n
	} else {
		n.leftMostChild = m
		n.rightMostChild = m
		m.parent = n
	}

	n.degree++
	return n
}

func (n *Node[P, T]) GetValue() T {
	return n.element
}

// Empties the list of all children into a slice. The children are still linked
func (n *Node[P, T]) extractAllChildren() []*Node[P, T] {
	all := make([]*Node[P, T], n.degree)
	for i := 0; i < n.degree; i++ {
		all[i] = n.leftMostChild
		all[i].parent = nil
		n.leftMostChild = all[i].rightSib
	}
	n.degree = 0
	n.rightMostChild = nil
	return all
}

func (n *Node[P, T]) cutoutForIncreasePriority(h *Heap[P, T]) {
	n.parent.degree--
	n.leftSib.rightSib = n.rightSib
	n.rightSib.leftSib = n.leftSib
	h.InsertNode(n)
	if n.parent.marked {
		n.parent.cutoutForIncreasePriority(h)
	}
}
