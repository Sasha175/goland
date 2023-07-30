package main
import (
 "errors"
 "fmt"
 "os"
 "strconv"
)

func main() {
	if len(os.Args) == 1 {
	fmt.Println("Please give one or more floats.")
	os.Exit(1)
	}
	arguments := os.Args
	var err error = errors.New("An error")
	k := 1
	//var n float64
	
	for err != nil {
		
		if k >= len(arguments) {
		fmt.Println("None of the arguments is a float!")
		return
		}
		
		_, err = strconv.ParseFloat(arguments[k], 64)
		k++
		}
		//fmt.Println(n)
		sum := 0
		
		for i := 1; i < len(arguments); i++ {
			ns, err := strconv.Atoi(arguments[i])
			if err == nil {
			sum = sum + ns
			}
			if arguments[i] == "END"{
				fmt.Println("sum:", sum)
				return
			}
			}
			
 fmt.Println("sum:", sum)
}	   