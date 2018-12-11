package filesys

import (
	"fmt"
	"github.com/mattn/go-zglob"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FileContent struct {
	Name    string
	Mode    os.FileMode
	IsDir   bool
	Content []byte
}

func (r *FileContent) String() string {
	return fmt.Sprintf("[%s](%v) %v", r.Name, r.IsDir, r.Mode)
}

func Walk(dir string, ignores []string) ([]*FileContent, error) {
	if dir != "./" {
		dir = strings.TrimPrefix(dir, "./")
	}

	var list []*FileContent
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		f := &FileContent{
			Name:  path,
			Mode:  info.Mode(),
			IsDir: info.IsDir(),
		}

		if IsIgnore(f, ignores) {
			return nil
		}

		if !info.IsDir() {
			bs, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			f.Content = bs
		}

		list = append(list, f)

		return nil
	})

	return list, nil
}

func CreateFolder(l []*FileContent) error {
	for _, v := range l {
		if err := CreateFileDir(v); err != nil {
			return err
		}
	}

	return nil
}

func CreateFileDir(file *FileContent) error {
	if file.IsDir {
		return os.MkdirAll(file.Name, file.Mode)
	}

	f, err := os.Create(file.Name)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := f.Chmod(file.Mode); err != nil {
		return err
	}

	if _, err := f.Write(file.Content); err != nil {
		return err
	}

	return nil
}

func GetDirectPath(dir string) (string, error) {
	if strings.HasPrefix(dir, "/") {
		return dir, nil
	}

	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(pwd, dir), nil
}

func IsIgnore(f *FileContent, ignores []string) bool {
	path := f.Name
	if f.IsDir && !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	for _, v := range ignores {
		if matched, err := zglob.Match(v, path); err == nil && matched {
			return true
		}
	}

	return false
}

func RemoveAll(f *FileContent) error {
	path := f.Name
	if f.IsDir {
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}

		files, err := filepath.Glob(path + "*")
		if err != nil {
			return err
		}
		for _, v := range files {
			if v == "./" {
				continue
			}
			if err := os.RemoveAll(v); err != nil {
				return err
			}
		}
	} else {
		if err := os.RemoveAll(path); err != nil {
			return err
		}
	}

	return nil
}
