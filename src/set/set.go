package set

import "fmt"

type Set[T comparable] map[T]struct{}

func New[T comparable](arr []T) *Set[T] {
	set := make(Set[T])
	for _, x := range arr {
		set[x] = struct{}{}
	}

	return &set
}

func (set *Set[T]) String() string {
	str := "("

	for k := range *set {
		str += fmt.Sprintf("%v ", k)
	}

	return str[:len(str)-1] + ")"
}

func (s *Set[T]) Len() int {
	return len(*s)
}

func (set *Set[T]) Copy() *Set[T] {
	copy := make(Set[T], len(*set))

	for k := range *set {
		copy[k] = struct{}{}
	}

	return &copy
}

func (set *Set[T]) Add(e T) {
	(*set)[e] = struct{}{}
}

func (set *Set[T]) Remove(e T) {
	delete(*set, e)
}

func (set *Set[T]) Contains(e T) bool {
	_, ok := (*set)[e]
	return ok
}

func (set *Set[T]) Clear() {
	clear(*set)
}

func (set *Set[T]) Values() []T {
	arr := make([]T, len(*set))
	i := 0
	for k := range *set {
		arr[i] = k
		i++
	}

	return arr
}
