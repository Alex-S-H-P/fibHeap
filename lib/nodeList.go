package heap

import "fmt"

var maxDegree int = 3 // by default, we can store about 2^10 values

type NodeList[P Number, T any] []*node[P, T]

// A type of NodeList such that at index i lies a tree of degree i
type DegreeArray[P Number, T any] NodeList[P, T]

func (d *DegreeArray[P, T]) assign(tree *node[P, T]) {
	if (*d)[tree.Degree()] == nil {
		(*d)[tree.Degree()] = tree
		fmt.Println("\t\tCould Put", tree.GetValue(), tree.priority)
	} else {
		otherTree := (*d)[tree.Degree()]
		(*d)[tree.Degree()] = nil
		fmt.Println("\t\tMerging", tree.GetValue(), tree.priority, " with ", otherTree.GetValue())
		d.assign(otherTree.merge(tree))
	}
	fmt.Printf("\t\t")
	d.Show()
}

func (d *DegreeArray[P, T]) Show() {
	fmt.Printf("| ")
	for _, tree := range *d {
		if tree == nil {
			fmt.Printf("[] | ")
		} else {
			fmt.Printf("%v, %v | ", tree.priority, tree.element)
		}
	}
	fmt.Printf("\n")
}
