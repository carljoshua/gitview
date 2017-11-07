package main

import (
    "io/ioutil"
    "fmt"
    "flag"
    "os"
    "net/http"
)

func main() {
    path := flag.String("f", "", "path to the README file")
    port := flag.String("p", "8000", "port number")
    flag.Parse()

    if *path != "" {
		if _, err := os.Stat(*path); err != nil {
            fmt.Printf("Can't find the file in `%s`\n", *path)
            os.Exit(1)
        }
	}else{
        fmt.Println("Missing file path")
        os.Exit(1)
    }

    var file []byte
    var err error
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        file, err = ioutil.ReadFile(*path)
        if err != nil{
            fmt.Printf("Cannot open the file in %s", *path)
            os.Exit(1)
        }
        tmpl := "<html><head><style>" + getCSS() + "</style></head><body><div class='container'>" +
                toHTML(file) + "</div></body></html>"
        fmt.Fprintf(w, "%s", tmpl)
    })

    fmt.Printf("Listening at port %s...\n", *port)
    erro := http.ListenAndServe(":" + *port, nil)
    if erro != nil {
        fmt.Printf("Invalid port: %s", *port)
    }
}
