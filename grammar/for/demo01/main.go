package main

import (
	"fmt"
	"time"
)

func main() {
	go test()
	time.Sleep(10 * time.Minute)
}

func test() {
	timeout := time.After(5 * time.Second)
	// 每 10 秒获取一次应用状态
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	defer ticker.Stop()
ForEnd:
	for {
		select {
		case <-ticker.C:
			i += 1
			if i == 3 {
				break ForEnd
			}
			fmt.Println("每秒")
		case <-timeout:
			fmt.Println("结束")
			break ForEnd
		}
	}
	fmt.Println("111")
}
