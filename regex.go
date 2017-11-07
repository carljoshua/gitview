package main

import (
    "regexp"
)

// Returns a slice of string that are in between the cutset
// e.g.     extract("This is **just** a test **of the** extract function", "**")
//          returns [ just, of the ]
func extract(raw string, cutset string) []string  {
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

// Returns url inside the line
func getURL(raw string) []string {
    url := regexp.MustCompile("(http|https)://+[a-zA-Z0-9:#@%/;$~_?\\+-=&]*")
    return url.FindAllString(raw, -1)
}

// Returns the url and the alternative text
func extractAnchorData(raw string) (string, string) {
    var url, alt string
    url_toggle := false
    alt_toggle := false

    for _, c := range(raw) {
        if c == '[' {
            alt_toggle = true
        }else if c == ']' && alt_toggle {
            alt_toggle = false
        }else if alt_toggle {
            alt = alt + string(c)
        }

        if c == '(' {
            url_toggle = true
        }else if c == ')' && url_toggle {
            url_toggle = false
        }else if url_toggle {
            url = url + string(c)
        }
    }

    return url, alt
}
