package heap

// A Node stores a value and a priority such that getting the minimal priority of a Heap is easy
type Node[PRIORITY Number, T any] struct {
	// all of the elements
	priority PRIORITY
	element  T
	// the list of all children
	children NodeList[PRIORITY, T]

	siblings NodeList[PRIORITY, T]
	sib_idx  int
}

func makeNode[PRIORITY Number, T any](el T, priority PRIORITY) *Node[PRIORITY, T] {
	var n *Node[PRIORITY, T] = new(Node[PRIORITY, T])

	n.priority = priority
	n.element = el

	return n
}

func (n *Node[PRIORITY, T]) Degree() int {
	return len(n.children)
}

func (n *Node[PRIORITY, T]) merge(m *Node[PRIORITY, T]) *Node[PRIORITY, T] {
	// if n is not the root of the new tree
	if n.priority < m.priority {
		// we flip the calls
		return m.merge(n)
	}
	// m is the subtree
	m.sib_idx = len(n.children)
	m.siblings = n.children
	n.children = append(n.children, m)
	return n
}

func (n *Node[PRIORITY, T]) GetValue() T {
	return n.element
}
