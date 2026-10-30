package recetario

type ProcesoEleccion struct {
	alergias      []Alergeno
	intolerancias []string
	dieta         Dieta
	inventario    []Alimento
	limiteMinutos int
}
