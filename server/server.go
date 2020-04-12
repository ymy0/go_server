package main

import (
	//	"database/sql"
	"fmt"
	"log"

	//	"os"
	//"strconv"

	"../common"
	//_ "github.com/go-sql-driver/mysql"
)

func main() {
	/* log日志初始化 */
	logpath := "../log/" + "go_test_server"
	var logfile *(common.Logger) = new(common.Logger)
	err := logfile.Init(logpath)
	if err != nil {
		fmt.Println(err)
		return
	}
 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
 	log.SetOutput(logfile)
	log.Println("启动http服务器")
	var http_handler Http
	http_handler.StartHttp()
}

