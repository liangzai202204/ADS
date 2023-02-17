package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 获取文件名
	list := os.Args
	log.Print(list)
	if len(list) != 5 {
		fmt.Println("输入有误")
		return
	}
	log.Print("命令：", list[1], ",地址：", list[2], ",命令：", list[3], ",文件：", list[4])

	log.Print(list)
}
