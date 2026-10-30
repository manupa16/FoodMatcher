package recetario

type ProcesoEleccion struct {
	alergias      []Alergeno
	intolerancias []Intolerancia
	dieta         Dieta
	inventario    []Alimento
	limiteMinutos int
}
