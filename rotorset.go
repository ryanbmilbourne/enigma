package enigma

import "fmt"

// RotorSet Represents a set of Rotors that are part of an Enigma machine.
// Set to represent a Type 1 Enigma (3 Rotors)
type RotorSet struct {
	refRotor    Rotor
	leftRotor   Rotor
	middleRotor Rotor
	rightRotor  Rotor

	// Here for posterity.  In the real machine, this rotor converts the electrical signal
	// used by the keyboard/plugboard to the mechanical contact system used by the rotors.
	// In this application, it has no use.
	echoRotor Rotor
}

// NewRotorSet Initializes a new rotor set using the provided initial rotor/ring settings.
func NewRotorSet(settings Settings) *RotorSet {
	r := RotorSet{
		refRotor:    *NewRotor(ReflectorC, 'a', 'a'), // Reflector rotor does not have ring settings or positionality
		leftRotor:   *NewRotor(settings.RotorTypes[0], settings.RingOffsets[0], settings.RotorInits[0]),
		middleRotor: *NewRotor(settings.RotorTypes[1], settings.RingOffsets[1], settings.RotorInits[1]),
		rightRotor:  *NewRotor(settings.RotorTypes[2], settings.RingOffsets[2], settings.RotorInits[2]),
		echoRotor:   *NewRotor(Echo, 'a', 'a'),
	}

	return &r
}

// Map Passes a letter through the rotor set
func (r *RotorSet) Map(inByte byte) byte {
	fmt.Println("==========================================")

	r.turnRotors()

	// Before doing anything, turn the rotors
	fmt.Printf("Rotor states: (%v%v%v)\n",
		string(r.leftRotor.GetPosition()),
		string(r.middleRotor.GetPosition()),
		string(r.rightRotor.GetPosition()),
	)

	// First Pass
	fmt.Println("RIGHT")
	val := r.rightRotor.Enc(inByte)
	fmt.Println("MIDDLE")
	val = r.middleRotor.Enc(val)
	fmt.Println("LEFT")
	val = r.leftRotor.Enc(val)

	fmt.Println("REFLECTOR")
	val = RotorCiphers[r.refRotor.Type][val-97]

	fmt.Println("-------------------------------")

	// On the flip-side.
	fmt.Println("LEFT")
	val = r.leftRotor.Dec(val)
	fmt.Println("MIDDLE")
	val = r.middleRotor.Dec(val)
	fmt.Println("RIGHT")
	val = r.rightRotor.Dec(val)
	fmt.Println("==========================================")
	return val
}

// RotateRotors turns the rotors in the set, turning over as needed.
func (r *RotorSet) turnRotors() {
	if r.rightRotor.GetPosition() == Turnovers[r.rightRotor.Type] {
		// Right rotor has turnedover, so we need to rotate the middle rotor
		if r.middleRotor.GetPosition() == Turnovers[r.middleRotor.Type] {
			// Middle rotor also turned, so rotate the left rotor.
			r.leftRotor.Rotate()
		}
		r.middleRotor.Rotate()
	} else {
		// Account for the double stepping of the middle rotor.
		if r.middleRotor.GetPosition() == Turnovers[r.middleRotor.Type] {
			r.middleRotor.Rotate()
			r.leftRotor.Rotate()
		}
	}
	r.rightRotor.Rotate()
}
