package filesys

import (
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

func Walk(dir string) ([]*FileContent, error) {
	dir=strings.TrimPrefix(dir,"./")

	var list []*FileContent
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		var content []byte
		if !info.IsDir() {
			bs, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			content = bs
		}

		list = append(list, &FileContent{
			Name:    path,
			Mode:    info.Mode(),
			IsDir:   info.IsDir(),
			Content: content,
		})

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
