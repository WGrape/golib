package set

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	var (
		s1 = []int64{
			1, 5, 6, 6, 6, 7,
		}
		s2 = []int64{
			1, 6, 7, 7,
		}
	)
	if !reflect.DeepEqual(Intersect(s1, s2), []int64{1, 6, 7}) {
		t.Fail()
		return
	}

	s1 = []int64{1}
	s2 = []int64{1}
	if !reflect.DeepEqual(Intersect(s1, s2), []int64{1}) {
		t.Fail()
		return
	}

	s1 = []int64{1}
	s2 = []int64{2}
	if len(Intersect(s1, s2)) != 0 {
		t.Fail()
		return
	}

	s1 = []int64{2, 3}
	s2 = []int64{2}
	if !reflect.DeepEqual(Intersect(s1, s2), []int64{2}) {
		t.Fail()
		return
	}

	s1 = []int64{2, 3}
	s2 = []int64{2, 3}
	if !reflect.DeepEqual(Intersect(s1, s2), []int64{2, 3}) {
		t.Fail()
		return
	}

	s1 = []int64{2, 2, 2, 3}
	s2 = []int64{2, 3, 3}
	if !reflect.DeepEqual(Intersect(s1, s2), []int64{2, 3}) {
		t.Fail()
		return
	}

	s1 = []int64{2, 2, 2, 2}
	s2 = []int64{2, 3, 3}
	if !reflect.DeepEqual(Intersect(s1, s2), []int64{2}) {
		t.Fail()
		return
	}
}

func TestDifference(t *testing.T) {
	var (
		s1 = []int64{
			1, 5, 5, 6, 6, 6, 7,
		}
		s2 = []int64{
			1, 6, 7, 7,
		}
	)
	if !reflect.DeepEqual(Difference(s1, s2), []int64{5}) {
		t.Fail()
		return
	}

	s1 = []int64{2, 3}
	s2 = []int64{1, 2}
	if !reflect.DeepEqual(Difference(s1, s2), []int64{3}) {
		t.Fail()
		return
	}

	s1 = []int64{2}
	s2 = []int64{1, 2}
	if len(Difference(s1, s2)) != 0 {
		t.Fail()
		return
	}

	s1 = []int64{2, 2, 2, 2}
	s2 = []int64{1, 2}
	if len(Difference(s1, s2)) != 0 {
		t.Fail()
		return
	}

	s1 = []int64{2, 2, 2, 2}
	s2 = []int64{1, 2, 2, 2, 3}
	if len(Difference(s1, s2)) != 0 {
		t.Fail()
		return
	}

	s1 = []int64{2, 2, 2, 2, 3, 4}
	s2 = []int64{1, 2, 2, 2, 3}
	if !reflect.DeepEqual(Difference(s1, s2), []int64{4}) {
		t.Fail()
		return
	}

	s1 = []int64{2, 2, 2, 2, 3, 4}
	s2 = []int64{1}
	if !reflect.DeepEqual(Difference(s1, s2), []int64{2, 3, 4}) {
		t.Fail()
		return
	}

	s1 = []int64{1, 2, 3, 4, 5, 6}
	s2 = []int64{3, 5}
	if !reflect.DeepEqual(Difference(s1, s2), []int64{1, 2, 4, 6}) {
		t.Fail()
		return
	}

	s1 = []int64{1, 2}
	s2 = []int64{2, 2, 3, 5}
	if !reflect.DeepEqual(Difference(s1, s2), []int64{1}) {
		t.Fail()
		return
	}
}
