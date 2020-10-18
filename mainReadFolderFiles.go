package main

import (
	"data_struct/stackarray"
	"errors"
	"fmt"
	"io/ioutil"
)

// GetAllFilesByStack 使用stack的方式来获取所有文件
// stack 作为一个临时缓冲区，用来存放读取到的所有文件夹路径
func GetAllFilesByStack(path string, files []string) ([]string, error) {
	files = append(files, path)
	stack := stackarray.NewStack()
	stack.Push(path)

	for !stack.IsEmpty() {
		path, _ := stack.Pop()
		// 读取文件夹下的所有路径
		fileInfo, err := ioutil.ReadDir(path.(string))
		if err != nil {
			return files, errors.New("Fail to read the file path")
		}
		for _, file := range fileInfo {
			subFilePath := path.(string) + "/" + file.Name()
			files = append(files, subFilePath)

			if file.IsDir() {
				// 将所有文件夹路径存入stack中
				stack.Push(subFilePath)
			}
		}
	}

	return files, nil
}

// GetAllFiles 递归的读取当前路径及其子路径下所有的文件夹和文件名
func GetAllFiles(path string, files []string) ([]string, error) {
	files = append(files, path)
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("Fail to read the file path")
	}
	for _, file := range fileInfo {
		subFilePath := path + "/" + file.Name()
		files = append(files, subFilePath)
		if file.IsDir() {
			// 递归读取所有子文件
			files, _ = GetAllFiles(subFilePath, files)
		}
	}

	return files, nil
}

func main123() {
	path := "/Users/xinc2/workspace/go/src"
	files := []string{}
	// files, _ = GetAllFiles(path, files)
	files, _ = GetAllFilesByStack(path, files)
	for _, val := range files {
		fmt.Println(val)
	}
}
