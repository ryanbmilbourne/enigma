package enigma

import "errors"

// Machine represents an Enigma machine
type Machine struct {
	PlugBoard PlugBoard
	Rotors    RotorSet
	Settings  Settings
}

// NewMachine sets up an Enigma machine with the provided operator settings.
func NewMachine(settings Settings) *Machine {
	m := Machine{
		PlugBoard: *NewPlugBoard(settings),
		Rotors:    *NewRotorSet(settings),
		Settings:  settings,
	}

	return &m
}

// Map is the equivelent of inputting a Key on the Enigma.
// Routes the given byte through the plugboard and rotor set.
func (m *Machine) Map(inByte byte) (byte, error) {
	if err := CheckByte(inByte); err != nil {
		return 0, errors.New("Character must be a-z, lowercase")
	}

	// Pass through the Plugboard. There...
	val := m.PlugBoard.Transform(inByte)

	val = m.Rotors.Map(val)

	// ...and back again.
	val = m.PlugBoard.Transform(val)

	return val - 32, nil //Convert to upper-case to signify encrypted
}
