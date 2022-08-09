package heap

import (
	"fmt"
	"sync"
)

/*
A subtype of Heap where the content is comparable,
    ie one node can have equal content to another
*/
type CHeap[P Number, T comparable] Heap[P, T]

// asynchronously queries every node to get the list of nodes that contain T
func (ch *CHeap[P, T]) GetNodes(content T) ([]*Node[P, T], error) {
	if ch.numberOfRoots == 0 {
		return nil, fmt.Errorf("This heap is empty")
	}

	gatherer := make(chan *Node[P, T])
	doneChan := make(chan bool)
	var wg sync.WaitGroup

	go func() {
		wg.Wait()
		doneChan <- true
	}()

	reslt := make([]*Node[P, T], 0, 32)
	curNode := ch.leftMostRoot
	wg.Add(ch.numberOfRoots)
	for rootidx := 0; rootidx < ch.numberOfRoots; rootidx++ {
		go findDesc(curNode, content, gatherer, &wg)
		curNode = curNode.rightSib
	}

GATHER:
	for {
		select {
		case n := <-gatherer:
			reslt = append(reslt, n)
		case <-doneChan:
			break GATHER
		}
	}

	return reslt, nil
}

func findDesc[P Number, T comparable](root *Node[P, T], content T,
	c chan *Node[P, T], wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(root.degree)
	fmt.Println(root.GetValue(), content)
	if root.GetValue() == content {
		c <- root
	}
	curNode := root.leftMostChild
	for i := 0; i < root.degree; i++ {
		if curNode == nil {
			continue
		}
		go findDesc[P, T](curNode, content, c, wg)
		curNode = curNode.rightSib
	}
}

func (ch *CHeap[P, T]) clean() { (*Heap[P, T])(ch).clean() }

func (ch *CHeap[P, T]) IncreasePriority(n *Node[P, T], newPriority P) {
	(*Heap[P, T])(ch).IncreasePriority(n, newPriority)
}

func (ch *CHeap[P, T]) ExtractMax() { (*Heap[P, T])(ch).ExtractMax() }

func (ch *CHeap[P, T]) GetMax() { (*Heap[P, T])(ch).GetMax() }

func (ch *CHeap[P, T]) Insert(priority P, element T) {
	(*Heap[P, T])(ch).Insert(priority, element)
}

func (ch *CHeap[P, T]) InsertNode(n *Node[P, T]) {
	(*Heap[P, T])(ch).InsertNode(n)
}
