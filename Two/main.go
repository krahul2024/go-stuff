package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	output := "This is some string in the output.txt file."
	file, err := os.Create("./output.txt")
	if err != nil {
		fmt.Println("There was an error")
		panic(err)
	}
	length, err := io.WriteString(file, output)
	if err != nil {
		fmt.Println("There was an error")
		panic(err)
	}

	fmt.Println("The size of file is : ", length)
	defer file.Close()

	readFile("output.txt")
}

func readFile(filename string) {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error")
		panic(err)
	}
	// fmt.Println(buffer) // this would print ascii for each character in file present in decimal format : [84 104 105 115 32 105 115 32 115 111 109 101 32 115 116 114 105 110 103 32 105 110 32 116 104 101 32 111 117 116 112 117 116 46 116 120 116 32 102 105 108 101 46]
	content := ""
	for i := range buffer {
		content += string(buffer[i])
	}
	fmt.Println(content)
}
