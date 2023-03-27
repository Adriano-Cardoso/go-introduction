package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	for {

		exibiMenu()

		switch leComando() {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs do programa...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)

		}

	}

}

func exibeIntroducao() {
	nome := "Adriano"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("------------------------------------")
}

func leComando() int {
	var leComando int
	fmt.Scan(&leComando)
	fmt.Println("O comando escolhido foi: ", leComando)

	return leComando
}

func exibiMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println("------------------------------------")

}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "Erro ao carregar o site. Status code: ", resp.StatusCode)
	}

}

func devolveNomeEidade() (string, int) {
	nome := "Adriano"
	idade := 24
	return nome, idade
}
