package notary

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Index     int
	Timestamp int64
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

func (b *Block) calculateHash() string {
	record := fmt.Sprintf("%d%d%s%s%d",
		b.Index, b.Timestamp, b.Data, b.PrevHash, b.Nonce)
	sum := sha256.Sum256([]byte(record))
	return hex.EncodeToString(sum[:])
}
