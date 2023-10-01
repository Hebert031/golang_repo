package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramento = 5 // Número de vezes que o monitoramento é executado.
const tempo = 7         // Tempo de espera entre os ciclos de monitoramento (em segundos).
const tempo2 = 1        // Tempo de espera entre as verificações individuais dos sites (em segundos).

func main() {
	// Chama a função exibeIntroducao() para mostrar uma saudação.
	exibeIntroducao()

	for {
		// Exibe o menu de opções.
		exibeMenu()
		// Lê o comando inserido pelo usuário.
		comando := leComando()

		switch comando {
		case 1:
			// Se o comando for 1, inicia o monitoramento de sites.
			iniciarMonitoramento()
		case 2:
			// Se o comando for 2, exibe logs (ainda não implementado neste código).
			fmt.Println("Exibindo Logs...")
		case 0:
			// Se o comando for 0, encerra o programa.
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			// Se o comando não for reconhecido, exibe uma mensagem de erro e encerra o programa com código de erro -1.
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	// Define variáveis nome e versao.
	nome := "Hebert"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	// Exibe as opções do menu.
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	// Lê um número inteiro inserido pelo usuário.
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// Lista de sites a serem monitorados.
	sites := []string{"http://www.igornetoadv.com.br/", "http://www.theends.com.br", "http://www.casahomecare.com.br", "http://www.hstech.cloud", "http://www.igornetoadv.adv.br/"}
	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println(i, ":", site)
			time.Sleep(tempo2 * time.Second)
			testaSite(site)
		}
		time.Sleep(tempo * time.Second)
	}
	os.Exit(0)
}

// Função para testar a resposta HTTP de um site.
func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		// Em caso de erro, imprime a mensagem de erro.
		fmt.Println("Erro ao acessar o site:", site)
		fmt.Println(err)
		return
	}

	if resp.StatusCode == 200 {
		// Se o status da resposta for 200, o site está OK.
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		fmt.Println(site, ":", resp)
	} else {
		// Caso contrário, o site está com algum problema.
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
