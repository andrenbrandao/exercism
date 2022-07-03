package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	if i == nil {
		return nil
	}

	res := make(Ints, 0)

	for _, val := range i {
		if filter(val) {
			res = append(res, val)
		}
	}

	return res
}

func (i Ints) Discard(filter func(int) bool) Ints {
	return i.Keep(func(k int) bool { return !filter(k) })
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	if l == nil {
		return nil
	}

	res := make(Lists, 0)

	for _, val := range l {
		if filter(val) {
			res = append(res, val)
		}
	}

	return res
}

func (s Strings) Keep(filter func(string) bool) Strings {
	if s == nil {
		return nil
	}

	res := make(Strings, 0)

	for _, val := range s {
		if filter(val) {
			res = append(res, val)
		}
	}

	return res
}
