package main

import (
    "strings"
)

func toHTML(file []byte) string {
    var temp string
    lines := strings.Split(string(file), "\n")

    for _, line := range(lines) {
        if line != ""{
            switch string(line[0]) {
            case "#":
                temp = temp + headerize(line)
            default:
                // Bold Parser
                for _, words := range(trim(line, "**")){
                    line = strings.Replace(line, "**" + words + "**", "<strong>" + words + "</strong>", -1)
                }
                for _, words := range(trim(line, "__")){
                    line = strings.Replace(line, "__" + words + "__", "<strong>" + words + "</strong>", -1)
                }

                // Italic Parser
                for _, words := range(trim(line, "*")){
                    line = strings.Replace(line, "*" + words + "*", "<em>" + words + "</em>", -1)
                }
                for _, words := range(trim(line, "_")){
                    line = strings.Replace(line, "_" + words + "_", "<em>" + words + "</em>", -1)
                }

                // Strikethrough Parser
                for _, words := range(trim(line, "~~")){
                    line = strings.Replace(line, "~~" + words + "~~", "<del>" + words + "</del>", -1)
                }

                temp = temp + "<p>" + line + "</p>"
            }
        }
    }

    return temp
}
