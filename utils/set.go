package utils

type IntSet map[int]struct{}

func NewIntSet(ints []int) IntSet {
	is := make(IntSet)
	for _, i := range ints {
		is.Add(i)
	}
	return is
}

func (is IntSet) Contains(i int) bool {
	_, ok := is[i]
	return ok
}

func (is IntSet) Add(i int) {
	is[i] = struct{}{}
}

func (is IntSet) Rm(i int) {
	delete(is, i)
}

type Set[E comparable] map[E]struct{}

func (s Set[E]) Contains(v E) bool {
	_, ok := s[v]
	return ok
}
func (s Set[E]) Add(v E) {
	s[v] = struct{}{}
}

func (s Set[E]) Rm(v E) {
	delete(s, v)
}

func NewSet[E comparable](vals ...E) Set[E] {
	set := Set[E]{}
	for _, v := range vals {
		set[v] = struct{}{}
	}
	return set
}
