package main

import "fmt"
import "reflect"

func main(){

	var nome  = "Adson";
	var versao  = 1.1
	//var idade  = 24
	variaVelNova := "Outra maneira de declarar variável"

	fmt.Println("Hello, ",nome)
	fmt.Println("Versao: ",versao)

	fmt.Println("O timpo da variável nome é ",reflect.TypeOf(nome));

	fmt.Println(variaVelNova)
}