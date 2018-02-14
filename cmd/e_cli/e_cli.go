package main

import "fmt"
import "github.com/ryanbmilbourne/enigma"

func main() {
	settings := enigma.Settings{
		Rotors: [3]int{enigma.Type1, enigma.Type2, enigma.Type3},
	}
	theMachine := enigma.NewMachine(settings)

	for _, b := range []byte("hello") {
		out, err := theMachine.Map(b)
		if err != nil {
			out = '.'
		}
		fmt.Printf(
			"%s -> %s\n",
			string(b),
			string(out),
		)
	}
}
