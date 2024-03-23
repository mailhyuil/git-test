package main

<<<<<<< HEAD
var baseUrl = "https://www.fmkorea.com/index.php?mid=best&page=5"

func main() {
=======
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
>>>>>>> f0bc622 (chore:update)
}
