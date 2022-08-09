package heap

// A Node stores a value and a priority such that getting the minimal priority of a Heap is easy
type Node[PRIORITY Number, T any] struct {
	// all of the elements
	priority PRIORITY
	element  T
	// the list of all children
	leftMostChild  *Node[PRIORITY, T]
	rightMostChild *Node[PRIORITY, T]

	degree int

	parent *Node[PRIORITY, T]

	leftSib  *Node[PRIORITY, T]
	rightSib *Node[PRIORITY, T]
}

func makeNode[PRIORITY Number, T any](el T, priority PRIORITY) *Node[PRIORITY, T] {
	var n *Node[PRIORITY, T] = new(Node[PRIORITY, T])

	n.priority = priority
	n.element = el

	return n
}

func (n *Node[PRIORITY, T]) Degree() int {
	return n.degree
}

func (n *Node[PRIORITY, T]) merge(m *Node[PRIORITY, T]) *Node[PRIORITY, T] {
	// if n is not the root of the new tree
	if n.priority < m.priority {
		// we flip the calls
		return m.merge(n)
	}
	// m is the subtree
	n.degree++
	n.rightMostChild.rightSib = m
	m.leftSib = n.rightMostChild
	n.rightMostChild = m
	m.parent = n
	return n
}

func (n *Node[PRIORITY, T]) GetValue() T {
	return n.element
}

// Empties the list of all children into a slice. The children are still linked
func (n *Node[PRIORITY, T]) extractAllChildren() []*Node[PRIORITY, T] {
	all := make([]*Node[PRIORITY, T], n.degree)
	for i := 0; i < n.degree; i++ {
		all[i] = n.leftMostChild
		all[i].parent = nil
		n.leftMostChild = all[i].rightSib
	}
	n.degree = 0
	n.rightMostChild = nil
	return all
}
