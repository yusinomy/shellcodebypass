package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const MAX_THREADS = 100

var lock sync.Mutex

func main() {
	var host, port string
	flag.StringVar(&host, "h", "172.18.1.1", "指定 ip 地址")
	flag.StringVar(&port, "p", "", "指定端口或端口范围（格式：80、80-100）")
	flag.Parse()
	if len(os.Args) != 5 {
		fmt.Println("go run test.go -h 172.18.1.1 -p 80")
		os.Exit(0)
	}
	var ports []int
	if port != "" {
		portRangeArr := strings.Split(port, "-")
		if len(portRangeArr) == 1 {
			// 指定特定的端口
			p, err := strconv.Atoi(port)
			if err != nil {
				fmt.Println("请输入正确的端口或端口范围")
				return
			}
			ports = append(ports, p)
		} else if len(portRangeArr) == 2 {
			// 指定端口范围
			startPort, err := strconv.Atoi(portRangeArr[0])
			if err != nil {
				fmt.Println("请输入正确的端口范围")
				return
			}
			endPort, err := strconv.Atoi(portRangeArr[1])
			if err != nil {
				fmt.Println("请输入正确的端口范围")
				return
			}
			for i := startPort; i <= endPort; i++ {
				ports = append(ports, i)
			}
		}
	}

	// 扫描端口
	sem := make(chan int, MAX_THREADS)
	for _, p := range ports {
		sem <- 1
		go func(port int) {
			con, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), time.Duration(50)*time.Millisecond)
			if err != nil {
				<-sem
				return
			}
			con.Close()
			lock.Lock()
			fmt.Println(port, "is alive")
			lock.Unlock()
			<-sem
		}(p)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- 1
	}
}
