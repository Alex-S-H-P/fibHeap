package heap

type Heap[P Number, T any] struct {
	leftMostRoot  *Node[P, T]
	rightMostRoot *Node[P, T]

	numberOfRoots int

	maxNode *Node[P, T]
}

// InsertNodes and updates maximum
func (h *Heap[P, T]) InsertNode(n ...*Node[P, T]) {
	// OPTIMIZE: you can string them together beforehand and only update h once
	for _, newRoot := range n {
		if newRoot == nil {
			continue
		}
		// we update the number of roots
		h.numberOfRoots++

		if h.rightMostRoot != nil {
			h.rightMostRoot.rightSib = newRoot
			newRoot.leftSib = h.rightMostRoot
			h.rightMostRoot = newRoot
		} else {
			h.leftMostRoot = newRoot
			h.rightMostRoot = newRoot
			h.numberOfRoots = 1 // OPTIMIZE: not needed
		}
		// roots don't have parents
		newRoot.parent = nil

		if h.maxNode == nil || newRoot.priority >= h.maxNode.priority {
			h.maxNode = newRoot
		}
	}
}

/*
returns the stored Node with max value.
*/
func (h *Heap[P, T]) GetMax() *Node[P, T] {
	return h.maxNode
}

func (h *Heap[P, T]) Insert(priority P, element T) {
	n := makeNode(element, priority)
	h.InsertNode(n)
}

// Deletes the max, and returns it
func (h *Heap[P, T]) ExtractMax() (P, T) {
	if h.maxNode == nil {
		var reslt T
		return 0, reslt
	}

	result := h.maxNode
	h.InsertNode(result.extractAllChildren()...)
	if result.leftSib != nil {
		result.leftSib.rightSib = result.rightSib
	}
	if result.rightSib != nil {
		result.rightSib.leftSib = result.leftSib
	}
	if h.maxNode == h.leftMostRoot {
		h.leftMostRoot = result.rightSib
	}
	if h.rightMostRoot == h.maxNode {
		h.rightMostRoot = result.leftSib
	}
	result.leftSib, result.rightSib = nil, nil
	h.maxNode = nil
	h.numberOfRoots--

	// cleanup
	h.clean()
	if result == h.maxNode { // no operation did set a new maxNode
		h.maxNode = nil // we are empty
	}

	return result.priority, result.element
}

func (h *Heap[P, T]) clean() {
	var degreeArray DegreeArray[P, T] = make([]*Node[P, T], maxDegree+1)
	for i := 0; i < h.numberOfRoots; i++ {
		tree := h.leftMostRoot
		h.leftMostRoot = tree.rightSib
		degreeArray.assign(tree)
	}
	h.leftMostRoot = nil
	var maxPrioritySet bool = false
	h.numberOfRoots = 0

	for _, tree := range degreeArray {
		if tree == nil {
			continue
		}
		tree.parent = nil
		if h.leftMostRoot == nil {
			h.leftMostRoot = tree
			h.rightMostRoot = tree
			tree.leftSib = nil
			tree.rightSib = nil
		} else {
			h.rightMostRoot.rightSib = tree
			tree.leftSib = h.rightMostRoot
			h.rightMostRoot = tree
			tree.rightSib = nil
		}

		h.numberOfRoots++

		if maxPrioritySet && h.maxNode.priority > tree.priority {
		} else {
			h.maxNode = tree
			maxPrioritySet = true
		}
	}
}

/*
Increases the priority of a node
If the node is would get a larger priority than its parent, then adds it as a new root
*/
func (h *Heap[P, T]) IncreasePriority(n *Node[P, T], newPriority P) {
	if newPriority < n.priority {
		return
	}
	if n.parent == nil || n.parent.priority > newPriority {
		n.priority = newPriority
	} else {
		n.cutoutForIncreasePriority(h)
	}
}

/*
Merges h2 into h1 before emptying h2
*/
func (h *Heap[P, T]) Merge(h2 *Heap[P, T]) {
	h.rightMostRoot.rightSib = h2.leftMostRoot
	h2.leftMostRoot.leftSib = h.rightMostRoot.rightSib
	h.rightMostRoot = h2.rightMostRoot
	h.numberOfRoots += h2.numberOfRoots
	if h2.maxNode.priority >= h.maxNode.priority {
		h.maxNode = h2.maxNode
	}

	h2.leftMostRoot, h2.rightMostRoot, h2.maxNode = nil, nil, nil
	h2.numberOfRoots = 0

	if h.numberOfRoots > maxDegree {
		h.clean()
	}
}
