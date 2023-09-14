//vamos criar nosso primeiro site estático?
//passo a passo
//um exemplo simples
//mas real
//código fará: servir arquivos HTML de um local especifico no disco.

//passo1: comece criandoum diretório (pasta para armazenar o projeto) - servidorEstatico
//passo2: criar seu arquivo go - aula4.go
//passo3: criar um arquivo html - example.html (cdm:type nul > aula.html)
//abrir os arquivos e escrever os códigos (deixarei no ppt da aula)

// pacote principal
package main

//importar os pacotes que usaremos: net/http e log
import (
	"log"
	"net/http"
)
// teremos uma função principal
func main() {
	fs := http.FileServer(http.Dir("./estatico"))
	http.Handle("/", fs)
	log.Print("Listening on: 3000...", nil)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
		
	}
}