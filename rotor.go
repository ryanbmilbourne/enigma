package enigma

import "fmt"

const (
	// Type1 rotor wheel
	Type1 = 1
	// Type2 rotor wheel
	Type2 = 2
	// Type3 rotor wheel
	Type3 = 3
	// ReflectorB for Model IB
	ReflectorB = 5
	// ReflectorC for Model IC
	ReflectorC = 6
	// Echo is just a Pass-through
	Echo = 9
)

// RotorCiphers contains the substitution cipher for each rotor type
var RotorCiphers = map[int][]byte{
	Type1:      []byte("ekmflgdqvzntowyhxuspaibrcj"),
	Type2:      []byte("ajdksiruxblhwtmcqgznpyfvoe"),
	Type3:      []byte("bdfhjlcprtxvznyeiwgakmusqo"),
	ReflectorC: []byte("fvpjiaoyedrzxwgctkuqsbnmhl"),
	Echo:       []byte(Alphabet),
}

// Rotor represents an Enigma rotor assembly
type Rotor struct {
	Type            int
	substitutions   []byte
	AlphaRingOffset int
	Position        byte
}

// NewRotor initializes a new rotor using the given type (I,II,III, Echo, or Reflector)
func NewRotor(rotorType int, alphaRingSetting, rotorInitPosition byte) *Rotor {
	rotor := Rotor{}

	rotor.Type = rotorType

	rotor.substitutions = RotorCiphers[rotorType]

	rotor.AlphaRingOffset = int(alphaRingSetting - 'a')

	rotor.Position = rotorInitPosition

	return &rotor
}

// Accounts for the static ring setting and current position to provide a lookup index
// for the substitution.
func (r *Rotor) doRotorPrework(inByte byte) int {
	// Remove the ASCII offest for the ring and rotor math
	lutIdx := int(inByte - 97)

	// Account for the static ring setting
	lutIdx = lutIdx - r.AlphaRingOffset
	if lutIdx < 0 {
		lutIdx = lutIdx + 26
	}
	// Now, account for the position of the rotor
	lutIdx = (lutIdx + (int(r.Position) - 97)) % 26

	// Note that we leave the ASCII off.  This is because the result of this is used for a lookup

	return lutIdx
}

// Re-Accounts for the static ring setting and current position to provide the output byte
func (r *Rotor) doRotorPostwork(outByte byte) byte {
	// Strip off the ASCII for the math.
	outByte = outByte - 97

	// Re-account for the rotor position
	outByte = outByte - (r.Position - 97)
	if outByte < 0 {
		outByte = outByte + 26
	}

	// Re-account for the ring setting
	outByte = (outByte + byte(r.AlphaRingOffset)) % 26

	// Re-add the ASCII
	outByte = outByte + 97
	return outByte
}

// Enc maps a character through the rotor that is directionally headed **towards** the reflector.
func (r *Rotor) Enc(inByte byte) byte {
	fmt.Printf("Type: %v ... ", r.Type)
	fmt.Printf("In: %v (%v) -> ", inByte, string(inByte))

	// reflectors don't have ring settings or rotor positions
	if r.Type == ReflectorC {
		outByte := RotorCiphers[ReflectorC][inByte-97]
		fmt.Printf("Out: %v (%v)\n", outByte, string(outByte))
		return outByte
	}

	// Account for the ring setting and rotor position
	lutIdx := r.doRotorPrework(inByte)
	fmt.Printf("LutIdx: %v -> ", lutIdx)

	// Do the substitution
	outByte := RotorCiphers[r.Type][lutIdx]
	fmt.Printf("OutByte: %v (%v) -> ", outByte, string(outByte))

	// Un-account for the ring setting and rotor position
	outByte = r.doRotorPostwork(outByte)
	fmt.Printf("Out: %v (%v)\n", outByte, string(outByte))

	return outByte
}

// Dec maps a character through the rotor that is directionally headed **back from** the reflector.
func (r *Rotor) Dec(inByte byte) byte {
	fmt.Printf("Type: %v ... ", r.Type)
	fmt.Printf("In: %v (%v) -> ", inByte, string(inByte))

	// reflectors don't have ring settings or rotor positions
	if r.Type == ReflectorC {
		return 0
	}

	// Account for the ring setting and rotor position
	lutIdx := r.doRotorPrework(inByte)
	fmt.Printf("LutIdx: %v -> ", lutIdx)

	// Find the index of that byte in the cipher array
	foundIdx := 0
	for i, val := range RotorCiphers[r.Type] {
		if val == byte(lutIdx+97) {
			foundIdx = i
			break
		}
	}

	outByte := byte(foundIdx + 97)

	fmt.Printf("OutByte: %v (%v) -> ", outByte, string(outByte))

	// Un-account for the ring setting and rotor position
	outByte = r.doRotorPostwork(outByte)
	fmt.Printf("Out: %v (%v)\n", outByte, string(outByte))

	return outByte
}
