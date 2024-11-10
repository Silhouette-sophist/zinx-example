package client

import (
	"fmt"
	"log/slog"
	"net"
	"testing"
)

func TestName(t *testing.T) {
	conn, err := net.Dial("tcp", "0.0.0.0:8081")
	if err != nil {
		slog.Error("dial error: %v", err)
		return
	}
	for {
		_, err := conn.Write([]byte("client send"))
		if err != nil {
			slog.Error("write error: %v", err)
			return
		}
		bytes := make([]byte, 512)
		cnt, err := conn.Read(bytes)
		if err != nil {
			slog.Error("read error: %v", err)
			return
		}
		args := string(bytes[:cnt])
		fmt.Printf("read from server: %s\n", args)
	}
}
