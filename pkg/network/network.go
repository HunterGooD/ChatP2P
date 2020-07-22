package network

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

// Node /
type Node struct {
	Connections map[string]bool
	Address     Address
}

// Address /
// ip v4
type Address struct {
	IP   string
	Port string
}

// Package передача данных
type Package struct {
	To   string
	From string
	Data string
}

// NewNode возвращает созданый узел
// address ip:port
func NewNode(address string) *Node {
	splitAddr := strings.Split(address, ":")

	// Если не передан порт возвращаем nil
	if len(splitAddr) != 2 {
		return nil
	}
	return &Node{
		Connections: make(map[string]bool),
		Address: Address{
			IP:   splitAddr[0],
			Port: ":" + splitAddr[1],
		},
	}
}

// Run стартует функции пользователя как клиента и сервера
func (n *Node) Run(handleServer func(*Node), handleClient func(*Node)) {
	go handleServer(n)
	handleClient(n)
}

// ConnectTo Присоединяет пользователя к какому либо пользователю
func (n *Node) ConnectTo(addresses []string) {
	// TODO: Генерация сеансового ключа и передача
	for _, addr := range addresses {
		n.Connections[addr] = true
	}
}

// MyAddress getter address
func (n *Node) MyAddress() string {
	return n.Address.IP + n.Address.Port
}

// GetNetwork getter network
func (n *Node) GetNetwork() {
	for addr := range n.Connections {
		fmt.Println("|", addr)
	}
}

// SendMessageAll Отправляет сообщение всем пользователям
func (n *Node) SendMessageAll(message string) {
	var newPackage *Package = &Package{
		From: n.MyAddress(),
		Data: message,
	}
	for addr := range n.Connections {
		newPackage.To = addr
		n.Send(newPackage)
	}
}

// Send отправка сообщения
func (n *Node) Send(pack *Package) {
	conn, err := net.Dial("tcp", pack.To)
	if err != nil {
		delete(n.Connections, pack.To)
		return
	}
	defer conn.Close()
	jsonPack, _ := json.Marshal(&pack)
	conn.Write(jsonPack)
}
