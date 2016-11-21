package main

import (
	"errors"
	"fmt"
	"net"
	"os"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on tcp", ":8081")
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		defer conn.Close()
		// Handle connections in a new goroutine.
		handle := handleRequest(conn)
		if handle != nil {
			fmt.Println("operation failed")
		}
	}

}

func Check(response string) string {
	if len(response) != 0 {
		if (response == "Hello") || (response == "hello") {
			response = "Hi"
		} else if (response == "Name") || (response == "name") || (response == "Name??") {
			response = "Chitty :) "
		} else if (response == "Ok!! Goodbye") || (response == "Bye") || (response == "bye") {
			response = "Bye"
		} else {
			response = "What?"
		}

	} else {
		response = "Sorry??"
	}
	return response
}

// Handles incoming requests.
func handleRequest(conn net.Conn) error {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	for {

		err := Request(conn, buf)
		if err != nil {
			//fmt.Println("Error reading:", err.Error())
			return errors.New("ERROR")
		}

	}

	return nil
}

func Request(conn net.Conn, buf []byte) error {

	reqLen, err := conn.Read(buf)
	if err != nil {
		//fmt.Println("Error reading:", err.Error())
		return errors.New("Received empty")
	}
	// output message received
	request := string(buf[:reqLen-2])
	fmt.Println("Client Request	:", request)

	//reading response from server
	response := Check(request)
	fmt.Println("Server Response : ", response)

	// Send a response back to person contacting us.
	_, err1 := conn.Write([]byte(response + "\n"))
	//checking if there is an error through a message
	if err1 != nil {
		return errors.New("Connection Full")
	}
	return nil
}
