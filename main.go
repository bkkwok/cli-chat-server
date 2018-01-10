package main

import (
  "log"
  "net/http"
  "fmt"
  "justaprank"
  "flag"
)

var (
  env *string
  port *int
)

func init() {
  fmt.Println("init")
  env = flag.String("env", "development", "a string")
  port = flag.Int("port", 3000, "an int")
}

func main() {

  fs := http.FileServer(http.Dir("../public"))
  http.Handle("/", fs)

  p := justaprank.Prank{"benjamin"}

  fmt.Println(p)

  flag.Parse()
  fmt.Println("env: ", *env)
  fmt.Println("port: ", *port)

  log.Fatal(http.ListenAndServe(":8000", nil))

}