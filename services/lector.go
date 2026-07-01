package services

import (
	"encoding/json"
	"os"

	"inventario/models"
)

// LeerEquipos lee un archivo JSON y devuelve una lista de equipos.
func LeerEquipos(ruta string) ([]models.Equipo, error) {

	// Leer todo el contenido del archivo
	datos, err := os.ReadFile(ruta)
	if err != nil {
		return nil, err
	}

	// Variable donde se guardarán los equipos
	var equipos []models.Equipo

	// Convertir JSON -> []Equipo
	err = json.Unmarshal(datos, &equipos)
	if err != nil {
		return nil, err
	}

	return equipos, nil
}
