package models

type Cambio struct {
	Fecha    string `json:"fecha"`
	EquipoID string `json:"equipo_id"`
	Tipo     string `json:"tipo"`
	Campo    string `json:"campo"`
	Esperado string `json:"esperado"`
	Actual   string `json:"actual"`
}
