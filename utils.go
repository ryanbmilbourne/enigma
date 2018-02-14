package enigma

const Alphabet string = "abcdefghipqrstuvwxyz"

type Settings struct {
	Rotors    [3]int
	Rings     [3]byte
	Plugboard [][2]byte
}
