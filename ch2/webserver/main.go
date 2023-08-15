package main

import (
	"log"
	"net"
	"syscall"
)

const (
	host    = "127.0.0.1"
	port    = 8888
	message = "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"Content-Length: 25\r\n" +
		"\r\n" +
		"Server with syscall"
)

func main() {
	fd, err := startServer(host, port)
	if err != nil {
		log.Fatal("error (startServer) : ",  err)
	}
	for {
		// accept
		cSock, cAddr, err := syscall.Accept(fd)
		if err != nil {
			log.Fatal("error (accept) : ",  err)
		}
		// create a goroutine to handle incoming client request
		go func(clientSocket int, clientAddress syscall.Sockaddr) {
			err := syscall.Sendmsg(clientSocket, []byte(message), []byte{}, clientAddress, 0)
			if err != nil {
				log.Fatal("error (send) : ", err)
			}
			syscall.Close(clientSocket)
		}(cSock, cAddr)
	}
}

func startServer(host string, port int) (int, error) {
	// open a socket
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatal("error (listen) : ", err)
	}
	srv := &syscall.SockaddrInet4{Port: port}
	addrs, err := net.LookupHost(host)
	if err != nil {
		log.Fatal("error (listen) : ", err)
	}
	for _, addr := range addrs {
		ip := net.ParseIP(addr).To4()
		copy(srv.Addr[:], ip)
		// bing to port
		if err = syscall.Bind(fd, srv); err != nil {
			log.Fatal("error (bind) : ", err)
		}
	}
	// listen
	if err = syscall.Listen(fd, syscall.SOMAXCONN);err!= nil {
		log.Fatal("error (listen) : ", err)
	} else {
		log.Println("Listening on ", host, ":", port)
	}
	return fd, nil;
}