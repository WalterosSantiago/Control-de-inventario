package models

type Equipo struct {
	ID               string `json:"id"`
	Hostname         string `json:"hostname"`
	Procesador       string `json:"procesador"`
	RAM              string `json:"ram"`
	Disco            string `json:"disco"`
	SistemaOperativo string `json:"sistema_operativo"`
}
