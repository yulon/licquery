package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) > 1 {
		os.Chdir(os.Args[1])
	}

	if query("他人修改源码后，是否可以闭源？") {
		if query("新增代码是否采用同样许可证？") {
			output("GPL", gpl)
		}else{
			if query("是否需要对源码的修改之处提供说明文档？") {
				output("Mozilla", mozilla)
			}else{
				output("LGPL", lgpl)
			}
		}
	}else{
		if query("每一个修改过的文件，是否都必须放置版权说明？") {
			output("Apache", apache)
		}else{
			if query("衍生软件的广告，是否可以用你的名字促销？") {
				output("MIT", mit)
			}else{
				output("BSD", bsd)
			}
		}
	}
}

func query(q string) bool {
	var in string
	fmt.Println(q + "[y/n]")
	fmt.Scan(&in)
	switch in {
		case "y":
			return true
		case "n":
			return false
		default:
			fmt.Println("Exit!")
			os.Exit(0)
	}
	return false
}

func output(name string, cont string) {
	dir, _ := os.Getwd()
	fmt.Println("Writing " + name + " to " + dir + " ...")
	f, err := os.Create("LICENSE")
	if err != nil {
		fmt.Println("Error!")
		return
	}
	f.Write([]byte(cont))
	f.Close()
	fmt.Println("Done!")
}
