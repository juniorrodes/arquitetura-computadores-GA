package pipeline

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Labels map[string]int

var (
    SUPPORTED_INSTRUCTIONS = []string{ "add", "sum", "beq", "lw", "sw", "noop", "halt" }
    labels Labels = make(Labels)
)

type MemInstruction struct {
    OpCode string
    Operand1 string
    Operand2 string
    Operand3 string
}

func ParseInstructions(buffer []byte, state *State) (err error) {
    str := string(buffer)
    str = strings.TrimRight(str, "\n")
    lines :=  strings.Split(str, "\n")
    
    identifyLabels(lines, state)
    for _, line := range lines {
        line = strings.TrimLeft(line, " ")
        lineParsed := strings.Split(line, " ")
        var instruction MemInstruction
        if _, ok := labels[lineParsed[0]]; !ok {
            instruction, err = newInstruction(lineParsed)
            if err != nil {
                return
            }
            state.InstructionMemory = append(state.InstructionMemory, instruction)
        }
     }
     state.labels = labels
    return
}

func identifyLabels(lines []string, state *State) error {
    for i:= 0; i < len(lines); i++ {
        lineParsed := strings.Split(lines[i], " ")
        if !slices.Contains(SUPPORTED_INSTRUCTIONS, lineParsed[0]) {
            if len(lineParsed) == 3 && lineParsed[1] == ".fill" {
                i, err := findFirstNilIndex(state.MainMemory[:])
                if err != nil {
                    return errors.New("Memory is full")
                }
                value, err := strconv.Atoi(lineParsed[2])
                if err != nil {
                    return err
                }
                state.MainMemory[i] = &value
                labels[lineParsed[0]] = i
            }
            if _, ok := labels[lineParsed[0]]; !ok {
                labels[lineParsed[0]] = i
                lines[i] = strings.ReplaceAll(lines[i], lineParsed[0], "")
            }
        }
    }
    return nil
}

func findFirstNilIndex(slice []*int) (int, error) {
    for i := 0; i < len(slice); i++ {
        if slice[i] == nil {
             return i, nil
        }
    }
    return 0, errors.New("Not found")
}

func newInstruction(line []string) (MemInstruction, error) {
    var err error
    
    if !slices.Contains(SUPPORTED_INSTRUCTIONS[4:], line[0]) {
        for i := 1; i < 4; i++ {
            _, err = strconv.Atoi(line[i])
            if err != nil {
                if _, ok := labels[line[i]]; ok {
                    continue
                } else {
                    return MemInstruction{}, errors.New(fmt.Sprint("Did not found matching label: ", line[i]))
                }
            }
        }
    }
    var instruction MemInstruction
    if len(line) > 1 {
        instruction = MemInstruction{
            OpCode: line[0],
            Operand1: line[1],
            Operand2: line[2],
            Operand3: line[3],
        }
    } else {
        instruction = MemInstruction{
            OpCode: line[0],
        }
    }

    return instruction, nil
} 

func (m *MemInstruction) String() string {
        return fmt.Sprintf("{\"opCode\": %s,\"operand1\": %s,\"operand2\": %s,\"operand3\": %s}", m.OpCode, m.Operand1, m.Operand2, m.Operand3)
}
