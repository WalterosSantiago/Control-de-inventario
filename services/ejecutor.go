package services

import (
	"fmt"
)

func EjecutarComparacion() error {

	// Leer inventario base
	baseline, err := LeerEquipos("data/baseline.json")
	if err != nil {
		return err
	}

	// Leer inventario actual
	actual, err := LeerEquipos("data/actual.json")
	if err != nil {
		return err
	}

	// Comparar
	cambios := Comparar(baseline, actual)

	// Guardar historial
	err = GuardarHistorial(cambios, "data/historial.json")
	if err != nil {
		return err
	}

	// Mostrar resumen
	fmt.Println("\n==============================")
	fmt.Println("RESUMEN DE LA COMPARACIÓN")
	fmt.Println("==============================")
	fmt.Println("Equipos base   :", len(baseline))
	fmt.Println("Equipos actual :", len(actual))
	fmt.Println("Cambios        :", len(cambios))

	if len(cambios) == 0 {
		fmt.Println("No se detectaron cambios.")
	} else {
		fmt.Println("\nCambios detectados:")

		for _, cambio := range cambios {
			fmt.Printf("- [%s] %s", cambio.Tipo, cambio.EquipoID)

			if cambio.Tipo == "MODIFICACION" {
				fmt.Printf(" | %s: %s -> %s",
					cambio.Campo,
					cambio.Esperado,
					cambio.Actual)
			}

			fmt.Println()
		}
	}

	fmt.Println("==============================")

	return nil
}
