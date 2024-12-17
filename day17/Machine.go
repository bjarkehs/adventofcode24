package day17

import (
	"fmt"
	"strconv"
)

type Machine struct {
	RegisterA int
	RegisterB int
	RegisterC int

	InstructionPointer int

	Program []int
	Output  []int
}

func (machine *Machine) Run() {
	for machine.InstructionPointer < len(machine.Program) {
		machine.RunInstruction()
	}
}

func (machine *Machine) PrintOutput() {
	outputString := ""
	for index, output := range machine.Output {
		if index > 0 {
			outputString += ","
		}
		outputString += strconv.Itoa(output)
	}
	fmt.Println(outputString)
}

func (machine *Machine) RunInstruction() {
	instruction := machine.Program[machine.InstructionPointer]
	operand := machine.Program[machine.InstructionPointer+1]
	jumpToNextInstruction := true
	switch instruction {
	case 0:
		machine.RegisterA = machine.RegisterA >> machine.ValueOfComboOperand(operand)
	case 1:
		machine.RegisterB = machine.RegisterB ^ operand
	case 2:
		machine.RegisterB = machine.ValueOfComboOperand(operand) % 8
	case 3:
		if machine.RegisterA != 0 {
			machine.InstructionPointer = operand
			jumpToNextInstruction = false
		}
	case 4:
		machine.RegisterB = machine.RegisterB ^ machine.RegisterC
	case 5:
		machine.Output = append(machine.Output, machine.ValueOfComboOperand(operand)%8)
	case 6:
		machine.RegisterB = machine.RegisterA >> machine.ValueOfComboOperand(operand)
	case 7:
		machine.RegisterC = machine.RegisterA >> machine.ValueOfComboOperand(operand)
	}

	if jumpToNextInstruction {
		machine.InstructionPointer += 2
	}
}

func (machine *Machine) ValueOfComboOperand(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return machine.RegisterA
	case 5:
		return machine.RegisterB
	case 6:
		return machine.RegisterC
	case 7:
		panic("Invalid program")
	default:
		panic("Invalid program")
	}
}
