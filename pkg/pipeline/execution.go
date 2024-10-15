package pipeline

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//import "strconv"

const (
    ADD = "ADD"
    SUB = "SUB"
)

func (s *State) Fetch() {
    instruction := s.InstructionMemory[s.Pc]
    if instruction.OpCode != "halt" {
        fmt.Println("fetching: ", instruction)
        s.Pc += 1
        execInstruction := ExecInstructions {
            MemInstruction: instruction,
        }
        s.queue.Push(&execInstruction)
        return
    }
    os.Exit(0)
}

func (s *State) Decode() {
     if s.queue[1] != nil {
        instruction := s.queue[1]
        fmt.Println("Decoding: ", instruction)
        switch instruction.OpCode {
        case "halt", "noop":
            return
        default:
            operandToTemp(s, instruction, 1)
            operandToTemp(s, instruction, 2)
            operandToTemp(s, instruction, 3)
        }
    }   
}

func (s *State) Execute() {
    if s.queue[2] != nil && !s.queue[2].Invalid {
        instruction := s.queue[2]
        fmt.Println("Executing: ", instruction)
        switch instruction.OpCode {
        case "lw", "sw":
            instruction.Temp1 = instruction.Temp1 + instruction.Temp3
        case "add", "sub":
            instruction.Temp1 = alu(
                s.Registers[instruction.Temp1],
                s.Registers[instruction.Temp2],
                strings.ToUpper(instruction.OpCode),
            )
        case "beq":
            if s.Registers[instruction.Temp1] == s.Registers[instruction.Temp2] {
                s.queue[0].Invalid = true
                s.queue[1].Invalid = true
                s.Pc = instruction.Temp3
            }
        default:
            return
        }
    }
}

func (s *State) MemoryAccess() {
    if s.queue[3] != nil && !s.queue[3].Invalid {
        instruction := s.queue[3]
        fmt.Println("Memory access for: ", instruction)
        switch instruction.OpCode {
        case "lw":
            instruction.Temp3 = *s.MainMemory[instruction.Temp1]
        case "sw":
            val := s.Registers[instruction.Temp2]
            s.MainMemory[instruction.Temp1] = &val
        }
    }
}

func (s *State) WriteBack() {
    if s.queue[4] != nil && !s.queue[4].Invalid {
        instruction := s.queue[4]
        fmt.Println("Write back for: ", instruction)
        switch instruction.OpCode {
        case "lw":
            s.Registers[instruction.Temp2] = instruction.Temp3
        case "add", "sub":
            s.Registers[instruction.Temp3] = instruction.Temp1
        }
    }
}

func alu(valA, valB int, operation string) int {
    switch operation {
    case ADD:
        return valA + valB
    case SUB:
        return valA - valB
    }
    return 0
}

func operandToTemp(state *State, instruction *ExecInstructions, index int) {
    switch index{
    case 1:
        v, err := strconv.Atoi(instruction.Operand1)
        if err != nil {
            if v, ok := state.labels[instruction.Operand1]; ok {
                instruction.Temp1 = v
                return
            } else {
                log.Fatal(err)
            }
        }
        instruction.Temp1 = v
    case 2:
        v, err := strconv.Atoi(instruction.Operand2)
        if err != nil {
            if v, ok := state.labels[instruction.Operand2]; ok {
                instruction.Temp2 = v
                return
            } else {
                log.Fatal(err)
            }
        }
        instruction.Temp2 = v
    case 3:
        v, err := strconv.Atoi(instruction.Operand3)
        if err != nil {
            if v, ok := state.labels[instruction.Operand3]; ok {
                instruction.Temp3 = v
                return
            } else {
                log.Fatal(err)
            }
        }
        instruction.Temp3 = v
    }
}
