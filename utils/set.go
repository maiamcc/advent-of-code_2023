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
