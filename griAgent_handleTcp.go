//go:build windows
// +build windows

package griagent

import (
	"bufio"
	"io"
	"log"
	"strings"
)

// this will handle tcp connection
// between service and the server
func (s *GriAgent) handleTcpReqests() {
	clientReader := bufio.NewReader(s.Conn)
	for {

		clientRequest, err := clientReader.ReadString('\n')

		switch err {
		case nil:
			clientRequest := strings.TrimSpace(clientRequest)
			log.Println("SrvResp: ", string(clientRequest))
		case io.EOF:
			s.Stop()
			log.Println("client closed the connection")
			return
		default:
			s.Stop()
			log.Printf("client error: %v\n", err)
			return
		}
	}
}
