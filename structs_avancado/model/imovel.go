package model

import "errors"

//Imovel é um struct
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

/*ErrValorDeImovelInvalido valor inválido*/
var ErrValorDeImovelInvalido = errors.New("O valor informado não é válido para um imóvel")

/*ErrValorDeImovelAlto valor inválido*/
var ErrValorDeImovelAlto = errors.New("O valor informado é muito alto para um imóvel")

//SetValor define o valor do imóvel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorDeImovelAlto
		return
	}
	i.valor = valor
	return
}

//GetValor retorna o valor do Imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
