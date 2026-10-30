package recetario

type ProcesoEleccion struct {
	alergias      []string
	intolerancias []string
	dieta         Dieta
	inventario    []Alimento
	limiteMinutos int
}
