package atbash

func Atbash(s string) string {
    var res = make([]byte, 0)
    for i:=0; i<len(s); i++ {
        b := encode(byte(s[i]))
        if b == ' ' {
            continue
        }
        if len(res)%6 == 5 {
            res = append(res, ' ')
        }
        res = append(res, b)
    }
    return string(res)
}

func encode(b byte) byte {
    if b >= 'A' && b <= 'Z' {
        b = byte(b - 'A' + 'a')
    }
    if b >= 'a' && b <= 'z' {
        if b > ('a' + 'z') / 2 {
            b = byte('a' + ('z' - b))
        } else {
            b = byte('z' - (b - 'a'))
        }
        return b
    }
    if b >= '0' && b <= '9' {
        return b
    }
    return ' '
}
