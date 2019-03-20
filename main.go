package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum(x int, y int) int {
	return x + y
}

func getInput(args []string) (numA int, numB int, err error) {
	if len(args) < 3 {
		return 0, 0, fmt.Errorf("require 2 arguments")
	}
	if numA, err = strconv.Atoi(args[1]); err != nil {
		return 0, 0, err
	}
	if numB, err = strconv.Atoi(args[2]); err != nil {
		return 0, 0, err
	}
	return
}
func main() {
	a, b, e := getInput(os.Args)
	if e != nil {
		panic(e)
	}
	fmt.Println(sum(a, b))
}
