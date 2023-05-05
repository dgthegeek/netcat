package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	user, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Creating two chanell one for the message and one for the evantual errors
	messages := make(chan string)
	errors := make(chan error)

	// goroutine to read messages from the server
	// go func(){
	// 	for {
	// 		message, err := bufio.NewReader(user).ReadString('\n')
	// 		if err != nil{
	// 			errors <- err
	// 		}
	// 		messages <- message
	// 	}
	// }()

	//goroutine to read read input from the user and send it to the server
	go func(){
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(">")
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(user, text+"\n")
			//message, _ := bufio.NewReader(user).ReadString('\n')
			
		}
	}()

	// Listen for messages or errors
	for {
		select {
		case message := <- messages:
			fmt.Println("Message from the server:" + message)
		case err :=  <- errors:
			fmt.Println("error", err)
			os.Exit(1)
		}
	}

	
}   