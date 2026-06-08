package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func scanport(target string, port int) (bool, string) {
	address := fmt.Sprintf("%s:%d", target, port)
	conn, err := net.DialTimeout(
		"tcp",
		address,
		1*time.Second,
	)

	if err != nil {
		return false, ""
	}

	buffer := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	n, _ := conn.Read(buffer)

	banner := string(buffer[:n])

	conn.Close()

	service := "unknow"

	if strings.Contains(banner, "SSH") {
		service = "SSH"
	} else if strings.Contains(banner, "HTTP") {
		service = "HTTP"
	}
	return true, service + "|" + banner

}
