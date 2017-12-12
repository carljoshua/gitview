package main

import (
    "strings"
)

// toHTML converts the file contents into their corresponding HTML tags
func toHTML(file []byte) string {
    var temp string
    lines := strings.Split(string(file), "\n")

    code_toggle := false
    list_toggle := 0
    list := ""

    for _, line := range(lines) {
        if line != ""{
            if len(line) > 1 {
                if strings.TrimSpace(line)[0:2] == "* " {
                    list_toggle = 1
                }else if strings.TrimSpace(line)[0:2] != "* " && list_toggle == 1{
                    temp = temp + "<ul>" + list + "</ul>"
                    list_toggle = 0
                    list = ""
                }
            }

            if line[0:1] == "#"{
                // Adds Header Tags if the line starts with "#"
                temp = temp + headerize(line)
            }else if line[0:1] == ">" {
                // Adds Qoute Tag if the line starts with ">"
                temp = temp + "<blockquote>" + line[1:] + "</blockquote>"
            }else if len(line) > 2 && line[0:3] == "---"{
                // Adds a Line Break if the line starts with "---"
                temp = temp + "<hr />"
            }else if len(line) > 2 && line[0:3] == "```"{
                // Adds <pre> tags if the line starts with "```"
                if !code_toggle {
                    temp = temp + "<pre>"
                    code_toggle = true
                }else{
                    temp = temp + "</pre>"
                    code_toggle = false
                }
            }else if code_toggle {
                temp = temp + line + "\n"
            }else{
                line = applyStyle(line)

                if list_toggle == 1 {
                    list = list + "<li>" + strings.TrimSpace(line)[2:] + "</li>"
                }else{
                    temp = temp + "<p>" + line + "</p>"
                }
            }
        }
    }

    return temp
}
