package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	//Mensagem que vai aparecer "seria o nome do log"
	log.SetPrefix("greetings: ")
	//Comando para determinar o tipo do log, iniciando no 0
	log.SetFlags(2)

	names := []string{"William", "Cavalo", "Wagner"}
	/*
		Variavel para receber a função do greetings.go,
		Essa variavel pode puxar a função "Hello" ou a função "Hellos"
		a primeira para retornar uma saldação aleatória para um nome
		e a outra para terornar uma saldação aleatória para varios nomes
	*/
	messages, err := greetings.Hellos(names)
	//IF simples para identificar se o valor da função acima está vazia
	if err != nil {
		//Comando para aparecer o status do erro
		log.Fatal(err)
	}
	//Print para mostrar o resultado no console
	fmt.Println(messages)

}
