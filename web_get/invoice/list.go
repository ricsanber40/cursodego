package invoice

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//ListInvoice retorna a lista de invoices
func ListInvoice() {

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	request, err := http.NewRequest("GET", "http://testing-api.billingdelivery.com/api/invoices?timestamp=0", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o Google Brasil. Erro: ", err.Error())
		return
	}

	request.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjE0YTFlN2VlMjUzZWNjMmYxMmQ2NjNlOWI4MDUyMDQ4ZGQ2NDI1MWRmZTBlNGY5N2UyNDIxMTA0YmJjZjcwYjJmM2UyNjNmNzA2YTZlZTQ5In0.eyJhdWQiOiIyIiwianRpIjoiMTRhMWU3ZWUyNTNlY2MyZjEyZDY2M2U5YjgwNTIwNDhkZDY0MjUxZGZlMGU0Zjk3ZTI0MjExMDRiYmNmNzBiMmYzZTI2M2Y3MDZhNmVlNDkiLCJpYXQiOjE1MjQwMDA0ODgsIm5iZiI6MTUyNDAwMDQ4OCwiZXhwIjoxNTI1Mjk2NDg4LCJzdWIiOiIzNzEiLCJzY29wZXMiOltdfQ.Tj4MYYXgKh3sASHTdpHpAmJyGTXaw21h9od0RcvItSt98SDt-mvjE1O5hU1fvkzvX42IYIRBhWVV05Rpt_ntJcCTZLObQczBp4C_bREFLOog7jMHAZvi9mmcuL8eUyt9Lc62EvoipQeSMXIq_BvXtcPQlQzFlFpicWLWcb92xg0-MS6hgZ29_vxfVSwQ5eJkvbFbwqoKma2D6QtNexEyZdB6Bue-OuzSYyzGedrMl6xaWVvBs2-VfsUR9q1g_QYEITkgXwfVzYNQ63BqcTnADwsszPEVJvgL4gmrIUQEiuWvySZPcyMcDZW9z9GooZ7ITdVkcnnF7OldlNzOsBta6uSOcG-NWojmi-4SY61y_Cpt27wymKWIsDZObpKW0AL40UGaEsIUXB1HkysZPyj4xHUi78ApQTcaR5Cj2SbYrf5j_ooCS6JOzlKJjOaqMdX2HGp4ckL8Btpg4fLXLmCFT1rbpoPwpUFNRdxGbIowHV2pL32z7EIehk7RGEcjOfxQN4PwSd_GVdniHwPVUudDtBst73gf2WvjpKUji17UMKp809AsOTpHjBPk2OKWQrVDYGoBwRAc55UM7gddMKQqMBpIrGwp2yVZ6x4cDPwpOjedR646JE62euNxCYRqsqbdhLH4zj-WuuTKVUnwBwNl1u8SAJpm1Zg5UdSKkU65OqQ")
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao listar invoices. Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	corpo, err := ioutil.ReadAll(resposta.Body)
	if err != nil {
		fmt.Println("[main] Erro ao ler o conteudo da lista de invoice. Erro: ", err.Error())
		return
	}
	fmt.Println(" ")
	fmt.Println(string(corpo))

}
