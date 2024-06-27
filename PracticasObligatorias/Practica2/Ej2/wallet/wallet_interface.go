package wallet

import "Ej2/blockchain"

type WalletInterface interface {
    GetBalance(userID string, chain *blockchain.Blockchain) float64
    SendTransaction(sender, receiver string, amount float64, chain *blockchain.Blockchain) error
}
