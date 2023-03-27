package main

import "fmt"

func main() {

	aprensentacaoLanchonete()

	switch leComandoExecutado() {
	case 1:
		validaPedidoExistente()
	case 2:

	}

}

func validaPedidoExistente() {
	var lePizza int
	fmt.Scan(&lePizza)
	fmt.Println("Em nosso cardápio tem as seguintes pizza. ")
	fmt.Println("1 - Pizza de Calabresa")
	fmt.Println("2 - Pizza de Calabresa")
	fmt.Println("3 - Pizza de Calabresa")
	fmt.Println("0 - Para falar com um de nossos atendentes")

	fmt.Println("Digite a opção para prosserguimos com o atendimento")

	if lePizza == 1 {
		fmt.Println("Obrigado por escolher nossa Pizza de Calabresa Iremos prepara-la com amor e")
	}
}

func aprensentacaoLanchonete() {
	var leNome string
	fmt.Scan(&leNome)
	fmt.Println("Olá, seja bem-vindo a nossa Pizzaria!")
	fmt.Println("Qual seu nome? ", leNome)

	fmt.Println("----------------//----------//-----------")
}

func leComandoExecutado() int {
	var valorComando int
	fmt.Scan(&valorComando)
	fmt.Println("O Comando escolhido foi. ", valorComando)
	fmt.Println("----------------//----------//-----------")

	return valorComando
}
