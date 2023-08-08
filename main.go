package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 1; i < till; i++ {
		fmt.Println(i, message)
	}
}

func main() {
	// jika tidak diberikan maka akan menggunakan default-nya (runtime.NumCPU())
	runtime.GOMAXPROCS(4) // digunakan untuk menentukan jumplah core yang digunakan untuk mengeksekusi program

	// menerapkan goruntine menggunakan keyword "go"
	go print(5, "Hello")
	print(5, "Apa Kabar")

	var input1, input2, input3 string
	fmt.Scanln(&input1, &input2, &input1)
	fmt.Println("input :", input1, input2, input3)
}
