package main

import (
	"fmt"
	"strconv" // Paquete para conversion de strings
)

func main() {
	edad := 22

	edad_str := strconv.Itoa(edad) // Conversion de Integer to Array(String)

	fmt.Println("La edad es " + edad_str)
}
