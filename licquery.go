package main

import (
	"fmt"
	"os"
)

func main(){
	if query("他人修改源码后，是否可以闭源？") {
		if query("新增代码是否采用同样许可证？") {
			fmt.Println("GPL")
		}else{
			if query("是否需要对源码的修改之处提供说明文档？") {
				fmt.Println("Mozilla")
			}else{
				fmt.Println("LGPL")
			}
		}
	}else{
		if query("每一个修改过的文件，是否都必须放置版权说明？") {
			fmt.Println("Apache")
		}else{
			if query("衍生软件的广告，是否可以用你的名字促销？") {
				fmt.Println("MIT")
			}else{
				fmt.Println("BSD")
			}
		}
	}
}

func query(q string) bool {
	var in string
	fmt.Println(q)
	print("[y/n] ")
	fmt.Scan(&in)
	println("")
	switch in {
		case "y":
			return true
		case "n":
			return false
		default:
			fmt.Println("licquery exit!")
			os.Exit(0)
	}
	return false
}
