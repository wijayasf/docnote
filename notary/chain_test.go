package notary

import "testing"

func TestChain(t *testing.T) {
	bc := NewBlockchain(2)
	blk := bc.AddBlock("abc")
	if blk.Index != 1 || blk.Data != "abc" {
		t.Fatal("Block not added correctly")
	}
	if len(bc.GetChain()) != 2 {
		t.Fatal("Chain length should be 2")
	}
}
