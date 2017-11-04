package main

import (
    "strings"
    "strconv"
)

func trim(raw string, cutset string) []string  {
    csl := len(cutset)
    pre := 0
    trimmed := ""
    raw = raw + "<<"
    var trimmed_set []string

    for i, c := range(raw[:len(raw) - csl]) {
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

func headerize(raw string) string {
    hs := 0
    for _, c := range(raw){
        if c != '#'{
            break
        }
        hs++
    }

    return "<h" + strconv.Itoa(hs) + ">" + strings.Replace(raw, "#", "", -1) + "</h" + strconv.Itoa(hs) + ">"
}
