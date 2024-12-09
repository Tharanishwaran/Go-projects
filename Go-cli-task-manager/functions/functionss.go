package main

import "fmt"

func main() {

  fmt.Println("Hello, World!")

  defer add(1,10)

  show()


}

func add(a int,b int){

  
  fmt.Println(a+b)

}

func show() {

  fmt.Println("showed")


}