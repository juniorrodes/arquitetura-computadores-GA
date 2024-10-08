package main

import (
	"fmt"
	"log"
	"os"

	"github.com/juniorrodes/arquitetura-computadores-GA/pkg/instructions"
)

type State struct {
    Pc int
    InstructionMemory []instructions.MemInstruction
    Registers [32]int
}

func main() {
    state := Init()

    for i, instruction := range state.InstructionMemory {
        fmt.Println("instruction ", i, ": ")
        fmt.Println(instruction)
    } 
}

func Init() State {
    fileContent, err := os.ReadFile("test.asm")
    if err != nil {
        log.Fatal(err) 
    }
    
    i, err := instructions.ParseInstructions(fileContent)
    if err != nil {
        log.Fatal(err)
    }

    return State{
        Pc: 0,
        InstructionMemory: i,
    }
}
