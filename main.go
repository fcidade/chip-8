package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"time"
	"runtime"

	"github.com/franciscocid/chip8/chip8"
)

func setup() {
	runtime.LockOSThread()

	// log.SetOutput(ioutil.Discard)
	log.Println("Starting...")

	rand.Seed(time.Now().UnixNano())
}

func main() {
	setup()

	// program, err := ioutil.ReadFile("./programs/random_number_test.ch8")
	program, err := ioutil.ReadFile("./programs/test_opcode.ch8")
	if err != nil {
		panic(err)
	}

	g := chip8.NewGuiMonitor(600, 400)
	c8 := chip8.NewChip8(g)

	c8.LoadProgram(program)
	c8.LoadFonts()

	g.Update(func() {
		c8.Tick()
		// time.Sleep(time.Second / 5)
		time.Sleep(time.Second / 60)
	})
}
