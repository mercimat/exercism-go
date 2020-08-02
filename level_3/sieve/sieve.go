package sieve

func Sieve(limit int) (res []int) {
    if limit < 2 {
        return res
    }

    multiples := make([]bool, limit+1, limit+1)

    for i:=2; i<=limit; i++ {
        if multiples[i] {
            continue
        }
        res = append(res, i)

        for j:=i+i; j <= limit; j+=i {
            multiples[j] = true
        }
    }
    return
}
