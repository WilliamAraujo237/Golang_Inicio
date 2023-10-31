package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
// Função Hello retorna uma saldação para o nome da pessoa
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	//Retorna uma sadação com a mensagem, "Nome Vazio"
	if name == "" {
		return "", errors.New("Nome Vazio")
	} else {
		//Caso o nome for informado retorne a saudação e um erro nulo
		message := fmt.Sprintf(RandolFormat(), name)

		//Esse exemplo abaixo informa uma quebra do teste unitario
		//Não informando um nome para que seja adicionado uma saldação
		//  --> message := fmt.Sprintf(RandolFormat())
		return message, nil
	}

}

// Hellos retorna um mapa que associa cada uma das pessoas nomeadas
// a uma mensagem de saudação
func Hellos(nome []string) (map[string]string, error) {
	//Um mapa para associar um nome a uma mensagem
	messages := make(map[string]string)
	//percorre o slice de nome recebido chamando a função
	//Hello para obter uma mensagem para cada um
	for _, nome := range nome {
		message, err := Hello(nome)

		if err != nil {
			return nil, err
		}
		//No mapa, associa a mensagem recuperada
		//ao nome
		messages[nome] = message
	}
	return messages, nil
}

// Função que armazena um Slice com saudações e retorna uma das saudações de maneira
// aleatoria
func RandolFormat() string {
	formats := []string{
		"Hi %v, greate see you",
		"Hi %v, Welcome",
		"Nice see you %v",
	}
	return formats[rand.Intn(len(formats))]
}
