package socket

import (
	_ "../model"
	"runtime"
)

var os int

func socketBind() {
	currentOS := runtime.GOOS
	if currentOS == "windows" {
		os = 1
	} else if currentOS == "darwin" || currentOS == "freebsd" || currentOS == "linux" {
		os = 2
	} else {
		os = 0
	}

}

// handleWindowsSocket
// WinSock is using send,recv,WSASend,WSARecv to handle TCPStream
// Using sendto,recvfrom,WSASendTo,WSARecvFrom to handle UDPStream
// Using WSASend function to send more buf to improve experience than calling "Send".
func handleWindowsSocket() {
	// PDUHeader was UDP-Like.

}
