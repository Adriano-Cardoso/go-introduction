package main

import (
	"fmt"
	"os"
)

func main() {

	aprensentacaoLanchonete()

	exibeMenu()

	switch leComandoExecutado() {
	case 1:
		fmt.Println("Obrigado por escolher nossa Pizza de Calabresa, Iremos prepara-la com amor e carinho")
	case 2:
		fmt.Println("Obrigado por escolher nossa Pizza Portuguesa, Iremos prepara-la com amor e carinho")
	case 3:
		fmt.Println("Obrigado por escolher nossa Pizza de Mussarela, Iremos prepara-la com amor e carinho")
	case 0:
		fmt.Print("Iremos te transferir para nosso atendente")
		os.Exit(0)
	default:
		fmt.Println("Não conheço o comando")
		os.Exit(-1)
	}

}

func exibeMenu() {
	fmt.Println("Em nosso cardápio tem as seguintes pizza. ")
	fmt.Println("1 - Pizza de Calabresa")
	fmt.Println("2 - Pizza Portuguesa")
	fmt.Println("3 - Pizza Mussarela")
	fmt.Println("0 - Para falar com um de nossos atendentes")

	fmt.Println("Digite a opção para prosserguimos com o atendimento")

}

func aprensentacaoLanchonete() string {
	leNome := "Adriano"
	fmt.Println("Olá", leNome, "! seja bem-vindo a nossa Pizzaria!")
	fmt.Println("----------------//----------//-----------")
	return leNome
}

func leComandoExecutado() int {
	var valorComando int
	fmt.Scan(&valorComando)
	fmt.Println("O Comando escolhido foi. ", valorComando)
	fmt.Println("----------------//----------//-----------")

	if valorComando == 1 {

	}

	return valorComando
}
