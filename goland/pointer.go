package main
import "fmt"
type Human struct{
	Name string
	age uint16
	Weight uint16
	}

func main()  {
	name := "Artur"
	changeName(&name)
	fmt.Println(name)

	h:=newHumen("Alexsandr", 30, 100)

	fmt.Println(h)
	h.changeNHumen("Alexsandr", 35, 100)
	fmt.Println(h)

}

func changeName(name *string){

*name = *name + "_new"

}

func (h *Human) changeNHumen(Name string, age uint16, Weight uint16) {

	h.Name = Name
	h.age = age 
	h.Weight = Weight
}

func newHumen(newName string, newage uint16, newWeight uint16) Human{
	return Human{
		Name:newName,
		age: newage,
		Weight: newWeight,
	}
}