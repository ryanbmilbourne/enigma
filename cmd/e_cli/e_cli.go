package main

import (
	"flag"
	"fmt"

	"github.com/ryanbmilbourne/enigma"
)

func main() {

	var (
		key     string
		message string
	)

	defaultRotorTypes := [3]int{enigma.Type1, enigma.Type2, enigma.Type3}
	defaultRingOffsets := [3]byte{'w', 'n', 'm'}
	defaultRotorInits := [3]byte{'r', 'a', 'o'}
	//defaultPlugBoard := [][2]byte{}

	const (
		defaultKey = "abc"
		defaultMsg = ""
	)

	flag.StringVar(&key, "key", defaultKey, "key to use for encryption")

	flag.StringVar(&message, "msg", defaultMsg, "message to encrypt")

	flag.Parse()

	if message == "" {
		fmt.Println("non-null message required")
		return
	}

	settings := enigma.Settings{
		RotorTypes:  defaultRotorTypes,
		RingOffsets: defaultRingOffsets,
		RotorInits:  defaultRotorInits,
	}

	initMachine := enigma.NewMachine(settings)

	keyText := enigma.Smash(key)
	encKey := make([]byte, 0)
	for _, b := range keyText {
		out, err := initMachine.Map(b)
		if err != nil {
			out = '.'
		}
		encKey = append(encKey, out)
	}

	fmt.Printf("ENC KEY: %s\n", encKey)
	inputKey := enigma.Smash(string(encKey))
	var keyArr [3]byte
	for i := 0; i < 2; i++ {
		keyArr[i] = inputKey[i]
	}

	settings = enigma.Settings{
		RotorTypes:  defaultRotorTypes,
		RingOffsets: defaultRingOffsets,
		RotorInits:  keyArr,
	}

	encMachine := enigma.NewMachine(settings)

	inputText := enigma.Smash(message)
	//inputText = enigma.Smash("KYRJR")

	outputText := make([]byte, 0, len(message))
	for _, b := range inputText {
		out, err := encMachine.Map(b)
		if err != nil {
			out = '.'
		}
		outputText = append(outputText, out)
	}

	fmt.Printf("%v -> %v\n", string(inputText), string(outputText))
}
