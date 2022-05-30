package gobj

import (
	"bufio"
	"os"
)

type Obj struct {
	name string
}

func LoadObj(path string) (*Obj, error) {

	f, err := os.OpenFile(path, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}
	return decodeFile(f)
}

func decodeFile(file *os.File) (*Obj, error) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		print(line)
	}
	return nil, nil
}
