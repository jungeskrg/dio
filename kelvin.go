package main

import "fmt"

const ebulicaoK = 373.0

func main() {
	tempK := ebulicaoK
	tempC := (tempK - 273)
	fmt.Printf("O ponto de ebulição da água em ºC é: %g e em ºK é: %g", tempC, tempK)
}
