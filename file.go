package gowrap

import (
	"fmt"
	"os"
	"path/filepath"
)

func Test() {

	fmt.Println("Test")
}

func CreateFile(path string) (*os.File, error) {
	p, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	fi, err := os.OpenFile(p, os.O_CREATE, 777)
	if os.IsNotExist(err) {
		err := os.Mkdir(filepath.Dir(p), 0777)
		if err != nil {
			return nil, err
		}
		fi, err = os.OpenFile(p, os.O_CREATE, 777)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return fi, nil
}

func CreatOrOpenFile(path string, flag int) (*os.File, error) {
	p, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	fi, err := os.OpenFile(p, flag, 777)

	if os.IsNotExist(err) {
		return CreateFile(p)
	}

	if err != nil {
		return nil, err
	}

	return fi, nil
}
