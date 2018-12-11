package action

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/radovskyb/watcher"

	"github.com/Chyroc/fs/internal/filesys"
)

func StartClient(host string, port int, pwd, dir string, hot bool, ignores []string) error {
	if !hot {
		return walkDirAndPushFiles(host, port, dir, ignores)
	}

	fmt.Println("hot mode.")

	lock := new(sync.Mutex)
	var fileMap = make(map[string]*filesys.FileContent)

	go func() {
		ticker := time.NewTicker(time.Millisecond * 200)
		for {
			select {
			case <-ticker.C:
				lock.Lock()
				if len(fileMap) == 0 {
					lock.Unlock()
					continue
				}
				fileMap = make(map[string]*filesys.FileContent)

				// todo 只同步 changed 文件
				if err := walkDirAndPushFiles(host, port, dir, ignores); err != nil {
					lock.Unlock()
					panic(err)
				}

				lock.Unlock()
			}
		}
	}()

	var handlerEvent = func(event watcher.Event) {
		lock.Lock()
		defer lock.Unlock()

		f := &filesys.FileContent{
			Name:  strings.TrimPrefix(event.Path, pwd+"/"),
			Mode:  event.Mode(),
			IsDir: event.IsDir(),
		}
		if filesys.IsIgnore(f, ignores) {
			return
		}
		fileMap[f.Name] = f
	}

	return watchDirWithCallback(dir, handlerEvent)
}

func watchDirWithCallback(dir string, f func(event watcher.Event)) error {
	w := watcher.New()
	w.FilterOps(watcher.Create, watcher.Write, watcher.Remove, watcher.Rename, watcher.Move, watcher.Chmod)

	go func() {
		for {
			select {
			case event := <-w.Event:
				f(event)
			case err := <-w.Error:
				panic(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.AddRecursive(dir); err != nil {
		return err
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		return err
	}

	return nil
}

func walkDirAndPushFiles(host string, port int, dir string, ignores []string) error {
	fmt.Printf("start sync: %s to %s", dir, fmt.Sprintf("%s:%d", host, port))
	defer fmt.Printf("\t\tdone.\n")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	defer conn.Close()

	l, err := filesys.Walk(dir, ignores)
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
