package main

import "fmt"

type user struct {
  name string
  age int
  posts []post
}

type post struct {
  title string
  content string
}

func main(){
  sb := user {
    name: "sb",
    age: 25,
  }
  fmt.Println(sb.name)
  fmt.Println(sb.age)
}
