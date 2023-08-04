package main

import "fmt"
import "os"
//import "net/http"
//import "reflect"


func main(){

	nome, idade := devolveNomeEIdade()
	fmt.Println("Nome: ",nome, "Idade: ",idade)

	//Forma de ignorar um retorno
	//_, idade := devolveNomeEIdade()
	//fmt.Println( "Idade: ",idade)
	

	//Chamada de função
	comando := leComando()



}

func condicionas(){

	var comando int

	//////////COndicionais if//////////////
	if comando == 1 {
		fmt.Println("Monitorando ...")
	}else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	}else if comando == 0{
		fmt.Println("Saindo...")
	}else{
		fmt.Println("Comando Inválido")
	}
	


	switch comando {

	case 1:		
		monitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Comando Inválido")
		os.Exit(-1)
		return
	}
}


/*
	Multiplas formas de retorno
*/
func devolveNomeEIdade() (string, int) {
	nome := "aDSON"
	idade := 24

	return nome , idade
}

func loop(){
	// o For faz um loop infinito
	for{

	}
}


func arays(){

	//Forma 1 Slice
	nomes := []string{"Adson","Souza","Jesus"}
	fmt.Println(nomes)

	//Forma 2 array
	var sites [3]string
	site[0] = "..."


	//Adicionar item no array
	nomes = append(nomes,"Silva")
}

