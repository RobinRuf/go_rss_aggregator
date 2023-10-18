package main

import (
  "fmt"
  "os"
  "log"
) 

func main() {
  portString := os.Getenv("PORT")
  if portString == "" {
    log.Fatal("PORT is not found in the environment")
  }

  fmt.Println("PORT:", portString)
}
