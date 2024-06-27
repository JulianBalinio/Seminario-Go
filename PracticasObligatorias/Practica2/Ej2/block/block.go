package block

import (
    "Ej2/transaction"
    "crypto/sha256"
    "time"
)

type Block struct {
    Hash          []byte
    PreviousHash  []byte
    Transactions  []transaction.Transaction
    Timestamp     time.Time
}

//Faltan agregar las transacciones para validar el Hash
func (b *Block) CalculateHash() []byte {
    data := append([]byte(b.PreviousHash), []byte(b.Timestamp.String())...)
    h := sha256.New()
    h.Write(data)
    return h.Sum(nil)
  }

func compareSlices(a, b []byte) bool {
  if len(a) != len(b) {
    return false
  }
  for i := range a {
    if a[i] != b[i] {
      return false
    }
  }
  return true
}
  
func (b *Block) IsValid(previousBlock *Block) bool {
    calculatedHash := b.CalculateHash()
    return compareSlices(calculatedHash, b.Hash) && compareSlices([]byte(b.PreviousHash), previousBlock.Hash)
}

func NewBlock(transactions []transaction.Transaction, previousHash []byte) *Block {
    return &Block{
        Hash:         nil,
        PreviousHash: previousHash,
        Transactions: transactions,
        Timestamp:    time.Now(),
    }
}
