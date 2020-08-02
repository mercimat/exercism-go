package strain

type Ints []int
type Lists [][]int
type Strings []string

func (s Ints) Keep(f func(int) bool) (r Ints) {
    for _, v := range s {
        if f(v) {
            r = append(r, v)
        }
    }
    return
}

func (s Ints) Discard(f func(int) bool) (r Ints) {
    for _, v := range s {
        if !f(v) {
            r = append(r, v)
        }
    }
    return
}

func (s Lists) Keep(f func([]int) bool) (r Lists) {
    for _, v := range s {
        if f(v) {
            r = append(r, v)
        }
    }
    return
}

func (s Strings) Keep(f func(string) bool) (r Strings) {
    for _, v := range s {
        if f(v) {
            r = append(r, v)
        }
    }
    return
}
