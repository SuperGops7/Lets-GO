package main

import "fmt"

func swap(a,b int) (int,int){
  return b,a
}

func main(){
  a := 100
  b := 200
  fmt.Println("Before swapping, a=", a, " b=", b)
  a,b = swap(a,b)
  fmt.Println("After swapping, a=", a, " b=", b)
}
