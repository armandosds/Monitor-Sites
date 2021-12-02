package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	introducao()
	for {
		parametros()

		comando := lerComandos()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}

}
func introducao() {
	fmt.Println("Olá, eu sou o Monitoramento de Sites")
	fmt.Println("Este programa irá verificar o status de um site")
	fmt.Println("Digite o seu nome.")
	var nome string
	fmt.Scan(&nome)
	versao := 1.1
	fmt.Println("Olá Sr(a).", nome)
	fmt.Println("Este programa está na versão", versao, "e foi desenvolvido por", "Armando Soares")
}

func parametros() {
	fmt.Println("Digite o número referente ao comando desejado")
	fmt.Println("1-> Iniciar Monitoramento")
	fmt.Println("2-> Exibir Logs")
	fmt.Println("0-> Sair do Programa")
}

func lerComandos() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	return comando
}

func iniciarMonitoramento() {
	var site string

	fmt.Println("Digite o site que deseja monitorar")
	fmt.Println("Você deve iniciar sua busca com http://")
	fmt.Scan(&site)
	fmt.Println("Monitorando...")
	resposta, _ := http.Get(site)

	fmt.Println("Codigo Status:", resposta.StatusCode)

	switch resposta.StatusCode {
	case 200:
		fmt.Println("Tudo certo,", site, "está ativo!")
	default:
		fmt.Println("Algo deu errado, o problema é:", resposta.Status)
		os.Exit(-1)
	}
}
