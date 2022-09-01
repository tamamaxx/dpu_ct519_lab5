package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

// func index(w http.ResponseWriter, req *http.Request){
// 	p := "static/index.html"
// 	http.ServeFile(w, req, p)
// }

// func about(w http.ResponseWriter, req *http.Request){
//     p := "static/about.html"
//     http.ServeFile(w, req, p)
// }

// func myresearch(w http.ResponseWriter, req *http.Request){
//     p := "static/myresearch.html"
//     http.ServeFile(w, req, p)
// }

func main() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    // http.HandleFunc("/index", index)
    // http.HandleFunc("/about", about)
    // http.HandleFunc("/myresearch", myresearch)
    http.Handle("/about.html", http.FileServer(http.Dir("./www")))
    http.Handle("/myresearch.html", http.FileServer(http.Dir("./www")))
    http.Handle("/", http.FileServer(http.Dir("./www")))

    http.ListenAndServe(":8090", nil)
}
