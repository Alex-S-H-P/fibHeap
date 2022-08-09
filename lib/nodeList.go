package heap

var maxDegree int = 10 // by default, we can store about 2^10 values

type NodeList[P Number, T any] []*Node[P, T]

// A type of NodeList such that at index i lies a tree of degree i
type DegreeArray[P Number, T any] NodeList[P, T]

func (d *DegreeArray[P, T]) assign(tree *Node[P, T]) {
	if (*d)[tree.Degree()] == nil {
		(*d)[tree.Degree()] = tree
	} else {
		otherTree := (*d)[tree.Degree()]
		(*d)[tree.Degree()] = nil
		d.assign(otherTree.merge(tree))
	}
}
