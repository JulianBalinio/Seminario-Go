package main

import (
    "Ej2/blockchain"
    "Ej2/wallet"
    "fmt"
)

func main() {
    // Crear una nueva cadena de bloques
    chain := blockchain.NewBlockchain()

    // Crear billeteras para dos usuarios
    julianWallet := wallet.CreateWallet("Julian", "Baliño")
    dexterWallet := wallet.CreateWallet("Dexter", "Morgan")

	julianWallet.AddFunds(100.00)
	fmt.Printf("Fondos Julian: $%.2f\n", julianWallet.GetFunds())
	fmt.Printf("Fondos Dexter: $%.2f\n", dexterWallet.GetFunds())

    //amount := 10.0
    err := wallet.SendTransaction(julianWallet, dexterWallet, 33.4, chain)
    if err != nil {
        fmt.Println("Error al enviar la transacción:", err)
        return
    }
    fmt.Println("Transacción enviada con éxito.")
	fmt.Printf("Fondos Julian: $%.2f\n", julianWallet.GetFunds())
    fmt.Printf("Fondos Dexter: $%.2f\n", dexterWallet.GetFunds())
	fmt.Printf("ID Julian: %d\n", julianWallet.GetID())
    fmt.Printf("ID Dexter: %d\n", dexterWallet.GetID())

	err = wallet.SendTransaction(julianWallet, dexterWallet, 5, chain)
    if err != nil {
        fmt.Println("Error al enviar la transacción:", err)
        return
    }
    fmt.Println("Transacción enviada con éxito.")

    //Balance de transacciones
    julianBalance := wallet.GetBalance(julianWallet.GetID(), chain)
    dexterBalance := wallet.GetBalance(dexterWallet.GetID(), chain)
    fmt.Printf("Balance de Julian: %.2f\n", julianBalance)
    fmt.Printf("Balance de Dexter: %.2f\n", dexterBalance)

    // Validar la cadena de bloques
    err = chain.IsChainValid()
    if err != nil {
        fmt.Println("La cadena de bloques no es válida:", err)
        return
    }
    fmt.Println("La cadena de bloques es válida.")
	fmt.Printf("Fondos Julian: $%.2f\n", julianWallet.GetFunds())
	fmt.Printf("Fondos Dexter: $%.2f\n", dexterWallet.GetFunds())
}
