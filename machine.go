package enigma

import "errors"

type Machine struct {
	PlugBoard PlugBoard
	Rotors    RotorSet
	Settings  Settings
}

func NewMachine(settings Settings) *Machine {
	m := Machine{
		PlugBoard: *NewPlugBoard(settings),
		Rotors:    *NewRotorSet(settings),
		Settings:  settings,
	}

	return &m
}

func (m *Machine) Map(inByte byte) (byte, error) {
	if inByte < 97 || inByte > 122 {
		return 0, errors.New("Character must be a-z, lowercase")
	}

	val := m.PlugBoard.Transform(inByte)

	val = m.Rotors.Map(val)

	return val - 32, nil //Convert to upper-case to signify encrypted
}
