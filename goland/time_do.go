package main
import (
 
 "fmt"
 "time"
 
)


func main() {
t0 := time.Now();
n := 1
	for i := 1; i < 1000000000; i++ {
		n = n + i

	}
	fmt.Println(n)
	t1 := time.Now();
    fmt.Printf("Elapsed time: %v", t1.Sub(t0));
}
