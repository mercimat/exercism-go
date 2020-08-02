package lsproduct

import "fmt"

func LargestSeriesProduct(digits string, span int) (int64, error) {
    res := int64(0)
    if span > len(digits) {
        return res, fmt.Errorf("%s is shorter than %d", digits, span)
    }
    if span < 0 {
        return res, fmt.Errorf("%d must be gt 0", span)
    }

    for i:=0; i <= len(digits) - span; i++ {
        product := int64(1)
        for j:=0; j < span; j++ {
            if digits[i+j] < '0' || digits[i+j] > '9' {
                return 0, fmt.Errorf("%s has unexpected characters", digits)
            }
            product *= int64(digits[i+j] - '0')
        }
        if product > res {
            res = product
        }
    }
    return res, nil
}
