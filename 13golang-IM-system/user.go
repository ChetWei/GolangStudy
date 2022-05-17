package main

import (
	"net"
)

type User struct {
	Name string
	Addr string
	C    chan string //用户的消息通道
	conn net.Conn

	server *Server //继承 Server
}

//根据连接 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string), //消息通道
		conn:   conn,              //连接
		server: server,
	}
	//启动监听当前 user channel消息的goroutine
	go user.ListenMessage()

	return user
}

//用户的上线业务
func (this *User) Online() {

	//用户上线了,将用户加入到online map中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this //将上线的用户添加到map中
	this.server.mapLock.Unlock()

	//广播当前用户上线的消息
	this.server.BroadCast(this, "已上线")
}

//用户的下线业务
func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name) //将上线的用户添加到map中
	this.server.mapLock.Unlock()

	//广播当前用户上线的消息
	this.server.BroadCast(this, "下线")
}

//给当前user对应的客户端发送消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

//用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询当前在线用户都有哪些
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else {
		//广播消息
		this.server.BroadCast(this, msg)
	}

}

//监听当前User channel 的方法，一旦有消息就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C                     //从当前用户的通道中取消息
		this.conn.Write([]byte(msg + "\n")) //向连接端发送消息
	}
}
