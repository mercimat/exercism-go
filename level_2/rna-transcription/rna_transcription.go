package strand

func ToRNA(dna string) string {
    rnaBytes := []rune(dna)
    for i := 0; i < len(rnaBytes); i++ {
        switch rnaBytes[i] {
        case 'G':
            rnaBytes[i] = 'C'
        case 'C':
            rnaBytes[i] = 'G'
        case 'T':
            rnaBytes[i] = 'A'
        case 'A':
            rnaBytes[i] = 'U'
        }
    }
    return string(rnaBytes)
}

