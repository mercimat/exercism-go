package flatten

func Flatten(input interface{}) []interface{} {
    q := make([]interface{}, 0)
    flatten(input, &q)
    return q
}

func flatten(input interface{}, q *[]interface{}) {
    if t, ok := input.([]interface{}); ok {
        for _, v := range t {
            flatten(v, q)
        }
    } else if input != nil {
        *q = append(*q, input)
    }
}
