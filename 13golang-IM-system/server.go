package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	//在线用户的列表 key 是用户名  v是用户对象
	OnlineMap map[string]*User
	mapLock   sync.RWMutex //同步锁,用于协程同步操作

	//消息广播的channel
	Message chan string
}

//创建一个server的端口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,                     //服务器ip
		Port:      port,                   //服务器监听端口
		OnlineMap: make(map[string]*User), //存放上线的用户
		Message:   make(chan string),      //广播消息的管道
	}
	return server
}

//处理建立好的连接
func (this *Server) Handeler(conn net.Conn) {
	//根据连接的消息创建一个用户登录用户的信息
	user := NewUser(conn, this)

	user.Online()

	//监听用户是否活跃的channel
	isLive := make(chan bool)

	//接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn Read err:", err)
				return
			}
			//提取用户的消息 去除换行符
			msg := string(buf[:n-1])

			//用户针对msg进行消息处理
			user.DoMessage(msg)

			//用户的任意消息，代表当前用户是一个活跃的
			isLive <- true
		}
	}()

	//当前handler 阻塞
	for {
		select {

		case <-isLive:
			//当前用户是活跃的，应该重置定时器
			// 不做任何事情，为了激活select，更新下面的定时器

		//time.After 表示定时器，会定时触发，进入这个case
		case <-time.After(time.Minute * 5):
			//已经超时10S，将当前的User强制关闭
			user.SendMsg("由于长时间没有发送消息，你已被强制下线\n")
			//销毁用户的资源
			close(user.C) //关闭管道
			//关闭连接
			conn.Close()
			//退出当前的handler
			return
		}
	}
}

//监听广播消息 Message通道，一旦有消息就发送给全部的在线User
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		//将msg发送给全部的在线user
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

//广播消息 向 Message通道 写入消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg //向 Server的 通道发送消息
}

//启动服务器
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))

	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}
	//close listen socket
	defer listener.Close()

	//启动监听Message的协程
	go this.ListenMessager()

	//循环等待建立连接
	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		//使用协程 do handler
		go this.Handeler(conn)
	}

}
