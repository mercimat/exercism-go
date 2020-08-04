package armstrong

import (
    "math"
)

// Tried creating a slice for each digit -> too slow and using too much memory
// Tried converting the number to a string using strconv.Itoa -> a bit slower
// than the suggested solution

func order(n int) (result int) {
    for n != 0 {
        result++
        n /= 10
    }
    return
}

func IsNumber(n int) bool {
    originalNumber := n
    pow := order(n)
    sum := 0
    for n != 0 {
        remainder := n % 10
        sum += int(math.Pow(float64(remainder), float64(pow)))
        n /= 10
    }
    return originalNumber == sum
}
