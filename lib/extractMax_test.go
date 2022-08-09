package heap

import (
	"testing"
)

const (
	VAL1 string = "hi"
	VAL2 string = "hello"
	VAL3 string = "GREETING"
	VAL4 string = "This is a public service announcement"
)

func TestHeap0(t *testing.T) {
	h := new(Heap[int, string])
	h.Insert(1, VAL1)
	if h.GetMax() == nil {
		t.Errorf("Failed to write valid max value")
	}
	if h.GetMax().GetValue() != VAL1 {
		t.Errorf("Incorrectly set up Node [Value]. Got %v instead of %v",
			h.GetMax().GetValue(), VAL1)
	}
	if h.GetMax().Degree() != 0 {
		t.Errorf("Incorrectly set up Node [Children]. Got %v instead of %v",
			h.GetMax().Degree(), 0)
	}
	p, v := h.ExtractMax()
	if p != 1 || v != VAL1 {
		t.Errorf("Could not extract properly [Value]. Got (%v, %v) instead of (%v, %v)",
			p, v, 1, VAL1)
	}
	p, v = h.ExtractMax()
	if p != 0 || v != "" {
		t.Errorf("Could not extract properly [ValuesWhenEmpty]. Got (%v, %v) instead of (%v, %v)",
			p, v, 0, "")
	}
}

func TestHeap1(t *testing.T) {
	h := new(Heap[int, string])
	h.Insert(1, VAL1)
	h.Insert(2, VAL2)
	h.Insert(4, VAL3)
	h.Insert(3, VAL4)
	/* h:

	roots 	: 	[(1 VAL1)	(2 VAL2)	(4 VAL3)	(3 VAL4)]
	maxNode	:	-> (4 VAL3)
	*/
	p, v := h.ExtractMax()
	if p != 4 || v != VAL3 {
		t.Errorf("Could not extract properly [Value]. Got (%v, %v) instead of (%v, %v)",
			p, v, 4, VAL3)
	}
	/* h:

	roots 	: 	[( 3 VAL4 {( 2 VAL2 {( 1 VAL1 )} )} )]
	maxNode	:	-> (3 VAL4)
	/
	if h.GetMax().Degree() != 1 {
		t.Errorf("Incorrectly set up node [Children]. Got %v instead of %v",
			h.GetMax().Degree(), 1)
	}
	*/
	p, v = h.ExtractMax()
	if p != 3 || v != VAL4 {
		t.Errorf("Could not extract properly [Value]. Got (%v, %v) instead of (%v, %v)",
			p, v, 3, VAL4)
	}

	/* h:

	roots 	: 	[( 2 VAL2 {( 1 VAL1 )} )]
	maxNode	:	-> (2 VAL2)
	*/
	p, v = h.ExtractMax()
	if p != 2 || v != VAL2 {
		t.Errorf("Could not extract properly [Value]. Got (%v, %v) instead of (%v, %v)",
			p, v, 2, VAL2)
	}
}

func TestHeap2(t *testing.T) {
	h := new(Heap[int, string])
	h.Insert(1, VAL1)
	h.Insert(2, VAL2)
	h.Insert(4, VAL3)
	h.Insert(3, VAL4)
	ch := (*CHeap[int, string])(h)
	arr, err := ch.GetNodes(VAL2)
	if err != nil {
		t.Error(err)
	}
	if len(arr) != 1 {
		t.Errorf("Did not find a good amount of solutions. Expected %v, found %v",
			1, len(arr))
	}
	ch.clean()
	if err != nil {
		t.Error(err)
	}
	if len(arr) != 1 {
		t.Errorf("Did not find a good amount of solutions. Expected %v, found %v",
			1, len(arr))
	}

}
