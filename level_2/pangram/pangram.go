package pangram

func IsPangram(sentence string) bool {

    //var binArray uint32
    byteArray := [26]uint8{}

    for _, c := range sentence {
        var i int
        if 'A' <= c && c <= 'Z' {
            i = int(c - 'A')
        } else if 'a' <= c && c <= 'z' {
            i = int(c - 'a')
        } else {
            continue
        }
        //binArray |= (1 << i)
        byteArray[i] = 1
    }
    //return binArray == uint32(0x03FFFFFF)

    for j := 0 ; j < len(byteArray); j++ {
        if byteArray[j] == 0 {
            return false
        }
    }
    return true
    //return false
}
