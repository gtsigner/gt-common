package assets

import (
    "fmt"
    "net/http"
    "testing"
    _ "zhaojunlike/snkrs/assets/server/statik"
)

func TestLoad(t *testing.T) {
    // Serve the contents over HTTP.
    err := Load("./server")
    fmt.Println(err)
    res, err := FileSystem.Open("./static/a.html")
    fmt.Println("JELL", res, err)
    http.Handle("/", http.StripPrefix("/", http.FileServer(FileSystem)))
    http.ListenAndServe(":8080", nil)
}
