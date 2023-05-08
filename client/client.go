package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	port := ""
	args := os.Args[1:]
	if len(args) == 1 {
		port = args[0]
	} else {
		fmt.Println("[USAGE]:go run . $port")
		os.Exit(0)
	}
	user, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Creating two chanell one for the message and one for the evantual errors
	messages := make(chan string)
	errors := make(chan error)

	// goroutine to read messages from the server
	go func() {
		for {
			message, err := bufio.NewReader(user).ReadString('\n')
			if err != nil {
				errors <- err
			}
			messages <- message
		}
	}()

	// goroutine to read input from the user and send it to the server
	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(">>")
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(user, text)
		}
	}()

	// Listen for messages or errors
	for {
		select {
		case message := <-messages:
			fmt.Print(message)
		case err := <-errors:
			fmt.Print("error", err)
			os.Exit(1)
		}
	}

}
