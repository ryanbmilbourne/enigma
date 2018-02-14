package enigma

type RotorSet struct {
	refRotor    Rotor
	leftRotor   Rotor
	middleRotor Rotor
	rightRotor  Rotor
	echoRotor   Rotor

	InChan   chan byte
	OutChan  chan byte
	exitChan chan bool
}

func NewRotorSet(settings Settings) *RotorSet {
	r := RotorSet{
		refRotor:    *NewRotor(ReflectorB),
		leftRotor:   *NewRotor(settings.Rotors[0]),
		middleRotor: *NewRotor(settings.Rotors[1]),
		rightRotor:  *NewRotor(settings.Rotors[2]),
		echoRotor:   *NewRotor(Echo),

		InChan:   make(chan byte, 32),
		OutChan:  make(chan byte),
		exitChan: make(chan bool),
	}

	//echoRightBridge := make(chan byte)
	//echoRightRefBridge := make(chan byte)

	//rightMiddleBridge := make(chan byte)
	//rightMiddleRefBridge := make(chan byte)

	//middleLeftBridge := make(chan byte)
	//middleLeftRefBridge := make(chan byte)

	//leftReflectBridge := make(chan byte)
	//leftReflectRefBridge := make(chan byte)

	//echo.InChan = make(chan byte)
	//echo.RefOutChan = make(chan byte)

	// Setup the channels
	//r.echoRotor.OutChan = echoRightBridge
	//r.echoRotor.RefInChan = echoRightRefBridge

	//r.rightRotor.InChan = echoRightBridge
	//r.rightRotor.RefOutChan = echoRightRefBridge
	//r.rightRotor.OutChan = rightMiddleBridge
	//r.rightRotor.RefInChan = rightMiddleRefBridge

	//r.middleRotor.InChan = rightMiddleBridge
	//r.middleRotor.RefOutChan = rightMiddleRefBridge
	//r.middleRotor.OutChan = middleLeftBridge
	//r.middleRotor.RefInChan = middleLeftRefBridge

	//r.leftRotor.InChan = middleLeftBridge
	//r.leftRotor.RefOutChan = middleLeftRefBridge
	//r.leftRotor.OutChan = leftReflectBridge
	//r.leftRotor.RefInChan = leftReflectRefBridge

	//echoWriteChan := r.echoRotor.InChan
	//echoReadChan := r.echoRotor.RefOutChan

	//r.refRotor.Start()
	//r.rightRotor.Start()
	//r.middleRotor.Start()
	//r.leftRotor.Start()
	//r.echoRotor.Start()

	//go func() {
	//	for {
	//		select {
	//		case inByte := <-r.InChan:
	//			// First Pass
	//			val := r.rightRotor.Enc(inByte)
	//			val = r.middleRotor.Enc(val)
	//			val = r.leftRotor.Enc(val)
	//			val = r.refRotor.Enc(val)

	//			// On the flip-side.
	//			val = r.leftRotor.Dec(val)
	//			val = r.middleRotor.Dec(val)
	//			val = r.rightRotor.Dec(val)

	//			r.OutChan <- val
	//		case <-r.exitChan:
	//			r.refRotor.Exit()
	//			r.leftRotor.Exit()
	//			r.middleRotor.Exit()
	//			r.rightRotor.Exit()
	//			r.echoRotor.Exit()

	//			return
	//		}
	//	}
	//}()

	return &r
}

func (r *RotorSet) Exit() {
	r.exitChan <- true
}

func (r *RotorSet) Map(inByte byte) byte {
	// First Pass
	val := r.rightRotor.Enc(inByte)
	val = r.middleRotor.Enc(val)
	val = r.leftRotor.Enc(val)
	val = r.refRotor.Enc(val)

	// On the flip-side.
	val = r.leftRotor.Dec(val)
	val = r.middleRotor.Dec(val)
	val = r.rightRotor.Dec(val)
	return val
}
