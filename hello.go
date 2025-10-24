package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	//"io/ioutil"
	"bufio"
	"strconv"
)

//import "reflect"

const monitoramentos = 3
const delay = 5

func main() {
	//Se não colocar nehum valor nas variaveis,sempre retorna vazio ou 0
	//O go tambem consegue inferir tipos===
	//var nome = "Bungas" = Ele vai entender que é uma String
	//Tambem é possivel declarar variaveis sem o "var' ==

	// nome := "Bungas"
	//versao := 1.1
	//idade := 22

	//fmt.Println("Esse é o meu patrão:",nome)
	//fmt.Println("Essa é a versão do meu patrão:",versao)

	//var comando int

	//&= Indica o endereço daquela variavel
	//fmt.Scanf("%d", &comando)
	//Tambem pode ser feito dessa forma sem passar o modificador

	//fmt.Scan(&comando)

	//fmt.Println("O comando escolhido foi", comando)

	//condicionais

	//if comando == 1 {
	//  fmt.Println("Monitorando...")
	//} else if comando == 2 {
	// fmt.Println("Exibindo Logs...")
	//} else if comando == 0 {
	// fmt.Println("Saindo do programa...")
	//} else {
	// fmt.Println("Não conheço este comando")
	//}

	//Quando não queremos saber de um dos retornos, queremos ignorá-lo,
	// nós utilizamos o operador de identificador em branco (_):

	// _, idade := devolveNomeEIdade()
	//fmt.Println(idade)

	//Caso o contrário
	//nome, idade := devolveNomeEIdade()
	//fmt.Println(nome, "tem", idade, "anos")

	exibeIntroducao()

	//Não existe while no Go
	for {
		exibeMenu()
		comando := leInput()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			//permite sair do programa de acordo com o numero passado
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			//Indica que ocorreu algum erro,algo inesperado no programa
			os.Exit(-1)
		}

	}

	//fmt.Println("Essa é a idade do meu patrão:",idade)

	//Tambem é possivel fazer assim
	//fmt.Println("Olá sr.", nome, "sua idade é", idade)

	//Mostra o tipo da variavel
	//fmt.Println("O tipo da variável idade é", reflect.TypeOf(versao))
}

// O go permite criar uma função que retorna dois valores, uma string e um inteiro.

//func devolveNomeEIdade() (string, int) {
//nome := "Douglas"
//idade := 24
//return nome, idade
//}

func exibeIntroducao() {
	nome := "Bungas"
	versao := 1.1

	fmt.Println("Esse é o meu patrão:", nome)
	fmt.Println("Essa é a versão do meu patrão:", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leInput() int {
	var comando int

	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
	fmt.Print("")

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	//Arrays em Go
	//var sites [4]string
	//sites[0] = "https://random-status-code.herokuapp.com/"
	//sites[1] = "https://www.alura.com.br"
	//sites[2] = "https://www.anroll.net/"

	sites := leSitesDoArquivo()

	//Pode ser feito assim ou de uma maneira mais facil
	//for i := 0; i < len(sites); i++ {
	//fmt.Println(sites[i])
	//}

	//Range obtem a posiçao e quem esta naquela posição
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Print("")
	}

	fmt.Print("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)

	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site , true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site , false)
	}
}

func leSitesDoArquivo() []string {

	//Comando para abrir arquivos em Go

	sites := []string{}

	arquivo, err := os.Open("sites.txt")

	//Le um arquivo como um todo em bytes 
	//arquivo,err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	//Le o arquivo passo a passo,linha por linha
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		//Remove espaços vazios
		linha = strings.TrimSpace(linha)
		
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}
	arquivo.Close()

	return sites
}

func registraLog(site string, status bool){
	//Para o Go criar um arquivo, podemos utilizar a função OpenFile, também do pacote os. 
	//Ela recebe o nome do arquivo, uma flag para representar o que fazer com o arquivo, e o seu tipo de permissão. 
	//Há diversas flags que podemos utilizar, elas podem ser vistar aqui.
    //No nosso caso, se o arquivo não existir, queremos criá-lo, então vamos utilizar a flag O_CREATE ou O_RDWR, 
	//para ler e escrever no arquivo. Além disso, vamos dar a permissão 0666 para ele:

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + 
    " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

// restante do código omitido

func imprimeLogs() {

    arquivo, err := os.ReadFile("log.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(string(arquivo))
}

//Exemplo de Slice em Go
//func exibeNomes() {
//nomes := []string{"Douglas", "Daniel", "Bernardo"}

//fmt.Println("O meu slice tem", len(nomes), "itens")
//fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

//Adiciona um elemento ao slice
//nomes = append(nomes, "Aparecida")

//Mostra quantos items tem no slice e sua capacidade
//fmt.Println("O meu slice tem", len(nomes), "itens")
//fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
//}
