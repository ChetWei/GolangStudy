package main

import (
	"flag"
	"fmt"
	"net"
)

//客户端

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn //套接字句柄
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
	}

	//连接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	//将连接句柄复制给对象
	client.conn = conn
	//返回对象
	return client
}

var serverIp string
var serrverPort int

//初始化
func init() {
	//绑定命令行参数
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器ip地址(默认是127.0.0.1)")
	flag.IntVar(&serrverPort, "port", 8888, "设置服务器的端口(默认是8888)")

}

func main() {

	//解析命令行  控制台输入   ./client -ip 127.0.0.1 -port 8888
	flag.Parse()

	client := NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println(">>>> 连接服务器失败...")
		return
	}

	fmt.Println(">>>>>> 连接服务器成功")

	//启动客户端业务
	select {}
}
