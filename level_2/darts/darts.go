package darts


func Score(x, y float64) int {
    radSquarred := x*x + y*y

    switch {
    case radSquarred <= 1.0:
        return 10
    case radSquarred <= 25.0:
        return 5
    case radSquarred <= 100.0:
        return 1
    default:
        return 0
    }
}
