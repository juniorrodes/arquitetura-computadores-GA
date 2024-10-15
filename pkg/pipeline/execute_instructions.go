package pipeline

type ExecInstructions struct {
    MemInstruction
    Temp1 int
    Temp2 int
    Temp3 int
    Invalid bool
}

//func (i *ExecInstructions) LoadWord(state *State) {
//    offset := i.Operand1
//    destiny := i.Operand2
//    address := i.Operand3
//
//    state.Registers[destiny] = *state.MainMemory[address + offset]
//}
//
//func (i *ExecInstructions) BranchEqual(state *State) {
//    i.Temp1 = state.Registers[i.MemInstruction.Operand1]
//    i.Temp2 = state.Registers[i.MemInstruction.Operand2]
//    i.Temp3 = i.MemInstruction.Operand3
//}
