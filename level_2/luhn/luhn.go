package luhn

func Valid(s string) bool {

    d := make([]int, 0, len(s))

    for _, c := range s {
        if c == ' ' {
            continue
        }
        if ! ('0' <= c && c <= '9') {
            return false
        }
        d = append(d, int(c - '0'))
    }
    if len(d) < 2 {
        return false
    }


    return (LuhnSum(d) % 10) == 0
}

func LuhnSum(d []int) int {
    sum := 0
    for i, digit := range d {
        j := len(d) -i
        if j % 2 == 0 {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }
        sum += digit
    }
    return sum
}
