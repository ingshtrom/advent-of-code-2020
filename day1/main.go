package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    fmt.Printf("error reading file: %v\n", err)
    os.Exit(1)
  }

  arr := strings.Split(string(data), "\n")
  

  for index, expense := range arr {
    for index2, expense2 := range arr {
      if index != index2 {
        e1, _ := strconv.Atoi(expense)
        e2, _ := strconv.Atoi(expense2)
        if e1 + e2 == 2020 {
          fmt.Printf("The numbers are %d (index: %d) and %d (index %d). Their product is %d\n", e1, index, e2, index2, e1*e2)
          os.Exit(0)
        }
      }
    }
  }
}
