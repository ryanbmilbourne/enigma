package enigma

import (
	"fmt"
	"strings"
)

// Alphabet is just...an alphabet
var Alphabet = []byte("abcdefghijklmnopqrstuvwxyz")

// Settings Represents the External inputs that the operator would provide to Enigma
type Settings struct {
	RotorTypes  [3]int    // Which rotor types to use in each slot
	RingOffsets [3]byte   // Alpha setting for the inner rotor ring.  Static.
	Plugboard   [][2]byte // Cords to use in the PlugBoard. Static
	RotorInits  [3]byte   // Initial Alpha settings for the rotor wheel. Dynamic.
}

// Turnovers notes the char at which each rotor type turns over
var Turnovers = map[int]byte{
	Type1: 'q',
	Type2: 'e',
	Type3: 'v',
}

// CheckByte Checks if a byte is valid input (lowercase alpha)
func CheckByte(inByte byte) error {
	if inByte < 97 || inByte > 122 {
		return fmt.Errorf("Invalid byte character: `%v`.  Chars must be [a-z]")
	}
	return nil
}

// Smash converts an input string to a byte array that the Enigma machine can process.
// It weeds out any invalid (non-alpha) characters
func Smash(text string) []byte {
	text = strings.ToLower(text)

	retText := make([]byte, 0, len(text))

	for _, char := range []byte(text) {
		if err := CheckByte(char); err != nil {
			continue
		}
		retText = append(retText, char)
	}
	return retText
}
