package prime

func Factors(n int64) []int64 {
    var factors = make([]int64, 0)
    var candidate = int64(2)

    for candidate*candidate <= n {
        for n%candidate == 0 {
            n /= candidate
            factors = append(factors, candidate)
        }
        candidate++
    }
    if n > 1 {
        factors = append(factors, n)
    }
    return factors
}
