package main


func trim(raw string, cutset string) []string  {
    csl := len(cutset)
    pre := 0
    trimmed := ""
    var trimmed_set []string

    for i, c := range(raw)[:len(raw) - csl] {
        if raw[i:i + csl] == cutset {
            pre = pre + 1
            if pre == 2 {
                trimmed_set = append(trimmed_set, trimmed[csl:])
                trimmed = ""
                pre = 0
            }
        }

        if pre == 1{
            trimmed = trimmed + string(c)
        }
    }

    return trimmed_set
}
