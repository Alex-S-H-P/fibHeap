package heap

type Heap[P Number, T any] struct {
	rootList NodeList[P, T]

	maxNode *node[P, T]
}

// inserts and updates maximum
func (h *Heap[P, T]) insert(n ...*node[P, T]) {
	var prevLen int = len(h.rootList)
	h.rootList = append(h.rootList, n...)

	// setting variables to their new correct values
	for i, newRoot := range n {
		newRoot.sib_idx = prevLen + i
		newRoot.siblings = h.rootList
		if h.maxNode == nil || newRoot.priority > h.maxNode.priority {
			h.maxNode = newRoot
		}
	}

}

/*
returns the stored node with max value.
*/
func (h *Heap[P, T]) GetMax() *node[P, T] {
	return h.maxNode
}

func (h *Heap[P, T]) Insert(priority P, element T) {
	n := makeNode(element, priority)
	h.insert(n)
}

// Deletes the max, and returns it
func (h *Heap[P, T]) ExtractMax() (P, T) {
	if h.maxNode == nil {
		var reslt T
		return 0, reslt
	}

	result := h.maxNode
	h.insert(h.maxNode.children...)
	h.rootList = append(h.rootList[:h.maxNode.sib_idx],
		h.rootList[h.maxNode.sib_idx+1:]...)

	// cleanup
	if len(h.rootList) > maxDegree {
		h.clean()
	}
	return result.priority, result.element
}

func (h *Heap[P, T]) clean() {
	var degreeArray DegreeArray[P, T] = make([]*node[P, T], maxDegree+1)
	for _, tree := range h.rootList {
		degreeArray.assign(tree)
	}

	var newHeap NodeList[P, T] = make([]*node[P, T], 0, len(degreeArray))
	var maxPrioritySet bool

	for _, tree := range degreeArray[1:] {
		if tree == nil {
			continue
		}

		if !maxPrioritySet || h.maxNode.priority < tree.priority {
			maxPrioritySet = true
			h.maxNode = tree
		}
		newHeap = append(newHeap, tree)
	}
	h.rootList = newHeap
}
