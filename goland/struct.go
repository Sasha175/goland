package main
import "fmt"
type Human struct{
Name string
age uint16
Weight uint16
}

func main() {

  me := Human {
	Name: "Artur",
	age: 32,
	Weight: 110,
   }

   fmt.Println(me)
}
