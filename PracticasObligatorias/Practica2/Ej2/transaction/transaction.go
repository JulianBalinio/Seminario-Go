package transaction

import "time"

type Transaction struct {
    amount       float64
    senderID     int
    receiverID   int 
    timestamp    time.Time //marca de tiempo de la transacción
}

//Constructor
func NewTransaction(amount float64, senderID, receiverID int) *Transaction {
    return &Transaction{
        amount:     amount,
        senderID:   senderID,
        receiverID: receiverID,
        timestamp:  time.Now(),
    }
}

//Getters
func (t *Transaction) GetAmount() float64 {
    return t.amount
}

func (t *Transaction) GetSenderID() int {
    return t.senderID
}

func (t *Transaction) GetReceiverID() int {
    return t.receiverID
}

func (t *Transaction) GetTimestamp() time.Time {
    return t.timestamp
}
