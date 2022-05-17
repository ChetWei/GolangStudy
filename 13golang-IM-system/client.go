package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

//客户端

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn //套接字句柄
	flag       int      //当前client的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
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

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名模式")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>请输入合法范围内的数字<<<")
		return false
	}

}

//公聊模式
func (client *Client) PublicChat() {
	var chatMsg string
	//提示用户输入消息
	fmt.Println(">>>>>请输入聊天内容，exit退出.")
	fmt.Scanln(&chatMsg)

	//一直循环知道用户输入exit退出
	for chatMsg != "exit" {
		//发给服务器
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg)) //发送消息
			if err != nil {
				fmt.Println("conn Write err:", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println(">>>>> 请输入聊天内容，exit退出.")
		fmt.Scanln(&chatMsg)
	}

}

//查询在线的用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write err:", err)
		return
	}
}

//私聊模式
func (client *Client) PrivateChat() {
	var remoteName string //选择的用户名字
	var chatMsg string    // 聊天的内容

	client.SelectUsers()
	fmt.Println(">>>请输入聊天的对象[用户名],exit退出:")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>请输入消息内容,exit退出:")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit" {
			//发给服务器
			if len(chatMsg) != 0 {
				//使用发送个用户的模式
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := client.conn.Write([]byte(sendMsg)) //发送消息
				if err != nil {
					fmt.Println("conn Write err:", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println(">>>>> 请输入聊天内容，exit退出.")
			fmt.Scanln(&chatMsg)
		}
	}

}

func (client *Client) UpdateName() bool {
	fmt.Println(">>> 请输入用户名:")
	fmt.Scanln(&client.Name)

	//构建重命名的语句
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

//接受服务器的消息，标准输出到控制台
func (client *Client) DealResponse() {
	//一旦clien.conn有数据，就直接copy到stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, client.conn)

}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
			fmt.Println("选择模式有误")
		}
		//根据用户选择的不同模式选择不同的业务
		switch client.flag {
		case 1:
			//公聊模式
			client.PublicChat()
			break
		case 2:
			//私聊模式
			client.PrivateChat()
			break
		case 3:
			//更新用户名
			client.UpdateName()
			break
		}

	}
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

	//开启一个goroutine处理服务器的返回消息
	go client.DealResponse()

	fmt.Println(">>>>>> 连接服务器成功")

	//启动客户端业务

	client.Run()
}
