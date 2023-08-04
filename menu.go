package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//import "io/ioutil"
//import "reflect"

const monitoramentos = 5
const delay = 10

func main() {

	for {

		exibeMenu()

		comando := leComando()

		switch comando {

		case 1:
			monitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Comando Inválido")
			os.Exit(-1)
			return
		}
	}

}

func leComando() int {
	var comando int

	fmt.Scan(&comando)

	return comando

}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")
}

func monitoramento() {

	sites := leSitesdoArquivo()

	fmt.Println("Monitorando ...")

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site ", i+1, ":", site)

			resp, err := http.Get(site)

			if err != nil {
				fmt.Println("Ocorreu um erro ", err)

			}

			if resp.StatusCode == 200 {
				fmt.Println("Site ", site, " está operante!")
				registraLog(site, true)
			} else {
				fmt.Println("Site ", site, " está inoperante! Status :", resp.StatusCode)
				registraLog(site, false)
			}
		}

		fmt.Println("")

		time.Sleep(delay * time.Second)
	}

	fmt.Println("")

}

func leSitesdoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ", err)

	}

	leitor := bufio.NewReader(arquivo)

	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	fmt.Println(sites)

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/06/2006 15:00:05") + "-" + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
