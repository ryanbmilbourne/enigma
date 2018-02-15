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
	Echo:       Alphabet,
}

// Rotor represents an Enigma rotor assembly
type Rotor struct {
	Type             int
	substitutions    []byte
	AlphaRingOffset  int
	AlphaRingSetting byte
	position         byte
}

// NewRotor initializes a new rotor using the given type (I,II,III, Echo, or Reflector)
func NewRotor(rotorType int, alphaRingSetting, rotorInitPosition byte) *Rotor {
	rotor := Rotor{}

	rotor.Type = rotorType

	rotor.substitutions = RotorCiphers[rotorType]

	rotor.AlphaRingOffset = int(alphaRingSetting) - 97
	rotor.AlphaRingSetting = alphaRingSetting

	rotor.position = rotorInitPosition

	return &rotor
}

// Accounts for the static ring setting and current position to provide a lookup index
// for the substitution.
func (r *Rotor) doRotorPrework(charNum int) int {
	// Remove the ASCII offest for the ring and rotor math
	lutIdx := charNum

	// Account for the static ring setting
	lutIdx = lutIdx - r.AlphaRingOffset
	if lutIdx < 0 {
		lutIdx = lutIdx + 26
	}
	// Now, account for the position of the rotor
	pos := int(r.position) - 97
	lutIdx = (lutIdx + pos) % 26

	// Note that we leave the ASCII off.  This is because the result of this is used for a lookup
	return lutIdx
}

// Re-Accounts for the static ring setting and current position to provide the output byte
func (r *Rotor) doRotorPostwork(charNum int) byte {
	pos := int(r.position) - 97

	// Re-account for the rotor position
	charNum = charNum - pos
	if charNum < 0 {
		charNum = charNum + 26
	}

	// Re-account for the ring setting
	charNum = (charNum + r.AlphaRingOffset) % 26

	// Re-add the ASCII
	charNum = charNum + 97
	return byte(charNum)
}

// Enc maps a character through the rotor that is directionally headed **towards** the reflector.
func (r *Rotor) Enc(inByte byte) byte {
	charNum := int(inByte) - 97

	fmt.Printf("Type: %v ... ", r.Type)
	fmt.Printf("In: %v -> ", string(inByte))

	// reflectors don't have ring settings or rotor positions
	if r.Type == ReflectorC {
		outByte := RotorCiphers[ReflectorC][charNum]
		fmt.Printf("Out: %v\n", string(outByte))
		return outByte
	}

	// Account for the ring setting and rotor position
	lutIdx := r.doRotorPrework(charNum)

	// Do the substitution
	outByte := RotorCiphers[r.Type][lutIdx]
	outCharNum := int(outByte) - 97

	// Un-account for the ring setting and rotor position
	outByte = r.doRotorPostwork(outCharNum)
	fmt.Printf("Out: %v\n", string(outByte))

	return outByte
}

// Dec maps a character through the rotor that is directionally headed **back from** the reflector.
func (r *Rotor) Dec(inByte byte) byte {
	charNum := int(inByte) - 97
	fmt.Printf("Type: %v ... ", r.Type)
	fmt.Printf("In: %v -> ", string(inByte))

	// Account for the ring setting and rotor position
	lutIdx := r.doRotorPrework(charNum)

	// Find the index of that byte in the cipher array
	foundIdx := 0
	for i, val := range RotorCiphers[r.Type] {
		if val == byte(lutIdx+97) {
			foundIdx = i
			break
		}
	}

	// Un-account for the ring setting and rotor position
	outByte := r.doRotorPostwork(foundIdx)
	fmt.Printf("Out: %v\n", string(outByte))

	return outByte
}

// Rotate advances the position of the Rotor by one step.
// Return true if the rotor has reached its turnover point.
// Each rotor type has a different notch turnover point.
func (r *Rotor) Rotate() {
	// ASCII so fun
	r.position = ((r.position + 1 - 97) % 26) + 97
}

// GetPosition returns the current position of the rotor wheel.
func (r *Rotor) GetPosition() byte {
	return r.position
}
