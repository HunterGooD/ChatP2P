// Package network для создания узлов передачи данных между ними
package network

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

// DisconnectBytes строка отправляющая сигнал об отключении пользователя
const DisconnectBytes = "\001\002\003\004\003\002\001"

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

// ConnectTo Присоединяет пользователей к какому либо пользователю
func (n *Node) ConnectTo(addresses []string) {
	// TODO: Генерация сеансового ключа и передача
	for _, addr := range addresses {
		splitAddr := strings.Split(addr, ":")
		if _, ok := n.Connections[":"+splitAddr[1]]; ok {
			delete(n.Connections, ":"+splitAddr[1])
		}
		n.Connections[addr] = true
	}
}

// Disconnect отключает соединения
func (n *Node) Disconnect(addresses []string, ok bool) {
	var p = &Package{
		From: n.MyAddress(),
		Data: DisconnectBytes,
	}
	for _, addr := range addresses {
		if ok {
			p.To = addr
			n.Send(p)
		}
		delete(n.Connections, addr)
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
