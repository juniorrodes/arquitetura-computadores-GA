package pipeline

import (
	"fmt"
	"strings"
)

type State struct {
    Pc int
    InstructionMemory []MemInstruction
    MainMemory        [100]*int
    Registers         [32]int
    queue             *Queue
    labels            Labels
}

func NewState() *State {
    return &State{
        Pc: 0,
        queue: &Queue{},
    }
}

func (s *State) String() string {
    var sb strings.Builder

    sb.WriteString("atual state: {\r\n\"pc\": ")
    sb.WriteString(fmt.Sprintf("%d,\r\n", s.Pc))
    sb.WriteString("Registers: {[\r\n")
    for i, r := range s.Registers {
        sb.WriteString(fmt.Sprintf("\"%d\": %d,", i, r))
    }
    sb.WriteString("]},")

    return sb.String()
}
