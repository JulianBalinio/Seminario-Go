package main

import (
    "fmt"
    "os"
    "strconv"
    "time"
    "Ej1/prime"
)


func main() {
    if len(os.Args) != 2 {
        fmt.Println("Uso: go run main.go N")
        return
    }

    N, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("N debe ser un nro entero valido")
        return
    }

    start := time.Now()
    singlePrimes := prime.SingleRoutine(N)
    singleTime := time.Since(start)
    fmt.Println("Números primos utilizando una única goroutine:", singlePrimes)

    numGoroutines := 4
    start = time.Now()
    multiPrimes := prime.MultipleRoutine(N, numGoroutines)
    multiTime := time.Since(start)
    fmt.Println("\nNúmeros primos utilizando", numGoroutines, "goroutines:", multiPrimes)
    

    speedUp := float64(singleTime) / float64(multiTime)
    fmt.Println("Tiempo con una única goroutine:", singleTime)
    fmt.Println("Tiempo con", numGoroutines, "goroutines:", multiTime)
    fmt.Printf("\nSpeed-up: %.2f\n", speedUp)
}
