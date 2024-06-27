package blockchain

import (
    "Ej2/block"
    "errors"
)

type BlockNode struct {
    Block *block.Block
    Next  *BlockNode
}

type Blockchain struct {
    Head *BlockNode // Puntero al primer nodo en la cadena
}

var ErrInvalidChain = errors.New("invalid chain")

//crea una nueva cadena de bloques vacía
func NewBlockchain() *Blockchain {
    return &Blockchain{} // Devuelve una nueva instancia de Blockchain vacia
}

func (chain *Blockchain) InsertBlock(newBlock *block.Block) {
    newNode := &BlockNode{Block: newBlock}
    if chain.Head == nil {
        chain.Head = newNode
        return
    }

    current := chain.Head
    for current.Next != nil {
        current = current.Next
    }
    current.Next = newNode
}

func (chain *Blockchain) IsChainValid() error {
    current := chain.Head
    for current != nil && current.Next != nil {
        if !current.Block.IsValid(current.Next.Block) {
            return ErrInvalidChain
        }
        current = current.Next
    }
    return nil
}
