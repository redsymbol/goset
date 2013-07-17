package set

import (
	"testing"
)

func Assert(t *testing.T, cond bool) {
	if ! cond {
		t.Fail()
	}
}

func TestSet(t *testing.T) {
	foo := NewSet("alpha", "delta")
	Assert(t, foo.Contains("alpha"))
	Assert(t, ! foo.Contains("epsilon"))
	foo.Add("epsilon")
	Assert(t, foo.Contains("epsilon"))
	Assert(t, 3 == foo.Len())
	
	bar := NewSet("gamma", "delta", "eta")
	Assert(t, bar.Contains("eta"))
	bar.Remove("eta")
	bar.Discard("eta")
	bar.Discard("eta")
	Assert(t, ! bar.Contains("eta"))
	
	fooAndBar := foo.Intersect(bar)
	Assert(t, 1 == fooAndBar.Len())
	fooOrBar := foo.Union(bar)
	Assert(t, 4 == fooOrBar.Len())
}