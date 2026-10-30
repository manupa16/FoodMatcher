package recetario

type Alimento struct {
	nombre    string
	alergenos map[Alergeno]struct{}
}
