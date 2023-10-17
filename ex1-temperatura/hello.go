package main

import "fmt"

const ebulicaoCelcius = 100.0;

func main() {
    kelvin := ebulicaoCelcius - 273;
    fmt.Printf("Ponto de Ebulição para Kelvin: %g", kelvin)
}
