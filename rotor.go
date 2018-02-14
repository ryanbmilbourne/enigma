package enigma

import "fmt"

const (
	Type1      = 1
	Type2      = 2
	Type3      = 3
	ReflectorB = 5 // The reflector used in the Enigma I B
	ReflectorC = 6 // Used in Enigma 1 C
	Echo       = 9
)

// Each rotor is its own substitution cipher.
var RotorCiphers = map[int][]byte{
	Type1:      []byte("ekmflgdqvzntowyhxuspaibrcj"),
	Type2:      []byte("ajdksiruxblhwtmcqgznpyfvoe"),
	Type3:      []byte("bdfhjlcprtxvznyeiwgakmusqo"),
	ReflectorC: []byte("fvpjiaoyedrzxwgctkuqsbnmhl"),
	Echo:       []byte(Alphabet),
}

type Rotor struct {
	Type          int
	substitutions []byte
	started       bool

	// Channels for the "input" side
	InChan  chan byte
	OutChan chan byte
	// Channels for the "output" side
	RefInChan  chan byte
	RefOutChan chan byte
	// Goroutine exit chan
	exitChan  chan bool
	rExitChan chan bool
}

func NewRotor(rotorType int) *Rotor {
	rotor := Rotor{}

	rotor.Type = rotorType

	//rotor.InChan = make(chan byte)
	//rotor.OutChan = make(chan byte)
	//rotor.RefInChan = make(chan byte)
	//rotor.RefOutChan = make(chan byte)

	//rotor.exitChan = make(chan bool)
	//rotor.rExitChan = make(chan bool)

	rotor.substitutions = RotorCiphers[rotorType]

	return &rotor
}

func (r *Rotor) Start() {
	r.started = true
	// Start the goroutines to parse input to the rotor
	go func() {
		for {
			select {
			case inByte := <-r.InChan:
				fmt.Printf("Type %v Rx %v :: ", r.Type, inByte)
				if r.Type == ReflectorC {
					r.RefOutChan <- r.Enc(inByte)
				} else {
					// Encrypt and output
					r.OutChan <- r.Enc(inByte)
				}
			case <-r.exitChan:
				return
			}
		}
	}()

	if r.Type != ReflectorC {
		go func() {
			for {
				select {
				case inByte := <-r.RefInChan:
					// Encrypt and output
					r.RefOutChan <- r.Dec(inByte)
				case <-r.rExitChan:
					return
				}
			}
		}()
	}
}

func (r *Rotor) Exit() {
	r.exitChan <- true
	if r.Type != ReflectorC {
		r.rExitChan <- true
	}
}

func (r *Rotor) Enc(inByte byte) byte {
	return inByte
}

func (r *Rotor) Dec(inByte byte) byte {
	return inByte
}
