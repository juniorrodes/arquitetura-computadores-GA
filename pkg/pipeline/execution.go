package pipeline

import (
	"log"
	"strconv"
)

//import "strconv"


func (s *State) Fetch() {
    instruction := s.InstructionMemory[s.Pc]
    s.Pc += 1
    execInstruction := ExecInstructions {
        MemInstruction: instruction,
    }
    s.queue.Push(&execInstruction)
}

func (s *State) Decode() {
     if s.queue[1] != nil {
        instruction := s.queue[1]
        switch instruction.OpCode {
            case "lw":
                operandToTemp(s, instruction, 1)
                operandToTemp(s, instruction, 2)
                operandToTemp(s, instruction, 3)
        }
    }   
}

func (s *State) Execute() {
    if s.queue[2] != nil {
        instruction := s.queue[2]
        switch instruction.OpCode {
            case "lw":
                instruction.Temp1 = instruction.Temp1 + instruction.Temp3
        }
    }
}

func (s *State) MemoryAccess() {
    if s.queue[3] != nil {
        instruction := s.queue[3]
        switch instruction.OpCode {
            case "lw":
                instruction.Temp3 = *s.MainMemory[instruction.Temp1]
        }
    }
}

func (s *State) WriteBack() {
    if s.queue[4] != nil {
        instruction := s.queue[4]
        switch instruction.OpCode {
            case "lw":
                s.Registers[instruction.Temp2] = instruction.Temp3
        }
    }
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
