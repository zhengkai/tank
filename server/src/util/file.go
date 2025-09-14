package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"project/config"
	"project/zj"
	"strings"

	"github.com/zhengkai/zu"
	"google.golang.org/protobuf/proto"
)

type DownloadFunc func(url string) (ab []byte, err error)

func Mkdir(filename string) error {
	dir := filepath.Dir(Static(filename))
	if dir == config.StaticDir {
		return nil
	}
	return os.MkdirAll(dir, 0755)
}

func FileDelete(filename string) error {
	return os.Remove(Static(filename))
}

// FileExists ...
func FileExists(filename string) bool {
	return zu.FileExists(Static(filename))
}

// IsURL ...
func IsURL(s string) bool {
	return strings.HasPrefix(s, `https://`) || strings.HasPrefix(s, `http://`)
}

// ReadFile ...
func ReadFile(file string) (ab []byte, err error) {
	ab, err = os.ReadFile(Static(file))
	return
}

// Static ...
func Static(file string) string {
	file = strings.TrimPrefix(file, config.StaticDir+`/`)
	return fmt.Sprintf(`%s/%s`, config.StaticDir, file)
}

// SaveData ...
func SaveData(name string, p proto.Message) (err error) {

	defer zj.Watch(&err)

	ab, err := proto.Marshal(p)
	if err == nil {
		WriteFile(name+`.pb`, ab)
	}

	ab, err = json.Marshal(p)
	if err == nil {
		WriteFile(name+`.json`, ab)
	}

	return
}

// WriteFile ...
func WriteFile[T []byte | string](file string, content T) (err error) {
	file = Static(file)
	Mkdir(file)
	return writeFile(file, content)
}

func TouchFile(file string) error {
	file = Static(file)
	Mkdir(file)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	return f.Close()
}

func writeFile[T []byte | string](file string, content T) (err error) {

	tmpName, err := WriteTemp(file, content)
	if err != nil {
		return
	}

	err = os.Chmod(tmpName, 0644)
	if err != nil {
		return
	}

	// 这个函数存在的意义在这一步，改名是个原子操作，避免程序异常导致文件写不完整
	err = os.Rename(tmpName, file)
	if err != nil {
		os.Remove(tmpName)
		return
	}
	return
}

func WriteTemp[T []byte | string](file string, content T) (filename string, err error) {

	f, err := os.CreateTemp(Static(`tmp`), zu.TempFilePattern)
	if err != nil {
		return
	}

	switch v := any(content).(type) {
	case string:
		_, err = f.WriteString(v)
	case []byte:
		_, err = f.Write(v)
	}
	f.Close()
	if err != nil {
		return
	}
	return f.Name(), nil
}
