package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/huntergood/ChatP2P/pkg/network"
	"github.com/huntergood/chatP2P/pkg/input"
)

func init() {
	// Проверка указан ли IP:pport для запуска
	os.Exit(0)
	if len(os.Args) != 2 {
		panic("Incorect cmd args run")
	}
}

func main() {
	// пример простого запуска
	network.NewNode(os.Args[1]).Run(handleServer, handleClient)
}

func handleConnection(node *network.Node, conn net.Conn) {
	defer conn.Close()
	var (
		buffer     = make([]byte, 512)
		resPackage = new(network.Package)
		message    string
	)
	for {
		length, err := conn.Read(buffer)
		if err != nil {
			break
		}
		message += string(buffer[:length])
	}
	if err := json.Unmarshal([]byte(message), &resPackage); err != nil {
		return
	}
	node.ConnectTo([]string{resPackage.From})
	fmt.Println(resPackage.Data)
}

func handleServer(node *network.Node) {
	listen, err := net.Listen("tcp", node.MyAddress())
	if err != nil {
		panic("Sorry, listen close")
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			break
		}
		go handleConnection(node, conn)
	}
}

func handleClient(node *network.Node) {
	var message string
	for {
		message = input.IUser()
		splitMessage := strings.Split(message, " ")

		switch splitMessage[0] {

		case "/exit":
			os.Exit(0)

		case "/connect":
			node.ConnectTo(splitMessage[1:])

		case "/network":
			node.GetNetwork()

		default:
			node.SendMessageAll(message)
		}
	}
}
