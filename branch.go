package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 70:
		g = "C"
	case score == 100:
		g = "A"
	default:
		panic(fmt.Sprint("Wraning"))
	}
	return g
}
func main() {
	const filename = "a.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	// contents 只在if块中有效
	if contents, err := ioutil.ReadFile(filename); err != nil {
		println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	println(
		grade(100),
	)
}
