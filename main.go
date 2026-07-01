package main

import (
	"fmt"
	"log"

	"inventario/services"
)

func main() {

	// Leer inventario base
	baseline, err := services.LeerEquipos("data/baseline.json")
	if err != nil {
		log.Fatal(err)
	}

	// Leer inventario actual
	actual, err := services.LeerEquipos("data/actual.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("===== INVENTARIO BASE =====")
	for _, equipo := range baseline {
		fmt.Printf("%s - %s\n", equipo.ID, equipo.Hostname)
	}

	fmt.Println()

	fmt.Println("===== INVENTARIO ACTUAL =====")
	for _, equipo := range actual {
		fmt.Printf("%s - %s\n", equipo.ID, equipo.Hostname)
	}

	cambios := services.Comparar(baseline, actual)

	fmt.Println()
	fmt.Println("Cantidad de cambios encontrados:", len(cambios))
	for _, cambio := range cambios {

		fmt.Println("----------------------------")
		fmt.Println("Fecha:", cambio.Fecha)
		fmt.Println("Equipo:", cambio.EquipoID)
		fmt.Println("Tipo:", cambio.Tipo)
		fmt.Println("Campo:", cambio.Campo)
		fmt.Println("Esperado:", cambio.Esperado)
		fmt.Println("Actual:", cambio.Actual)
	}

	fmt.Println("\n=== CAMBIOS DETECTADOS EN ESTA EJECUCIÓN ===")

	for _, cambio := range cambios {
		fmt.Printf("%+v\n", cambio)
	}
	err = services.GuardarHistorial(cambios, "data/historial.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("Historial guardado correctamente.")
}
