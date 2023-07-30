package main

import "fmt"

func finish() {
	fmt.Println("Program has been finished")
	//panic("Error")
}

func main() {
	defer finish()
	var users = []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}

	var hello string
	hello = "Hello World!"
	fmt.Println(hello)
	var i int = 6
	//fmt.Scan(&i)

	if i < 0 {
		fmt.Println("i меньше 0")
	} else if i > 0 {
		fmt.Println("i больше 0")
	} else {
		fmt.Println("i = 0")
	}

	users = append(users, "Alexandr")
	for index, value := range users {
		fmt.Println(fmt.Sprint(index) + ")" + value)
	}
	//удаляем 4-й элемент
	var n = 3
	users = append(users[:n], users[n+1:]...)

	u1 := users[2:6]
	fmt.Println(u1)

	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}

	//Добавить элемент
	people["Kate"] = 128
	//Удалить элемент
	delete(people, "Bob")

	fmt.Println(people["Alice"]) // 8
	fmt.Println(people["Bob"])   // 2
	people["Bob"] = 32
	fmt.Println(people["Bob"])

	var x int = 4
	var p *int = &x            // указатель получает адрес переменной
	fmt.Println("Address:", p) // значение указателя - адрес переменной x
	fmt.Println("Value:", *p)  // значение переменной x

	for i := 0; i < 5; i++ {
		fmt.Println("итерация " + fmt.Sprint(i))
	}
}
