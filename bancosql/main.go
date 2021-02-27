package main

import (
	"cursogo/bancosql/manipulador"
	"cursogo/bancosql/repro"
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("Vamos começar a subir o servidor....")
}

func main() {

	err := repro.AbreConexaoSQL()
	if err != nil {
		fmt.Println("Erro ao abrir banco de dados", err.Error())
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá mundo!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)

	fmt.Println("Estou escutando......")
	http.ListenAndServe(":8181", nil)

}
