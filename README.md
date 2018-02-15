# enigma
A technical simulation of an Enigma Machine

A simulation of an M3 Enigma machine.  Supports Type1,2,3 wheels, and one Reflector Type (C).

Pairs with a paper written for CIS628 (Introduction to Cryptography)

## Running

Build the e_cli binary, in "./cmd/e_cli/..." :
```
$ cd cmd/e_cli
$ go build
```

Once built, you can call it like so:

```
$ e_cli -key abc -msg "Hi Professor Kyung"
```

You will receive output as each character moves through each section of the Enigma, as well as
final output.  

To decrypt a message, all you need is the encryption key, the ciphertext, and knowledge of the initial settings.
Initial Rotor positions, ring settings and plug board settings can be modified in `e_cli.go`.  The defaults are:

```go
defaultRotorTypes := [3]int{enigma.Type1, enigma.Type2, enigma.Type3}
defaultRingOffsets := [3]byte{'w', 'n', 'm'}
defaultRotorInits := [3]byte{'r', 'a', 'o'}
defaultPlugBoard := [][2]byte{}
```

### Sample output:
Here is a sample encoding "hello" using a key of ABC.
```
$ ./e_cli -key "abc" -msg "hello"

====================
Encoding Key!
====================
Plugboard: a -> a
==========================================
Rotor states: (rap)
RIGHT
Type: 3 ... In: a -> Out: e
MIDDLE
Type: 2 ... In: e -> Out: t
LEFT
Type: 1 ... In: t -> Out: d
-------------------------------
REFLECTOR
Type: 6 ... In: d -> Out: j
-------------------------------
LEFT
Type: 1 ... In: j -> Out: f
MIDDLE
Type: 2 ... In: f -> Out: r
RIGHT
Type: 3 ... In: r -> Out: t
==========================================
Plugboard: t -> t
Plugboard: b -> b
==========================================
Rotor states: (raq)
RIGHT
Type: 3 ... In: b -> Out: h
MIDDLE
Type: 2 ... In: h -> Out: c
LEFT
Type: 1 ... In: c -> Out: w
-------------------------------
REFLECTOR
Type: 6 ... In: w -> Out: n
-------------------------------
LEFT
Type: 1 ... In: n -> Out: a
MIDDLE
Type: 2 ... In: a -> Out: g
RIGHT
Type: 3 ... In: g -> Out: q
==========================================
Plugboard: q -> q
Plugboard: c -> c
==========================================
Rotor states: (rar)
RIGHT
Type: 3 ... In: c -> Out: k
MIDDLE
Type: 2 ... In: k -> Out: i
LEFT
Type: 1 ... In: i -> Out: k
-------------------------------
REFLECTOR
Type: 6 ... In: k -> Out: r
-------------------------------
LEFT
Type: 1 ... In: r -> Out: h
MIDDLE
Type: 2 ... In: h -> Out: u
RIGHT
Type: 3 ... In: u -> Out: h
==========================================
Plugboard: h -> h
Encrypted Key: TQH

====================
Encoding Message!
====================
Plugboard: h -> h
==========================================
Rotor states: (tqe)
RIGHT
Type: 3 ... In: h -> Out: w
MIDDLE
Type: 2 ... In: w -> Out: b
LEFT
Type: 1 ... In: b -> Out: f
-------------------------------
REFLECTOR
Type: 6 ... In: f -> Out: a
-------------------------------
LEFT
Type: 1 ... In: a -> Out: t
MIDDLE
Type: 2 ... In: t -> Out: j
RIGHT
Type: 3 ... In: j -> Out: i
==========================================
Plugboard: i -> i
Plugboard: e -> e
==========================================
Rotor states: (tqf)
RIGHT
Type: 3 ... In: e -> Out: z
MIDDLE
Type: 2 ... In: z -> Out: a
LEFT
Type: 1 ... In: a -> Out: u
-------------------------------
REFLECTOR
Type: 6 ... In: u -> Out: s
-------------------------------
LEFT
Type: 1 ... In: s -> Out: w
MIDDLE
Type: 2 ... In: w -> Out: p
RIGHT
Type: 3 ... In: p -> Out: x
==========================================
Plugboard: x -> x
Plugboard: l -> l
==========================================
Rotor states: (tqg)
RIGHT
Type: 3 ... In: l -> Out: r
MIDDLE
Type: 2 ... In: r -> Out: m
LEFT
Type: 1 ... In: m -> Out: c
-------------------------------
REFLECTOR
Type: 6 ... In: c -> Out: p
-------------------------------
LEFT
Type: 1 ... In: p -> Out: f
MIDDLE
Type: 2 ... In: f -> Out: c
RIGHT
Type: 3 ... In: c -> Out: x
==========================================
Plugboard: x -> x
Plugboard: l -> l
==========================================
Rotor states: (tqh)
RIGHT
Type: 3 ... In: l -> Out: h
MIDDLE
Type: 2 ... In: h -> Out: i
LEFT
Type: 1 ... In: i -> Out: j
-------------------------------
REFLECTOR
Type: 6 ... In: j -> Out: d
-------------------------------
LEFT
Type: 1 ... In: d -> Out: x
MIDDLE
Type: 2 ... In: x -> Out: x
RIGHT
Type: 3 ... In: x -> Out: c
==========================================
Plugboard: c -> c
Plugboard: o -> o
==========================================
Rotor states: (tqi)
RIGHT
Type: 3 ... In: o -> Out: b
MIDDLE
Type: 2 ... In: b -> Out: p
LEFT
Type: 1 ... In: p -> Out: r
-------------------------------
REFLECTOR
Type: 6 ... In: r -> Out: k
-------------------------------
LEFT
Type: 1 ... In: k -> Out: s
MIDDLE
Type: 2 ... In: s -> Out: u
RIGHT
Type: 3 ... In: u -> Out: c
==========================================
Plugboard: c -> c

hello ->
IXXCC
```


