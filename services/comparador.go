package services

import (
	"inventario/models"
	"time"
)

func agregarCambio(cambios *[]models.Cambio, equipoID, campo, esperado, actual string) {

	cambio := models.Cambio{
		Fecha:    time.Now().Format("2006-01-02 15:04:05"),
		EquipoID: equipoID,
		Tipo:     "MODIFICACION",
		Campo:    campo,
		Esperado: esperado,
		Actual:   actual,
	}

	*cambios = append(*cambios, cambio)
}

// Comparar recibe el inventario base y el inventario actual.
// Más adelante devolverá todos los cambios encontrados.
func Comparar(baseline []models.Equipo, actual []models.Equipo) []models.Cambio {

	var cambios []models.Cambio
	// Crear un mapa para buscar rápidamente equipos del baseline
	baselineMap := make(map[string]models.Equipo)

	for _, equipo := range baseline {
		baselineMap[equipo.ID] = equipo
	}

	// Recorrer el inventario actual
	for _, equipoActual := range actual {

		// Buscar si existe en el baseline
		equipoBase, existe := baselineMap[equipoActual.ID]

		if !existe {

			cambio := models.Cambio{
				Fecha:    time.Now().Format("2006-01-02 15:04:05"),
				EquipoID: equipoActual.ID,
				Tipo:     "NUEVO",
				Campo:    "",
				Esperado: "",
				Actual:   "",
			}

			cambios = append(cambios, cambio)
		}

		if existe {

			// Comparar RAM
			if equipoBase.RAM != equipoActual.RAM {

				if equipoBase.RAM != equipoActual.RAM {
					agregarCambio(
						&cambios,
						equipoActual.ID,
						"RAM",
						equipoBase.RAM,
						equipoActual.RAM,
					)
				}
			}
			// Comparar Procesador
			if equipoBase.Procesador != equipoActual.Procesador {
				agregarCambio(
					&cambios,
					equipoActual.ID,
					"Procesador",
					equipoBase.Procesador,
					equipoActual.Procesador,
				)
			}

			//Comparar Disco
			if equipoBase.Disco != equipoActual.Disco {
				agregarCambio(
					&cambios,
					equipoActual.ID,
					"Disco",
					equipoBase.Disco,
					equipoActual.Disco,
				)
			}

			//Comparar S.O
			if equipoBase.SistemaOperativo != equipoActual.SistemaOperativo {
				agregarCambio(
					&cambios,
					equipoActual.ID,
					"Sistema Operativo",
					equipoBase.SistemaOperativo,
					equipoActual.SistemaOperativo,
				)
			}

		}
	}

	// Crear un mapa para buscar rápidamente equipos del inventario actual
	actualMap := make(map[string]models.Equipo)

	for _, equipo := range actual {
		actualMap[equipo.ID] = equipo
	}

	for _, equipoBase := range baseline {

		_, existe := actualMap[equipoBase.ID]

		if !existe {

			cambio := models.Cambio{
				Fecha:    time.Now().Format("2006-01-02 15:04:05"),
				EquipoID: equipoBase.ID,
				Tipo:     "ELIMINADO",
				Campo:    "",
				Esperado: "",
				Actual:   "",
			}

			cambios = append(cambios, cambio)
		}
	}

	return cambios
}
