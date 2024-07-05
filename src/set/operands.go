package set

func (set *Set[T]) Union(set2 *Set[T]) *Set[T] {
	union := make(Set[T])

	for k := range *set {
		union[k] = struct{}{}
	}

	for k := range *set2 {
		union[k] = struct{}{}
	}

	return &union
}

func (set *Set[T]) Intersection(set2 *Set[T]) *Set[T] {
	intersection := make(Set[T])

	for k := range *set {
		if _, ok := (*set2)[k]; ok {
			intersection[k] = struct{}{}
		}
	}

	return &intersection
}

func (set *Set[T]) Difference(set2 *Set[T]) *Set[T] {
	difference := make(Set[T])

	for k := range *set {
		if _, ok := (*set2)[k]; !ok {
			difference[k] = struct{}{}
		}
	}

	for k := range *set2 {
		if _, ok := (*set)[k]; !ok {
			difference[k] = struct{}{}
		}
	}

	return &difference
}

func (set *Set[T]) IsSubset(set2 *Set[T]) bool {
	for k := range *set {
		if _, ok := (*set2)[k]; !ok {
			return false
		}
	}

	return true
}

func (set *Set[T]) IsSuperset(set2 *Set[T]) bool {
	return set2.IsSubset(set)
}
