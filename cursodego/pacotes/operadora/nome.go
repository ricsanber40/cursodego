package operadora

import (
	"projetos/cursodego/pacotes/prefixo"
	"strconv"
)

//NomeOperadora comentário
var NomeOperadora = "XPTO TELECOM"

//PrefixoDaCapitalOperadora  comentario
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + "" + NomeOperadora
