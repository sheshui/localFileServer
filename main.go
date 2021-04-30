package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {

	// 获取当前目录
	path, _ := os.Getwd()
	fmt.Println("运行路径:", path)

	// 获取当前指定端口 未指定默认8080
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "8080"
	}
	fmt.Println("运行端口：", port)

	fmt.Println("本机IPV4地址:")
	// 获取网卡接口
	interf, _ := net.Interfaces()
	for _, value := range interf {
		// 判断接口是否有up的flag (是否启用？)
		if flagString := value.Flags.String(); strings.Contains(flagString, "up") {
			// 获取接口的地址
			addrs, _ := value.Addrs()
			for _, addrValue := range addrs {
				// 遍历并去除不可用的ip，去除回环地址，去除没有ipv4的地址
				if ipnet, ok := addrValue.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						fmt.Println(ipnet.IP.String())
					}
				}
			}
		}
		//fmt.Println(value)
		//flagString := value.Flags.String()
		//fmt.Println(flagString)
		//fmt.Println(strings.Contains(flagString, "up"))
		//fmt.Println()
	}

	//addr, err := net.InterfaceAddrs()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, value := range addr {
	//	if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	//		if ipnet.IP.To4() != nil {
	//			fmt.Println(ipnet.IP.String())
	//		}
	//	}
	//}

	// 启动文件服务
	http.Handle("/", http.FileServer(http.Dir('.')))
	_ = http.ListenAndServe(":"+port, nil)
}
