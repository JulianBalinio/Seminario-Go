package prime

import (
    "math"
    "sync"
)

func isPrime(num int) bool {
    if num <= 1 {
        return false
    }
    if num == 2 || num == 3 {
        return true
    }
    if num%2 == 0 {
        return false
    }
    sqrt := int(math.Sqrt(float64(num)))
    for i := 3; i <= sqrt; i += 2 {
        if num%i == 0 {
            return false
        }
    }
    return true
}

func SingleRoutine(N int) []int {
    primes := []int{}
    for i := 2; i <= N; i++ {
        if isPrime(i) {
            primes = append(primes, i)
        }
    }
    return primes
}

func producer(start, end int, result chan <- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i:= start; i <= end; i++ {
        if isPrime(i) {
            result <- i
        }
    }
}

func MultipleRoutine(N, numGoroutines int) []int {
    var wg sync.WaitGroup
    result := make(chan int, N)

    //Podria simplificarse a N/numGoroutines pero esti asumiria que N sea multiplo exacto
    rangeSize := (N + numGoroutines - 1) / numGoroutines

    for i:= 0; i < numGoroutines; i++ {
        //Calculo para el rango de cada goroutine
        start := (i * rangeSize) + 1
        end := start + rangeSize
        if end > N {
            end = N
        }
        wg.Add(1)
        go producer(start, end, result, &wg)
    }

    
    go func () {
        wg.Wait()
        close(result)
    }()

    var primes [] int
    for prime:= range result {
        primes = append(primes, prime)
    }
    return primes
}
