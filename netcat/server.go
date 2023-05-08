package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	port := "8989"
	args := os.Args[1:]
	if len(args) == 0 {
		port = port
	} else if len(args) == 1 {
		port = args[0]
	} else {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(0)
	}
	// Listening to the port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Something went wrong : ", err)
		return
	}
	defer listener.Close()
	fmt.Println("Listenning to port :", port)

	// boocle to accept the connections
	for i := 0; i < 11; i++ {
		user, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Print out the ubuntu logo
		file, err := os.Open("ubuntu.txt")
		if err != nil {
			fmt.Println("Error reading the ubuntu text file", err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Fprintln(user, string(scanner.Text()))
		}

		// Send a welcome messgae to the client
		fmt.Fprintf(user, "[ENTER YOUR NAME]:")

		reader := bufio.NewReader(user)
		name, _ := reader.ReadString('\n')
		name = name[:len(name)-1]

		// Print th history odf the chat
		historychat, err := os.Open("chathistory.txt")
		if err != nil {
			fmt.Println("Error reading the ubuntu text historychat", err)
		}
		defer historychat.Close()
		scanner1 := bufio.NewScanner(historychat)
		for scanner1.Scan() {
			fmt.Fprintln(user, string(scanner1.Text()))
		}

		// now := time.Now()
		// formatedTime := now.Format("2006-01-02 15:04:05")
		// fmt.Fprint(user, "[" + formatedTime + "]" + "[" + name + "]:")


		//Treat the connection in a go roution
		go handleConnection(user, name, &connections)
	}
}

// List of clients
var connections []net.Conn

func handleConnection(user net.Conn, name string, connections *[]net.Conn) {

	// add user connection to list of connections
	*connections = append(*connections, user)

	// cut the connection the the function not runing
	defer user.Close()

	// Retrieve and read the data sent by the client
	buffer := make([]byte, 1024)
	for {
		n, err := user.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		message := string(buffer[:n])
		now := time.Now()
		formatedTime := now.Format("2006-01-02 15:04:05")

		// broadcast message to all clients
		for _, conn := range *connections {
			if conn != user{
				fmt.Fprintln(conn)
				fmt.Fprint(conn, "["+formatedTime+"]"+"["+name+"]:"+message)
			}
		}
		//fmt.Fprint(user, "[" + formatedTime + "]" + "[" + name + "]:")

		WriteHistory("[" + formatedTime + "]" + "[" + name + "]:" + message)
	}
}

func WriteHistory(data string) {
	file, err := os.OpenFile("chathistory.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Print(err)
	}
}
