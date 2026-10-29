package recetario

type ProcesoEleccion struct {
	alergias      []string
	intolerancias []string
	dieta         string
	inventario    []Alimento
	limiteMinutos int
}
