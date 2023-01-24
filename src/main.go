package main

import(
	"os"
	"fmt"
	"net"
	"strings"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9898"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("Monopoly-game-service has started...")
	server, err := net.Listen(SERVER_TYPE, strings.Join([]string{SERVER_HOST, SERVER_PORT}, ":"))

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Monopoly-game-service is ready for user connections...")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			os.Exit(1)
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	buffer := make([]byte, 1024)

	for {
		bufferLen, err := connection.Read(buffer)

		if err != nil {
			fmt.Println("Read error:", err.Error())
			break
		}

		fmt.Println("Received: ", string(buffer[:bufferLen]))
		_, err = connection.Write([]byte("Thanks! Got your message:"))

		if err != nil {
			fmt.Println("Write error:", err.Error())
			break
		}
	}
}
