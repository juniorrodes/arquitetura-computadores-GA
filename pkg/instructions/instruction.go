package instructions

import "slices"

const SUPPORTED_INSTRUCTIONS = []string{ "add", "sum", "beq", "lw", "sw", "noop", "halt" }

import (
	"fmt"
	"strconv"
	"strings"
)

type MemInstruction struct {
    OpCode string
    Operand1 int
    Operand2 int
    Operand3 int
}

func ParseInstructions(buffer []byte) (instructions []MemInstruction, err error) {
    str := string(buffer)
    str = strings.TrimRight(str, "\n")
    lines :=  strings.Split(str, "\n")

    for _, line := range lines {
        lineParsed := strings.Split(line, " ")
        if slices.Contains(SUPPORTED_INSTRUCTIONS, lineParsed[0]) {
            operands := make([]int, 3)

            for i := 1; i < 4; i++ {
                operands[i - 1], err = strconv.Atoi(lineParsed[i])
                if err != nil {
                    return
                }
            }

            instruction := MemInstruction{
                OpCode: lineParsed[0],
                Operand1: operands[0],
                Operand2: operands[1],
                Operand3: operands[2],
            }

            instructions = append(instructions, instruction)
        }
    }

    return
}

func (m *MemInstruction) String() string {
    return fmt.Sprintf("{\n\t\"opCode\": %s,\n\t\"operand1\": %d,\n\t\"operand2\": %d,\n\t\"operand3\": %d,\n}", m.OpCode, m.Operand1, m.Operand2, m.Operand3)
}
