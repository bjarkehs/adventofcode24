package day15

func findMove(moveAsString string) int {
	switch moveAsString {
	case "^":
		return 0
	case ">":
		return 1
	case "v":
		return 2
	case "<":
		return 3
	default:
		panic("Invalid move")
	}
}
