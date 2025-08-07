package notary

var defaultBC = NewBlockchain(3)

func Notarize(data string) Block {
	return defaultBC.AddBlock(data)
}

func GetChain() []Block {
	return defaultBC.GetChain()
}
