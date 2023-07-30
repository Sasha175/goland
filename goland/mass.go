package main
import (

 "fmt"

)

func main() {
var arr [] int
all :=  [] string{"A", "B", "C"}

for i := 1; i<10; i++{
	arr = append(arr,i)
}

for index, value := range all{
    fmt.Println(index, value)
}


arr[0] = 10
fmt.Println(arr)
m := map[int]string{}
m[1] = "Понедельник"
m[2] = "Вторник"
m[3] = "Среда"
m[4] = "Четверг"
m[5] = "Пятница"
m[6] = "Суббота"
fmt.Println(m)



for k, v := range m {

	fmt.Print( "k: ", k)
	fmt.Print( " v: ", v)
	fmt.Println()
}

}