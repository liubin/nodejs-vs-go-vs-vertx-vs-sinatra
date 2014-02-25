package main

import (
    "net/http"
    "fmt"
    "crypto/md5"
    "encoding/hex"
    "time"
)

func Hello(w http.ResponseWriter, req *http.Request) {

    s := fmt.Sprintf("%d", time.Now().Unix())

    h := md5.New()
    h.Write([]byte(s))

    x := h.Sum(nil)

    y := make([]byte, 32)

    hex.Encode(y, x)

    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte(y))
}

func main() {
    http.HandleFunc("/", Hello)

    fmt.Println("Server running at http://127.0.0.1:5557/")

    http.ListenAndServe(":5557", nil)
}