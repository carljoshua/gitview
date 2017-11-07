package main

import (
    "strconv"
    "strings"
)

// Creates a header tag with the same number as the "#" (e.g. "###" = <h3></h3>)
func headerize(raw string) string {
    hs := 0
    n := 0
    for _, c := range(raw){
        if c != '#' && c != ' '{
            break
        }else if c == ' '{
            n++
        }else if c == '#'{
            hs++
            n = hs
        }
    }

    name := raw[n:]

    return "<h" + strconv.Itoa(hs) + " id='" + strings.ToLower(name) + "'>" + name + "</h" + strconv.Itoa(hs) + ">"
}

// Creates an <a> tag
func makeAnchor(href string, alt string) string {
    if alt != "" {
        return "<a href='" + href + "'>" + alt + "</a>"
    }
    return "<a href='" + href + "'>" + href + "</a>"
}

// Adds <strong>, <em>, <a>, <del> and <code> tags in the lines
func applyStyle(line string) string {
    // URL Parser
    url, alt := extractAnchorData(line)
    line = strings.Replace(line, "[" + alt + "](" + url + ")", makeAnchor(url, alt), -1)

    // Simple URL Parser
    for _, words := range(getURL(line)){
        line = strings.Replace(line, words, makeAnchor(words, ""), -1)
    }

    // Bold Parser
    for _, words := range(extract(line, "**")){
        line = strings.Replace(line, "**" + words + "**", "<strong>" + words + "</strong>", -1)
    }
    for _, words := range(extract(line, "__")){
        line = strings.Replace(line, "__" + words + "__", "<strong>" + words + "</strong>", -1)
    }

    // Italic Parser
    for _, words := range(extract(line, "*")){
        line = strings.Replace(line, "*" + words + "*", "<em>" + words + "</em>", -1)
    }
    for _, words := range(extract(line, "_")){
        line = strings.Replace(line, "_" + words + "_", "<em>" + words + "</em>", -1)
    }

    // Strikethrough Parser
    for _, words := range(extract(line, "~~")){
        line = strings.Replace(line, "~~" + words + "~~", "<del>" + words + "</del>", -1)
    }

    // Code Parser
    for _, words := range(extract(line, "`")){
        line = strings.Replace(line, "`" + words + "`", "<code>" + words + "</code>", -1)
    }

    return line
}
