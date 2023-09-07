package main

import (
	"go_blog/common"
	"go_blog/router"
	"log"
	"net/http"
)

func init() {
	//模板加载
	common.LoadTemplate()
}
func main() {
	//程序入口，一个项目 只能有一个入口
	//web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
