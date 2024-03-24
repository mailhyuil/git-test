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
  hyuil := user {
    name: "hyuil",
    age: 25,
  }
  fmt.Println(hyuil.name)
  fmt.Println(hyuil.age)
}
