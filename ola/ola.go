package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const prefixOlaPortugues = "Olá, "
const prefixOlaEspanhol = "Hola, "
const prefixOlaFrances = "Bonjour, "

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixOlaFrances
	case espanhol:
		prefixo = prefixOlaEspanhol
	default:
		prefixo = prefixOlaPortugues
	}
	return
}

func main() {
	fmt.Println(Ola("Mundo", ""))
}
