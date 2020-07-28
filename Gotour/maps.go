package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var nil_m map[string]Vertex //nil map
// key type: string, value type: vertex

func main() {
  m := make(map[string]Vertex)
  m["Bell Labs"] = Vertex{
    40.68433, -74.39967,
  }
  m["My Labs"] = Vertex{
    1.111, -2.222,
  }
/*
  nil_m["Nil Labs"] = Vertex{
    1.111, -2.222,
  }
*/
  //panic: assignment to entry in nil map

  fmt.Println(m["My Labs"])
}
