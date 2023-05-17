package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readByBytes(file *os.File) {
	data := make([]byte, 9)
	file.Read(data)
	fmt.Println(string(data))
}

func readByLines(file *os.File) {
	reader := bufio.NewReader(file)

	for true {
		lineContent, err := reader.ReadString('\n')
		fmt.Print(lineContent)
		//lineContent, _, _ := reader.ReadLine()
		//fmt.Println(string(lineContent))
		if err == io.EOF {
			break
		}
	}

}

func readByFile(file *os.File) {
	content, _ := ioutil.ReadFile("满江红")
	fmt.Println(string(content))
}

func main() {
	file, err := os.Open("满江红")
	if err != nil {
		fmt.Println("打开文件失败")
	}
	fmt.Println("file:", file)

	// 按字节读取
	//readByBytes(file)

	// 按行都
	//readByLines(file)

	// 读取整个文件
	//readByFile(file)
}
