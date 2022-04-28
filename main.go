package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// 判断输入是否符合规范
	if len(os.Args) < 2 {
		log.Fatalln("Usage: tail -f <filename>")
	}

	fileName := flag.String("f", "", "<filename>")
	// 解析命令行参数
	flag.Parse()
	// 读取文件
	go Read(*fileName)
	select {}
}

// 持续写入文件
func Write() {
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Panicln(err)
		return
	}
	count := 1
	for {

		_, err = file.WriteString("我是第" + strconv.Itoa(count) + "行\n")
		if err != nil {
			log.Panicln(err)
			return
		}
		count++
		time.Sleep(time.Second)
	}
}

// 持续读取文件
func Read(fileName string) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Panicln("读取文件失败", err)
		return
	}
	reader := bufio.NewReader(file)
	// 设置行号
	LineNumber := 1
	for {

		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Panicln(err)
			return
		}
		if line == "" {
			time.Sleep(time.Second)
			continue
		}
		if line != "" {
			fmt.Print(LineNumber, "  ", line)
		}
		LineNumber++
		time.Sleep(time.Microsecond * 500)
	}
}
