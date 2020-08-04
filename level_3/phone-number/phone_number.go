package phonenumber

import (
    "errors"
    "regexp"
)

var pattern = regexp.MustCompile("^\\D*([1]{0,1})\\D*([2-9]\\d{2})\\D*([2-9]\\d{2})\\D*(\\d{4})\\D*$")

func Number(number string) (string, error) {
    var res []byte
    for _, submatch := range pattern.FindAllSubmatch([]byte(number), -1) {
        res = append(res, submatch[2]...)
        res = append(res, submatch[3]...)
        res = append(res, submatch[4]...)
        break
    }
    if len(res) == 0 {
        return "", errors.New("wrong format")
    }
    return string(res), nil
}

func AreaCode(number string) (string, error) {
    for _, submatch := range pattern.FindAllSubmatch([]byte(number), -1) {
        return string(submatch[2]), nil
    }
    return "", errors.New("wrong format")
}

func Format(number string) (string, error) {
    for _, submatch := range pattern.FindAllSubmatch([]byte(number), -1) {
        var res []byte
        res = append(res, '(')
        res = append(res, submatch[2]...)
        res = append(res, ')', ' ')
        res = append(res, submatch[3]...)
        res = append(res, '-')
        res = append(res, submatch[4]...)
        return string(res), nil
    }
    return "", errors.New("wrong format")
}
