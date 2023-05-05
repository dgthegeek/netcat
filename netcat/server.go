package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Listening to the port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Something went wrong : ", err)
		return
	}
	defer listener.Close()
	fmt.Println("Listenning to port :8080")

	// boocle to accept the connections
	for i := 0; i < 11; i++ {
		user, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your name: ")
		name, _ := reader.ReadString('\n')

		fmt.Println("New connection from : ", name)

		//Treat the connection in a go roution
		go handleConnection(user, name)
	}
}

func handleConnection(user net.Conn, name string) {
	// cut the connection the the function not runing
	defer user.Close()

	// Let the client know that he is connected successfully

	// Retrieve and read the data sent by the client
	buffer := make([]byte, 1024)
	for {
		n, err := user.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		message := string(buffer[:n])
		fmt.Printf("Message received from %s: %s", user.RemoteAddr(), message)
	}
	//DIffuser le message dans les autre client connectes
}
