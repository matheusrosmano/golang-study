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

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs ...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa ...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando.")
			os.Exit(-1)
		}
		fmt.Println("")
	}
}

func exibeIntroducao() {
	nome := "Matheus"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	fmt.Println("O número digitado foi: ", comandoLido)
	fmt.Println("")

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando ...")
	sites := leSitesArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			siteOk, codigo := testeSite(site)
			if siteOk {
				fmt.Println("Site '", site, "' carregado com sucesso.")
				registraLog(site, true)
			} else {
				fmt.Println("Site '", site, "' com problema. Código:", codigo)
				registraLog(site, false)
			}
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testeSite(site string) (bool, int) {
	resp, _ := http.Get(site)
	return resp.StatusCode == 200, resp.StatusCode
}

func leSitesArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
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

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(arquivo)

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if nil != err {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}