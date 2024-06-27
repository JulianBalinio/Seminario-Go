// wallet.go
package wallet

import (
    "errors"
    "Ej2/blockchain"
    "Ej2/transaction"
    "Ej2/block"
)

type Wallet struct {
    id        int
    firstName string
    lastName  string 
    balance   float64
}

var (
    ErrInvalidAmount    = errors.New("invalid amount in operation.")
    ErrNotEnoughBalance = errors.New("insufficient balance.")
    lastWalletID       = 0
)

func CreateWallet(firstName, lastName string) *Wallet {
    lastWalletID++
    return &Wallet{
        id:        lastWalletID,
        firstName: firstName,
        lastName:  lastName,
        balance:   0.0, //saldo inicial
    }
}

func GetBalance(userID int, chain *blockchain.Blockchain) float64 {
    balance := 0.0
    current := chain.Head
    for current != nil {
        for _, transaction := range current.Block.Transactions {
            if transaction.GetSenderID() == userID {
                balance -= transaction.GetAmount()
            }
            if transaction.GetReceiverID() == userID {
                balance += transaction.GetAmount()
            }
        }
        current = current.Next
    }
    return balance
}

func SendTransaction(senderWallet, receiverWallet *Wallet, amount float64, chain *blockchain.Blockchain) error {
    if senderWallet.balance < amount {
        return ErrNotEnoughBalance
    }

    newTransaction := transaction.NewTransaction(amount, senderWallet.GetID(), receiverWallet.GetID())

    //Si la blockchain esta vacia creo el primer bloque y le agrego la transaccion al slice del bloque
    if chain.Head == nil {
        firstBlock := block.NewBlock([]transaction.Transaction{*newTransaction}, nil)
        chain.Head = &blockchain.BlockNode{Block: firstBlock}
		//Actualizo el balance correspondiente
        senderWallet.balance -= amount //No se xq carajo no funciona con AddFunds (A solucionar)
    	receiverWallet.balance += amount 
        return nil
    }

    //Agrego la transaccion al slice del bloque correspondiente
    currentBlock := chain.Head
    for currentBlock.Next != nil {
        currentBlock = currentBlock.Next
    }
    currentBlock.Block.Transactions = append(currentBlock.Block.Transactions, *newTransaction)

    //Actualizo el balance correspondiente
    senderWallet.balance -= amount
    receiverWallet.balance += amount

    return nil
}

func (w *Wallet) AddFunds(amount float64) error {
    if amount < 0 {
        return ErrInvalidAmount
    }
    w.balance += amount
    return nil
}

// Getters
func (w *Wallet) GetID() int {
    return w.id
}

func (w *Wallet) GetFirstName() string {
    return w.firstName
}

func (w *Wallet) GetLastName() string {
    return w.lastName
}

func (w *Wallet) GetFunds() float64 {
    return w.balance
}
