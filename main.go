package main

import (
	"io/ioutil"

	"github.com/franciscocid/vm/chip8"
)

func main() {

	// g := chip8.NewAsciiGraphics(64, 32, "X", " ")
	program, err := ioutil.ReadFile("./programs/random_number_test.ch8")
	if err != nil {
		panic(err)
	}

	c8 := chip8.New()
	c8.LoadProgram(program)
	c8.Run()
}
