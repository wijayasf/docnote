package notary

import "strings"

func (b *Block) Mine(difficulty int) {
  target := strings.Repeat("0", difficulty)
  for {
    b.Hash = b.calculateHash()
    if strings.HasPrefix(b.Hash, target) {
      break
    }
    b.Nonce++
  }
}
