package set

import (
	"fmt"
	"testing"
)

func Assert(t *testing.T, cond bool) {
	if !cond {
		t.Fail()
	}
}

func AssertSliceEqual(t *testing.T, a *[]string, b *[]string) {
	if len(*a) != len(*b) {
		t.Fail()
	}
	var ii int
	for ii = 0; ii < len(*a); ii++ {
		aval := (*a)[ii]
		bval := (*b)[ii]
		if aval != bval {
			t.Log(fmt.Sprintf("'%v' != '%v' at index %d", aval, bval, ii))
		}
	}
}

func TestSetString(t *testing.T) {
	bag := NewSet("delta", "alpha")
	Assert(t, bag.Contains("alpha"))
	Assert(t, !bag.Contains("epsilon"))
	bag.Add("epsilon")
	Assert(t, bag.Contains("epsilon"))
	Assert(t, 3 == bag.Len())

	bar := NewSet("gamma", "delta", "eta")
	Assert(t, bar.Contains("eta"))
	bar.Remove("eta")
	bar.Discard("eta")
	bar.Discard("eta")
	Assert(t, !bar.Contains("eta"))

	bagAndBar := bag.Intersect(bar)
	Assert(t, 1 == bagAndBar.Len())
	bagOrBar := bag.Union(bar)
	Assert(t, 4 == bagOrBar.Len())

	expected := []string{"alpha", "delta", "epsilon"}
	actual := bag.Sorted()
	AssertSliceEqual(t, &expected, &actual)
}

func TestSetMixed(t *testing.T) {
	bag := NewSet("alpha", 42)
	bag.Add(32)
	bag.Add("beta")
	bag.Add(2.7)

	Assert(t, bag.Contains("alpha"))
	Assert(t, bag.Contains("beta"))
	Assert(t, bag.Contains(32))
	Assert(t, bag.Contains(42))
	Assert(t, bag.Contains(2.7))

	// check sorting
	expected := []string{
		"2.7",
		"32",
		"42",
		"alpha",
		"beta",
	}
	actual := bag.Sorted()
	AssertSliceEqual(t, &expected, &actual)
}

func TestSubset(t *testing.T) {
	alpha := NewSet(7, 3, 2, 1)
	beta := NewSet(7, 2)
	gamma := NewSet(7, 3, 2, 9)
	empty := NewSet()

	Assert(t, beta.IsSubsetOf(alpha))
	Assert(t, alpha.IsSupersetOf(beta))
	Assert(t, ! gamma.IsSubsetOf(alpha))

	Assert(t, empty.IsSubsetOf(empty))
	Assert(t, empty.IsSupersetOf(empty))
	Assert(t, alpha.IsSupersetOf(empty))
	Assert(t, ! empty.IsSupersetOf(alpha))
}

func TestEquals(t *testing.T) {
	alpha := NewSet(2, 3, 4, 5)
	beta := NewSet(2, 4, 3, 5)
	gamma := NewSet(2, 3, 4)
	delta := NewSet(2, 3, 4, 5, 6)

	Assert(t, alpha.Equals(beta))
	Assert(t, beta.Equals(alpha))
	Assert(t, ! alpha.Equals(gamma))
	Assert(t, ! alpha.Equals(delta))
}

func TestCopy(t *testing.T) {
	src := NewSet(4, 7, "foo", "bar")
	clone := src.Copy()

	Assert(t, src.Equals(clone))
}

func TestDifference(t *testing.T) {
	first := NewSet(2, 4, 8)
	second := NewSet(3, 6, 8)

	expected1 := NewSet(2, 4)
	actual1 := first.Difference(second)
	Assert(t, expected1.Equals(actual1))

	expected2 := NewSet(2, 4)
	actual2 := second.Difference(first)
	Assert(t, expected2.Equals(actual2))
}