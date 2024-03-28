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

	return string(body), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Cliente: ")
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return
	}

	message = strings.TrimSpace(message)

	response, err := sendMessage(message)
	if err != nil {
		fmt.Println("Erro ao enviar mensagem:", err)
		return
	}

	fmt.Println("Servidor:", response)
}
