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

	expected2 := NewSet(3, 6)
	actual2 := second.Difference(first)
	Assert(t, expected2.Equals(actual2))
}

func TestPop(t *testing.T) {
	foo := NewSet(42)
	popped := foo.Pop()
	Assert(t, 42 == popped)
	Assert(t, 0 == foo.Len())

	bar := NewSet(2, 3, 4)
	Assert(t, 3 == bar.Len())
	bar.Pop()
	Assert(t, 2 == bar.Len())
	bar.Pop()
	Assert(t, 1 == bar.Len())
	bar.Pop()
	Assert(t, 0 == bar.Len())
}

func TestClear(t *testing.T) {
	foo := NewSet(2, 4, 6, 8, "who do we appreciate")
	Assert(t, foo.Len() > 0)
	foo.Clear()
	Assert(t, foo.Len() == 0)
}

func TestSymmetricDifference(t *testing.T) {
	foo := NewSet(2, 3, 6, 7)
	bar := NewSet(3, 6, 9, 12)
	expected := NewSet(2, 7, 9, 12)
	actual1 := foo.SymmetricDifference(bar)
	actual2 := bar.SymmetricDifference(foo)
	Assert(t, expected.Equals(actual1))
	Assert(t, expected.Equals(actual2))
}

func TestIsDisjoint(t *testing.T) {
	foo := NewSet(2, 3, 6, 7)
	bar := NewSet(3, 6, 9, 12)
	baz := NewSet(101, 202, 303)

	Assert(t, foo.IsDisjoint(baz))
	Assert(t, ! foo.IsDisjoint(bar))
}