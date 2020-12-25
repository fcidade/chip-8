package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/franciscocid/vm/chip8"
)

func main() {

	log.SetOutput(ioutil.Discard)

	rand.Seed(time.Now().UnixNano())

	program, err := ioutil.ReadFile("./programs/random_number_test.ch8")
	if err != nil {
		panic(err)
	}

	// g := chip8.NewAsciiMonitor(64, 32, "█", ".")
	g := chip8.NewGuiMonitor(64, 32)
	c8 := chip8.NewChip8(&g)
	c8.LoadProgram(program)
	c8.Run()
}
