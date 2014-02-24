package main

import (
    "net/http"
    "fmt"
    "crypto/md5"
    "encoding/hex"
    "time"
)

func Hello1(w http.ResponseWriter, req *http.Request) {

    h := md5.New()
    h.Write([]byte("1234"))

    x := h.Sum(nil)

    y := make([]byte, 32)
    hex.Encode(y, x)

    w.Write([]byte(y))
}

func Hello(w http.ResponseWriter, req *http.Request) {
  x := time.Now().Unix()
  s := fmt.Sprintf("%d", x)
  // TODO: add content-type
  w.Write([]byte(s))
}
func main() {
    http.HandleFunc("/", Hello)

    fmt.Println("Server running at http://127.0.0.1:5557/")

    http.ListenAndServe(":5557", nil)
}