package main

import (
	"fmt"
)
func sumarProductos(a, b, c int, chSuma chan int){
	suma := a + b + c
	chSuma <- suma
	fmt.Println("[sumarProductos] El resultado de la suma es:", suma, ". enviado al canal.")
}
func calcularIva(iva float32, chSuma, chIva, chNeto chan int){
	neto := <- chSuma
	impuesto := int(float32(neto)*iva)
	chIva <- impuesto
	chNeto <- neto
	fmt.Println("[calcularIva] El impuesto es:", impuesto, ". Enviado al canal")
}
func calcularBruto(chNeto, chIva, chBruto chan int){
	impuesto := <- chIva
	neto := <- chNeto
	bruto:= neto + impuesto
	chBruto <- bruto
	fmt.Println("[calcularBruto] Calculo enviado al canal")
}
func main(){
	const iva float32 = 0.19
	chSuma := make(chan int)
	chNeto := make(chan int)
	chIva := make(chan int)
	chBruto := make(chan int)

	go sumarProductos(3000, 2500, 1000, chSuma)
	go calcularIva(iva, chSuma, chIva, chNeto)
	go calcularBruto(chNeto, chIva, chBruto)

	final := <- chBruto

	fmt.Println("[main] El valor total bruto sera:", final)

}