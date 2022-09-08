package main

import "fmt"

func Run(sut ...int) {
	fmt.Println(len(sut))
}

func main() {
	Run()

}
