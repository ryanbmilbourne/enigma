package main

import "fmt"
import "github.com/ryanbmilbourne/enigma"

func main() {
	settings := enigma.Settings{
		RotorTypes:  [3]int{enigma.Type1, enigma.Type2, enigma.Type3},
		RingOffsets: [3]byte{'w', 'n', 'm'},
		RotorInits:  [3]byte{'r', 'a', 'o'},
	}

	theMachine := enigma.NewMachine(settings)

	inputText := enigma.Smash("Hello")
	//inputText = enigma.Smash("BJUFI")

	outputText := make([]byte, 0, len(inputText))

	for _, b := range inputText {
		out, err := theMachine.Map(b)
		if err != nil {
			out = '.'
		}
		outputText = append(outputText, out)
	}
	//for i, out := range outputText {
	//	fmt.Printf(
	//		"\n%s -> %s",
	//		string(inputText[i]),
	//		string(out),
	//	)
	//}

	fmt.Printf("%v -> %v\n", string(inputText), string(outputText))
}
