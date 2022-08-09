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
