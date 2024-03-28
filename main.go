package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Message struct {
	Question string `json:"question"`
}

type Response struct {
	Answer string `json:"answer"`
}

func sendMessage(message string) (string, error) {
	url := "http://localhost:8080/chatbot"
	requestBody, err := json.Marshal(Message{Question: message})
	if err != nil {
		return "", err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var jsonResponse Response
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return "", err
	}

	return jsonResponse.Answer, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Cliente: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler a entrada:", err)
			return
		}

		message = strings.TrimSpace(message)

		if strings.ToLower(message) == "sair" {
			fmt.Println("Saindo...")
			break
		}

		response, err := sendMessage(message)
		if err != nil {
			fmt.Println("Erro ao enviar mensagem:", err)
			return
		}

		fmt.Println("Servidor:", response)
	}
}
