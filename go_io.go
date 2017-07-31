package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
$ go run go_io.go
test.txt 203
在gevent中用到的主要模式是Greenlet, 它是以C扩展模块形式接入Python的轻量级协程。 Greenlet全部运行在主程序操作系统进程的内部，但它们被协作式地调度。
*/
func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}
	// 输出文件的名字和大小
	fmt.Println(stat.Name(), stat.Size())

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	fmt.Println(string(bs))

	// short way to read file
	//read_file()

	// 创建文件
	//create_file("test1.txt")

	// 读取目录
	read_directory()

	//work path
	work_filepath()
}

func read_file() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}

	fmt.Println(string(bs))
}

func create_file(file_name string) {
	file, err := os.Create(file_name)
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString("test string")
}

// 读取目录
func read_directory() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	file_infos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range file_infos {
		fmt.Println(fi.Name())
	}
}

func work_filepath() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(info)
		return nil
	})
}
