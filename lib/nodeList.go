package heap

var maxDegree int = 10 // by default, we can store about 2^10 values

type NodeList[P Number, T any] struct {
	all []*Node[P, T]
}

func (N *NodeList[P, T]) append(newN ...*Node[P, T]) {
	N.all = append(N.all, newN...)
}

func (N *NodeList[P, T]) Get(i int) *Node[P, T] {
	return N.all[i]
}

func (N *NodeList[P, T]) Len() int {
	return len(N.all)
}

func (N *NodeList[P, T]) Eliminate(idx int) {
	N.all = append(N.all[:idx], N.all[idx+1:]...)
}

func newNodeList[P Number, T any](cap int) *NodeList[P, T] {
	n := new(NodeList[P, T])
	n.all = make([]*Node[P, T], 0, cap)
	return n
}

// A type of NodeList such that at index i lies a tree of degree i
type DegreeArray[P Number, T any] []*Node[P, T]

func (d *DegreeArray[P, T]) assign(tree *Node[P, T]) {
	if (*d)[tree.Degree()] == nil {
		(*d)[tree.Degree()] = tree
	} else {
		otherTree := (*d)[tree.Degree()]
		(*d)[tree.Degree()] = nil
		d.assign(otherTree.merge(tree))
	}
}
