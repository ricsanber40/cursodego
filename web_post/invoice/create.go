package invoice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"web_post/model"
)

//Create acessa rota de criar invoice
func Create() {

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	invoice := model.Invoice{}
	invoice.DueDate = "2018-06-10"
	invoice.Value = 500.99
	invoice.Description = "Conta Escola Go"
	invoice.CodeLine = "75691333380101376770000434520011973080000032174"
	invoice.CompanyIdentifier = "28038705000117"

	conteudoEnviar, err := json.Marshal(invoice)
	if err != nil {
		fmt.Println("[main] Erro ao gerar o JSON do objeto invoice. Erro: ", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "http://testing-api.billingdelivery.com/api/invoices", bytes.NewBuffer(conteudoEnviar))
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para rota de criar invoice. Erro: ", err.Error())
		return
	}
	request.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImJmODZmYjRkZmIwNjc3YmM1YjkzODI3ODJiZDM2ZTY1OWQwYjIxNzY0MzZkYTllMjcwYzM5NWI4N2FlNWUyNmZiODc3NzZjMGVhNDVlNTVmIn0.eyJhdWQiOiIyIiwianRpIjoiYmY4NmZiNGRmYjA2NzdiYzViOTM4Mjc4MmJkMzZlNjU5ZDBiMjE3NjQzNmRhOWUyNzBjMzk1Yjg3YWU1ZTI2ZmI4Nzc3NmMwZWE0NWU1NWYiLCJpYXQiOjE1MjQwMDIxMDcsIm5iZiI6MTUyNDAwMjEwNywiZXhwIjoxNTI1Mjk4MTA3LCJzdWIiOiIzNzMiLCJzY29wZXMiOltdfQ.rS5axGBiZKLUiVQwQiFq393b9ruQUkvk3dLKM5fiTRqTgB-DrgD1AwV_zVKhFZQOeLWntOJL5ZmA8owXsk4x5ADvwdnqqGOWxRkSzKC71TH-8VU46xMgNCEVh2PQadum4zpdbV8LcJ2klavLw-XKJSgL61chwEEdGaxuCctGcRlXESoIrm4IxyVyYGCZAyYIKuo_zd9pBUZZeUuZhSgBxENV8L3e--YMbSXdkzVJ2oprQniedOpu8d0xNYtXB0xXYk-0sgkNkc9Dnacyrqo_WQbM1I65eKg166t4ClpXQvWfN7LuuN_7qqiemlRmXFOpcYuRu-o27wc5ABkKfWq2vk64zi6p2i_T0xcvvPoHTVOC_0n2IYlVGLNvF7b9HzVoLt-iwS8DznQQteGpcRQri3X9kuw9heMMemBSOSwjEtcSFr1ELQMCqRoo6dUvUQPPRS-it9YEFD__ZEkuYwY41UwZ71mhOWqendWglDwy0zXGnllbqct_WqdQ1WL2XYog-u2OAU91LKsgNIB50QDAMRmbqQO2kfMDCz1yhQxcd6V0nisJivkiy7ZDUKJjkISdwJYLI1QrFsE_I7tFuhmJAThps1NZ0ynPyM83crlGDVbEdczNSs8KrAe4W-U_k3fi89eXv8PRsBw5BIRhaCLEc5aTz5Hjl4Is_ciOXsnUQaE")
	request.Header.Set("content-type", "application/json; charset=utf-8")
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao chamar o servico de criar invoice. Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	corpo, err := ioutil.ReadAll(resposta.Body)
	if err != nil {
		fmt.Println("[main] Erro ao ler o conteudo retornado pela rota de criar invoice. Erro: ", err.Error())
		return
	}
	fmt.Println(" ")
	fmt.Println(string(corpo))

}
