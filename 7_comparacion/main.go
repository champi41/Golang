// Implementacion: Divide y venceras. Suma de valores de un arreglo
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func suma(parte []int, ch chan int) {
	sumaP := 0
	for i := 0; i < len(parte); i++ {
		sumaP += parte[i]
	}
	ch <- sumaP
}

func main() {
	numPartes := 8
	resultados := make(chan int, numPartes)
	var total int
	n := 200_000_000
	cantidad := n / numPartes
	arreglo := make([]int, n)

	for i := 0; i < n; i++ {
		arreglo[i] = rand.IntN(100)
	}

	// --- Version secuencial ---
	inicioSec := time.Now()
	var totalSecuencial int
	for i := 0; i < n; i++ {
		totalSecuencial += arreglo[i]
	}
	tiempoSec := time.Since(inicioSec)

	// --- Version concurrente (divide y vencerás) ---
	inicioConc := time.Now()

	for j := 0; j < numPartes; j++ {
		inicio := j * cantidad
		fin := inicio + cantidad
		if j == numPartes-1 {
			fin = n
		}
		go suma(arreglo[inicio:fin], resultados)
	}

	for k := 0; k < numPartes; k++ {
		total += <-resultados
	}
	tiempoConc := time.Since(inicioConc)

	// --- Resultados ---
	fmt.Println("numPartes:", numPartes)
	fmt.Println("Total secuencial:", totalSecuencial, "| tiempo:", tiempoSec)
	fmt.Println("Total concurrente:", total, "| tiempo:", tiempoConc)
	fmt.Println("¿Coinciden?", total == totalSecuencial)
}