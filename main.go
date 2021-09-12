package main

import (
  "net/http"
)
func main() {

    http.HandleFunc("/login",login.html)
    http.ListenAndServe(":8080", nil)
}
