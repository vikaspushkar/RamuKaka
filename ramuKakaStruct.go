package ramukaka

//NeuNet is the neural network node struct
type NeuNet struct {
	value            [36]byte
	endhere          bool
	qNa              map[string]*NeuNet
	frequentreferral [32]uint32
	next             [36]*NeuNet // 26 alphabets and 0-9 number
}
