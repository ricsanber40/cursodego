package model

//Galinha representa uma ave do tipo galinha
type Galinha interface {
	Cacareja() string
}

//Pata representa uma ave do tipo pato
type Pata interface {
	Quack() string
}

//Ave pode ser galinha ou pata
type Ave struct {
	Nome string
}

//Cacareja representa  som da galinha
func (a Ave) Cacareja() string {
	return "Cócóricó..."
}

//Quack representa  som da pata
func (a Ave) Quack() string {
	return "Quack quack..."
}
