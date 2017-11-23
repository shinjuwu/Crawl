package main

import (
	"Crawl/forblog"
	"fmt"
	"runtime"
)

func main() {
	//forblogCrawl
	go forblog.Spy("http:/blog.csdn.net")
	for url := range forblog.UrlChannel {
		fmt.Println("routines num = ", runtime.NumGoroutine(), "chan len = ", len(forblog.UrlChannel)) //通過runtime可以獲取當前運行時的一些相關參數等
		go forblog.Spy(url)
	}
	fmt.Println("a")
}
