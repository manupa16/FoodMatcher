package recetario

import (
	"errors"
	"slices"
)

type ProcesoEleccion struct {
	alergias      []string
	intolerancias []string
	dieta         string
	inventario    []Alimento
	limiteMinutos int
}

func NuevoProcesoEleccion(alergias []string, intolerancias []string, dieta string, inventario []Alimento, limiteMinutos int) (*ProcesoEleccion, error) {
	if limiteMinutos <= 0 {
		return nil, errors.New("el límite de tiempo tiene que ser un número de minutos mayor que cero")
	}
	return &ProcesoEleccion{
		alergias:      slices.Clone(alergias),
		intolerancias: slices.Clone(intolerancias),
		dieta:         dieta,
		inventario:    slices.Clone(inventario),
		limiteMinutos: limiteMinutos,
	}, nil
}
