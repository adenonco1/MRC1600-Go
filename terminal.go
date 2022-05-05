package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func testMenu() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("---Peek/Poke Menu---")
	fmt.Println("1) Peek")
	fmt.Println("2) Poke")
	text, _ := reader.ReadString('\n')

	//fmt.Println("the entered text is:", text)
	if strconv.Atoi(text) == 1 {
		peek()
	} else if strconv.Atoi(text) == 2 {
		poke()
	}

}

func peek() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("---Peek---")
	fmt.Println("Peek at address: ")
}

func poke() {
	fmt.Println("---Poke---")
	fmt.Println("Poke to address : ")
	fmt.Println("Poke data to address")
}

func termColor() {

	colorReset := "\033[0m"

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"

	fmt.Println(string(colorRed), "test")
	fmt.Println(string(colorGreen), "test")
	fmt.Println(string(colorYellow), "test")
	fmt.Println(string(colorBlue), "test")
	fmt.Println(string(colorPurple), "test")
	fmt.Println(string(colorWhite), "test")
	fmt.Println(string(colorCyan), "test", string(colorReset))
	fmt.Println("next")
}
