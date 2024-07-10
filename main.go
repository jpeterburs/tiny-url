package main

import (
  "errors"
  "fmt"
  "io"
  "net/http"
  "os"
)

func main() {
  http.HandleFunc("/", getRoot)

  err := http.ListenAndServe(":3333", nil)
  if errors.Is(err, http.ErrServerClosed) {
    fmt.Println("Server closed")
  } else if err != nil {
    fmt.Printf("Error starting server: %s\n", err)
    os.Exit(1)
  }
}

func getRoot(w http.ResponseWriter, r *http.Request) {
  fmt.Println("GET /")
  io.WriteString(w, "Hello, World!")
}