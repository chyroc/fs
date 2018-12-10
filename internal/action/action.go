package action

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/Chyroc/fs/internal/filesys"
)

func StartFolderSync(mode string, port int) error {
	if mode == "push" {
		return fmt.Errorf("push serber is deving")
	}

	return startTCPServer(port)
}

func startTCPServer(port int) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}

		go func() {
			if err := connHandler(conn); err != nil {
				fmt.Println("err:", err)
			}
		}()
	}

	return nil
}

func connHandler(conn net.Conn) error {
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		return err
	}

	var files []*filesys.FileContent
	if err := json.Unmarshal(bs, &files); err != nil {
		return err
	} else if len(files) == 0 {
		return nil
	}

	fmt.Println("sync:", files[0].Name)

	if err = os.RemoveAll(files[0].Name); err != nil {
		return err
	}

	return filesys.CreateFolder(files)
}
