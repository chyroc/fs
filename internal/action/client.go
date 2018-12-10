package action

import (
	"encoding/json"
	"fmt"
	"github.com/Chyroc/fs/internal/filesys"
	"net"
)

func StartClient(host string, port int, dir string) error {
	return connTCPServer(host, port, dir)
}

func connTCPServer(host string, port int, dir string) error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}

	return clientConnHandler(conn, dir)
}

func clientConnHandler(conn net.Conn, dir string) error {
	fmt.Println(dir)

	l, err := filesys.Walk(dir)
	if err != nil {
		return err
	}

	bs, err := json.Marshal(l)
	if err != nil {
		return err
	}

	_, err = conn.Write(bs)
	return err
}
