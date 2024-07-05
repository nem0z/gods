package main

import (
	"fmt"

	"github.com/nem0z/gods/set"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 3, 2, 2, 3, 1, 5}
	arr2 := []int{2, 2, 4, 4, 7, 8, 2, 1, 1, 2}

	set1 := set.New(arr1)
	set2 := set.New(arr2)

	fmt.Println(set1)
	fmt.Println(set2)

	fmt.Println(set1.Intersection(set2))
	fmt.Println(set1.Union(set2))
	fmt.Println(set1.Difference(set2))
	fmt.Println(set2.Difference(set1))

	fmt.Println(set1.IsSubset(set2))
	fmt.Println(set1.IsSuperset(set2))
}
