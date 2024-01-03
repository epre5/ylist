package main

import (
    "flag"
    "fmt"
    "os"
    "log"
    "net/http"
)

func main() {
    // 定义命令行参数
    var dir string
    var port string
    flag.StringVar(&dir, "dir", ".", "定义静态文件目录，默认为当前目录")
    flag.StringVar(&port, "port", "8080", "定义监听端口，默认为8080")

    // 自定义帮助信息
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        flag.PrintDefaults()
    }

    // 解析命令行参数
    flag.Parse()

    // 设置静态文件服务
    fs := http.FileServer(http.Dir(dir))
    http.Handle("/", fs)

    // 启动服务器
    log.Printf("Listening on :%s...\n", port)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatal(err)
    }
}
