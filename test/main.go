package main

import (
	fu "filuap/lib"
	"fmt"
)

func main() {
	file := fu.ReadFile("./test/main.go")

	data := fu.Data(file)

	msg, err := fu.WriteDataFile(data, "mainCopy.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(msg)

}
