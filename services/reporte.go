package services

import (
	"encoding/json"
	"os"

	"inventario/models"
)

func existeCambio(historial []models.Cambio, nuevo models.Cambio) bool {

	for _, cambio := range historial {

		if cambio.EquipoID == nuevo.EquipoID &&
			cambio.Tipo == nuevo.Tipo &&
			cambio.Campo == nuevo.Campo &&
			cambio.Esperado == nuevo.Esperado &&
			cambio.Actual == nuevo.Actual {

			return true
		}
	}

	return false
}

func GuardarHistorial(cambios []models.Cambio, ruta string) error {

	// Historial existente
	var historial []models.Cambio

	// Intentar leer el archivo si existe
	datos, err := os.ReadFile(ruta)

	if err == nil {

		// Si existe, convertir JSON -> []Cambio
		json.Unmarshal(datos, &historial)

	}

	// Agregar los cambios nuevos
	for _, cambio := range cambios {

		if !existeCambio(historial, cambio) {

			historial = append(historial, cambio)

		}

	}

	// Convertir nuevamente a JSON
	datos, err = json.MarshalIndent(historial, "", "    ")
	if err != nil {
		return err
	}

	// Guardar
	err = os.WriteFile(ruta, datos, 0644)
	if err != nil {
		return err
	}

	return nil
}
