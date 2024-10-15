package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/juniorrodes/arquitetura-computadores-GA/pkg/pipeline"
)

func main() {
	state := pipeline.NewState()

	file, err := os.Open("test.asm")
	if err != nil {
		log.Fatal(err)
	}

	buffer, err := io.ReadAll(file)
	if err = pipeline.ParseInstructions(buffer, state); err != nil {
		log.Fatal(err)
	}

	for _, instruction := range state.InstructionMemory {
		fmt.Println(instruction)
	}

	for {
		state.Fetch()
		state.Decode()
		state.Execute()
		state.MemoryAccess()
		state.WriteBack()
		fmt.Print(state)
		if _, debug := os.LookupEnv("DEBUG"); !debug {
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}
}
