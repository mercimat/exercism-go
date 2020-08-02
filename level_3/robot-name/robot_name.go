package robotname

import (
    "fmt"
    "crypto/rand"
    "math/big"
)

type Robot struct {
    name string
}

var spaceSize = 26*26*10*10*10
var availableIDs = make([]int, spaceSize)

func init() {
    for i:=0; i<len(availableIDs); i++ {
        availableIDs[i] = i
    }
}

func randInt() int {
    max := int64(len(availableIDs))
    nBig, _ := rand.Int(rand.Reader, big.NewInt(max))
    return int(nBig.Int64())
}

func getUniqueID() int {
    i := randInt()
    id := availableIDs[i]
    availableIDs[i] = availableIDs[len(availableIDs)-1]
    availableIDs = availableIDs[:len(availableIDs)-1]
    return id
}

func generateName() string {
    id := getUniqueID()
    name := make([]rune, 5)
    remains := id
    div := spaceSize
    for i:=0; i<5; i++ {
        size := 10
        r := int('0')
        if i < 2 {
            size = 26
            r = int('A')
        }
        div /= size
        elem := remains / div
        remains %= div
        name[i] = rune(r + elem)
    }
    return string(name)
}

func (r *Robot) Reset() {
    r.name = ""
}

func (r *Robot) Name() (string, error) {
    if r.name != "" {
        return r.name, nil
    }
    if len(availableIDs) < 1 {
        return "", fmt.Errorf("no more name available")
    }
    newName := generateName()
    r.name = newName
    return newName, nil
}
