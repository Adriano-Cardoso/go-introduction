package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Definir os nomes dos times
	teamNames := []string{"Time 1", "Time 2", "Time 3", "Time 4"}

	// Definir os nomes dos jogadores
	playerNames := []string{
		"Jogador 1",
		"Jogador 2",
		"Jogador 3",
		"Jogador 4",
		"Jogador 5",
		"Jogador 6",
		"Jogador 7",
		"Jogador 8",
		"Jogador 9",
		"Jogador 10",
		"Jogador 11",
		"Jogador 12",
		"Jogador 13",
		"jogador 14",
		"Jogador 15",
		"Jogador 16",
		"Jogador 17",
		"Jogador 18",
		"Jogador 19",
		"Jogador 20",
		"Jogador 21",
		"Jogador 22",
		"Jogador 23",
		"Jogador 24",
	}

	// Inicializar a seed do gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	// Criar um mapa para armazenar os times e seus jogadores
	teams := make(map[string][]string)

	// Sortear aleatoriamente os jogadores para cada time
	for i := 0; i < len(teamNames); i++ {
		// Criar uma fatia vazia para armazenar os jogadores do time atual
		teams[teamNames[i]] = make([]string, 0)

		// Sortear aleatoriamente 6 jogadores para o time atual
		for j := 0; j < 4; j++ {
			if len(playerNames) == 0 {
				// Se a fatia de jogadores estiver vazia, interromper o loop interno
				break
			}

			// Escolher um jogador aleatoriamente da lista de jogadores disponíveis
			index := rand.Intn(len(playerNames))
			player := playerNames[index]

			// Adicionar o jogador escolhido ao time atual
			teams[teamNames[i]] = append(teams[teamNames[i]], player)

			// Remover o jogador escolhido da lista de jogadores disponíveis para evitar que seja escolhido novamente
			playerNames = append(playerNames[:index], playerNames[index+1:]...)
		}
	}

	// Imprimir os times e seus jogadores
	for team, players := range teams {
		fmt.Printf("%s: %v\n", team, players)
	}
}
