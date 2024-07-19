package set

import "testing"

type testCase struct {
	name   string
	input  []interface{}
	input2 []interface{}
	output []interface{}
	length int
}

func setEqArr[T comparable](s *Set[T], arr []T) bool {
	if s.Len() != len(arr) {
		return false
	}

	for _, v := range arr {
		if _, ok := (*s)[v]; !ok {
			return false
		}
	}

	return true
}

func TestNew(t *testing.T) {
	testCases := []testCase{
		{
			name:   "(1 2 3 4 5)",
			input:  []interface{}{1, 2, 2, 3, 4, 5, 5, 1, 2},
			output: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:   "(a b ab c abc)",
			input:  []interface{}{"a", "b", "ab", "c", "abc", "a", "c", "ab"},
			output: []interface{}{"a", "b", "ab", "c", "abc"},
		},
		{
			name:   "(1 2 4 8 16 32 64)",
			input:  []interface{}{1, 2, 4, 8, 16, 32, 64},
			output: []interface{}{1, 2, 4, 8, 16, 32, 64},
		},
	}

	for _, tc := range testCases {
		set := New(tc.input)
		if !setEqArr(set, tc.output) {
			t.Errorf("Test case '%s': got %v", tc.name, set)
		}
	}
}

// func TestString(t *testing.T) {
// 	testCases := []testCase{
// 		{
// 			name:  "(1 2 3 4 5)",
// 			input: []interface{}{1, 2, 2, 3, 4, 5, 5, 1, 2},
// 		},
// 		{
// 			name:  "(a b ab c abc)",
// 			input: []interface{}{"a", "b", "ab", "c", "abc", "a", "c", "ab"},
// 		},
// 		{
// 			name:  "(1 2 4 8 16 32 64)",
// 			input: []interface{}{1, 2, 4, 8, 16, 32, 64},
// 		},
// 	}

// 	for _, tc := range testCases {
// 		set := New(tc.input)
// 		if set.String() != tc.name {
// 			t.Errorf("Test case '%s': got %s", tc.name, set)
// 		}
// 	}
// }

func TestLen(t *testing.T) {
	testCases := []testCase{
		{
			name:   "(1 2 3 4 5)",
			input:  []interface{}{1, 2, 2, 3, 4, 5, 5, 1, 2},
			length: 5,
		},
		{
			name:   "(a b ab c abc)",
			input:  []interface{}{"a", "b", "ab", "c", "abc", "a", "c", "ab"},
			length: 5,
		},
		{
			name:   "(1 2 4 8 16 32 64)",
			input:  []interface{}{1, 2, 4, 8, 16, 32, 64},
			length: 7,
		},
	}

	for _, tc := range testCases {
		set := New(tc.input)
		if set.Len() != tc.length {
			t.Errorf("Test case '%s': got %s", tc.name, set)
		}
	}
}

func TestCopy(t *testing.T) {
	testCases := []testCase{
		{
			name:  "(1 2 3 4 5)",
			input: []interface{}{1, 2, 2, 3, 4, 5, 5, 1, 2},
		},
		{
			name:  "(a b ab c abc)",
			input: []interface{}{"a", "b", "ab", "c", "abc", "a", "c", "ab"},
		},
		{
			name:  "(1 2 4 8 16 32 64)",
			input: []interface{}{1, 2, 4, 8, 16, 32, 64},
		},
	}

	for _, tc := range testCases {
		set := New(tc.input)
		copy := set.Copy()

		if set.Len() != copy.Len() {
			t.Errorf("Test case '%s': copy got size %v", tc.name, copy.Len())
		}

		for k := range *set {
			if _, ok := (*copy)[k]; !ok {
				t.Errorf("Test case '%s': %v doesn't contains %v", tc.name, copy, k)
				break
			}
		}
	}
}

func TestAdd(t *testing.T) {
	set := New([]int{})
	if set.Len() > 0 {
		t.Errorf("Error when initializing set : len %v", set.Len())
		return
	}

	set.Add(2)
	set.Add(4)
	set.Add(8)

	if set.Len() != 3 {
		t.Errorf("Error when adding elements : len %v", set.Len())
		return
	}

	set.Add(2)
	set.Add(5)
	set.Add(8)

	if set.Len() != 4 {
		t.Errorf("Error when adding elements : len %v", set.Len())
		return
	}

	if !setEqArr(set, []int{2, 4, 8, 5}) {
		t.Errorf("set doesn't contains correct elements : %v", set)
	}
}

func TestRemove(t *testing.T) {
	set := New([]int{1, 2, 2, 3, 4, 5, 5, 1, 2})
	if set.Len() != 5 {
		t.Errorf("Error when initializing set : len %v", set.Len())
		return
	}

	set.Remove(2)

	if set.Len() != 4 {
		t.Errorf("Error when removing elements : len %v", set.Len())
		return
	}

	set.Remove(2)
	set.Remove(5)
	set.Remove(8)

	if set.Len() != 3 {
		t.Errorf("Error when adding elements : len %v", set.Len())
		return
	}

	if !setEqArr(set, []int{1, 3, 4}) {
		t.Errorf("set doesn't contains correct elements : %v", set)
	}
}

func TestContains(t *testing.T) {
	set := New([]int{1, 2, 2, 3, 4, 5, 5, 1, 2})
	if set.Len() != 5 {
		t.Errorf("Error when initializing set : len %v", set.Len())
		return
	}

	if !set.Contains(1) {
		t.Errorf("set should coutains 1 : %v", set)
	}

	if !set.Contains(2) {
		t.Errorf("set should coutains 2 : %v", set)
	}

	if set.Contains(10) {
		t.Errorf("set should not coutains 10 : %v", set)
	}

	if !set.Contains(1) {
		t.Errorf("set should coutains 1 : %v", set)
	}

	set.Remove(1)

	if set.Contains(15) {
		t.Errorf("set should not coutains 1 anymore : %v", set)
	}
}

func TestClear(t *testing.T) {
	set := New([]int{1, 2, 2, 3, 4, 5, 5, 1, 2})
	if set.Len() != 5 {
		t.Errorf("Error when initializing set : len %v", set.Len())
		return
	}

	set.Clear()

	if set.Len() > 0 {
		t.Errorf("Error clearing set : len %v", set.Len())
	}
}

func TestValues(t *testing.T) {
	set := New([]int{1, 2, 2, 3, 4, 5, 5, 1, 2})
	if set.Len() != 5 {
		t.Errorf("Error when initializing set : len %v", set.Len())
		return
	}

	values := set.Values()

	if set.Len() != len(values) {
		t.Errorf("Values doesn't contains as much elements as the initial set : %v != %v", set.Len(), len(values))
	}

	if !setEqArr(set, values) {
		t.Errorf("Values doesn't contains values from set : %v != %v", set, values)
	}
}
