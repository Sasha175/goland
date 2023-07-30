package main

import(
"io"
"os"
)


func main() {

var myStr string = ""
argument := os.Args
if len(argument) == 1{
	myStr = "Введите аргумент"
}else{
	myStr = argument[1] + "\n"
}

io.WriteString(os.Stdout, myStr)
myStr = ""
}
