package enigma

import (
	"errors"
	"fmt"
)

type PlugBoard struct {
	transforms map[byte]byte
}

// Plug'n'chug
func NewPlugBoard(settings Settings) *PlugBoard {
	plugb := PlugBoard{}
	for _, buple := range settings.Plugboard {
		err := plugb.AddPlug(buple)
		if err != nil {
			fmt.Println(err)
		}
	}
	return &plugb
}

// Use the plugboard to transform a letter to the other side of its "plug"
func (p *PlugBoard) Transform(inByte byte) byte {
	// Are there any plugs...plugged in?
	if p.transforms == nil {
		return inByte
	}
	// If there's a transformation mapped, use it.  Else return the original value
	if val, ok := p.transforms[inByte]; ok {
		return val
	}
	return inByte
}

// Add a new wire to the plugboard
func (p *PlugBoard) AddPlug(buple [2]byte) error {
	// A plug can't go into itself
	if buple[0] == buple[1] {
		return errors.New("A plug must go to two characters!")
	}

	// Are there any plugs...plugged in?
	if p.transforms == nil {
		p.transforms = make(map[byte]byte)
	}

	// Because we add a transform for both directions, we only need to check the keys
	for _, b := range buple {
		if _, ok := p.transforms[b]; ok {
			return fmt.Errorf("%v is already in use on the Plug Board!", b)
		}
	}

	// Add the character mapping
	p.transforms[buple[0]] = buple[1]
	p.transforms[buple[1]] = buple[0]

	return nil
}

// Rip the plugs out.  Roughly.
func (p *PlugBoard) Reset() {
	p.transforms = make(map[byte]byte)
}
