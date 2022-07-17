package permutation

import (
	"github.com/WGrape/golib/array"
	"strings"
	"testing"
)

// TestGetCombinations tests the function of GetCombinations.
func TestGetCombinations(t *testing.T) {
	combinations := GetCombinations([]string{"123", "456", "789"})
	for index, combination := range combinations {
		if index == 0 && !array.IsStringsEqual(combination, []string{"123", "456", "789"}) {
			t.Fail()
			return
		}
		if index == 1 && !array.IsStringsEqual(combination, []string{"123", "456"}) {
			t.Fail()
			return
		}
		if index == 2 && !array.IsStringsEqual(combination, []string{"123", "789"}) {
			t.Fail()
			return
		}
		if index == 3 && !array.IsStringsEqual(combination, []string{"456", "789"}) {
			t.Fail()
			return
		}
		if index == 4 && !array.IsStringsEqual(combination, []string{"123"}) {
			t.Fail()
			return
		}
		if index == 5 && !array.IsStringsEqual(combination, []string{"456"}) {
			t.Fail()
			return
		}
		if index == 6 && !array.IsStringsEqual(combination, []string{"789"}) {
			t.Fail()
			return
		}
	}
}

// TestGetCombinationsWithImplode tests the function GetCombinationsWithImplode.
func TestGetCombinationsWithImplode(t *testing.T) {
	s := strings.Join(GetCombinationsWithImplode([]string{"123", "456", "789"}, ";"), "/")
	if s != "123;456;789/123;456/123;789/456;789/123/456/789" {
		t.Fail()
		return
	}

	s = strings.Join(GetCombinationsWithImplode([]string{"78", "79", "80", "81", "82"}, ";"), "/")
	if s != "78;79;80;81;82/78;79;80;81/78;79;80;82/78;79;81;82/78;80;81;82/79;80;81;82/78;79;80/78;79;81/78;79;82/78;80;81/78;80;82/78;81;82/79;80;81/79;80;82/79;81;82/80;81;82/78;79/78;80/78;81/78;82/79;80/79;81/79;82/80;81/80;82/81;82/78/79/80/81/82" {
		t.Fail()
		return
	}
}
