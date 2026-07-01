package main

import (
	"log"

	"inventario/services"
)

func main() {

	err := services.VigilarArchivo("data")
	if err != nil {
		log.Fatal(err)
	}

}
