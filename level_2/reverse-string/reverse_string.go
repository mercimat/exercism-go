package reverse

func Reverse(s string) string {

    runes := []rune(s)
    for i := 0; i < (len(runes) / 2); i++ {
        runes[i], runes[len(runes)-1 - i] = runes[len(runes)-1 - i], runes[i]
    }
    return string(runes)
}
