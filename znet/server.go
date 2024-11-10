package znet

import (
	"fmt"
	"log/slog"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

// Start 启动
func (s *Server) Start() {
	slog.Info("[Start] Server listener at IP: %s, Port: %d, is starting", s.IP, s.Port)

	go func() {
		// 1.获取tcp监听地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			slog.Error("resolve tcp addr error: %v", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			slog.Error("listen %v error: %v", addr, err)
			return
		}
		slog.Info("start server: %v", s)
		// 2.监听所有连接，并处理
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				slog.Error("accept err: %v", err)
				continue
			}
			// 3.创建连接后启动协程处理
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						slog.Error("conn read error: %v", err)
						continue
					}
					if _, err := conn.Write(buf[:cnt]); err != nil {
						slog.Error("conn write error: %v", err)
						continue
					}
				}
			}()
		}
	}()
}

// Stop 停止服务端
func (s *Server) Stop() {
	slog.Warn("server stop: %v", s)
}

func (s *Server) Serve() {
	s.Start()

	select {}
}
