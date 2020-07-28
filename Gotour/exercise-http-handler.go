package main

import (
  "fmt"
  "net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, s)
}

type Struct struct {
  Greeting string
  Punch string
  Who string
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, s)
}

func main() {
  http.Handle("/string", String("I am a frayed knot."))
  http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
  http.ListenAndServe("localhost:4000", nil)
}
