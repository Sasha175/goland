package main
import (
	"fmt"
	"math/rand"
	"time"
	"packGo"

	
)
//sec1 := rand.New(rand.NewSource(10))
//fmt.Println( sec1)

func main() {
	//var arr = [...]int{2,4,6,1,3,5,7,9,8}
	var arr [] int
	for i := 1; i<100000; i++{
		arr = append(arr, rand.Intn(100))
	}


	// for _, value := range arr{
	// 	fmt.Println( value)
	// }
	fmt.Println("----------------------------")
	l:=0
	t0 := time.Now();
	for i:=0; i<len(arr);i++{		
		for j:=i+1; j<(len(arr));j++{
			if arr[i] > arr[j]{
				l = arr[i]
				arr[i] = arr[j]
				arr[j] = l
			}
		}

			}
			fmt.Println("End")		
	// for _, value := range arr{
	// 	fmt.Println( value)

	// }

	t1 := time.Now();
    fmt.Printf("Elapsed time: %v", t1.Sub(t0));

	packGop.Print("")
}
