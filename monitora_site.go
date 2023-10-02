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
			imprimeLogs()
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
	sites := leSitesDoArquivo()
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
		registraLog(site, true)
	} else {
		// Caso contrário, o site está com algum problema.
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	// Cria uma slice vazia para armazenar os sites lidos do arquivo.
	var sites []string

	// Abre o arquivo "sites.txt" para leitura.
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		// Em caso de erro ao abrir o arquivo, imprime uma mensagem de erro.
		fmt.Println("Ocorreu um erro:", err)
	}

	// Cria um leitor para ler o arquivo linha por linha.
	leitor := bufio.NewReader(arquivo)

	// Loop infinito para ler todas as linhas do arquivo.
	for {
		// Lê uma linha do arquivo.
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha) // Remove espaços em branco e caracteres de nova linha.
		sites = append(sites, linha)     // Adiciona a linha à slice de sites.
		if err == io.EOF {
			// Se chegarmos ao final do arquivo, saímos do loop.
			break
		}
	}

	// Fecha o arquivo após a leitura.
	arquivo.Close()

	// Retorna a slice contendo os sites lidos do arquivo.
	return sites
}

func registraLog(site string, status bool) {
	// Abre ou cria o arquivo "log.txt" para escrita, anexando ao final (os.O_APPEND) e com permissão de escrita (0666).
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		// Em caso de erro ao abrir ou criar o arquivo, imprime uma mensagem de erro.
		fmt.Println("Ocorreu um erro:", err)
	}

	// Escreve uma linha no arquivo de log no formato "site - online: true/false".
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	// Fecha o arquivo após a escrita.
	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		// Em caso de erro ao abrir o arquivo, imprime uma mensagem de erro.
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
