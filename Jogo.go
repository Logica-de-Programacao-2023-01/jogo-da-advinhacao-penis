package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	jogar()
}

func jogar() {
	fmt.Println("Bem-vindo ao Jogo de Adivinhação!")
	fmt.Println("Tente adivinhar o número entre 1 e 100.")

	var todasTentativas [][]int

	for {
		tentativas := jogarJogo()

		todasTentativas = append(todasTentativas, tentativas)

		if !jogarNovamente() {
			break
		}
	}

	PrintarTodasTentativas(todasTentativas)
}

func jogarJogo() []int {
	resposta := rand.Intn(100) + 1
	fmt.Println("Novo jogo iniciado.")

	var tentativas []int

	for {
		fmt.Print("Digite um número: ")
		var palpiteStr string
		_, err := fmt.Scanln(&palpiteStr)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número inteiro.")
			continue
		}

		palpite, err := strconv.Atoi(palpiteStr)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número inteiro.")
			continue
		}

		tentativas = append(tentativas, palpite)

		if palpite < resposta {
			fmt.Println("O número é maior. Tente novamente!")
		} else if palpite > resposta {
			fmt.Println("O número é menor. Tente novamente!")
		} else {
			fmt.Printf("Parabéns! Você acertou o número em %d tentativas.\n", len(tentativas))
			break
		}
	}

	return tentativas
}

func jogarNovamente() bool {
	var escolha string
	fmt.Print("Você deseja jogar novamente? (s/n): ")
	_, err := fmt.Scanln(&escolha)
	if err != nil || (escolha != "s" && escolha != "n") {
		fmt.Println("Entrada inválida. Por favor, digite 's' para jogar novamente ou 'n' para sair.")
		return jogarNovamente()
	}

	if escolha == "s" {
		return true
	} else {
		fmt.Println("Obrigado por jogar! Até a próxima.")
		return false
	}
}

func PrintarTodasTentativas(todasTentativas [][]int) {
	fmt.Println("Histórico de tentativas:")
	for i, tentativas := range todasTentativas {
		fmt.Printf("Jogada %d: %d tentativas\n", i+1, len(tentativas))
	}
}
