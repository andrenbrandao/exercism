package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) (out Ints) {
	for _, val := range i {
		if filter(val) {
			out = append(out, val)
		}
	}

	return out
}

func (i Ints) Discard(filter func(int) bool) Ints {
	return i.Keep(func(k int) bool { return !filter(k) })
}

func (l Lists) Keep(filter func([]int) bool) (out Lists) {
	for _, val := range l {
		if filter(val) {
			out = append(out, val)
		}
	}

	return out
}

func (s Strings) Keep(filter func(string) bool) (out Strings) {
	for _, val := range s {
		if filter(val) {
			out = append(out, val)
		}
	}

	return out
}
