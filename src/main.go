package main

import (
  "net/http"
   "os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
  message, _ := os.Hostname()
  message = message + "\n"
  w.Write([]byte(message))
}

func main() {
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":80", nil); err != nil {
    panic(err)
  }
}
