package main

import (
	"bufio" //扫描io的包
	"fmt"
	"net"
	"strings"
	"time"
)
//声明一个类型，client，实际上就一个可以接收string数据类型的管道
//type client chan string //重命名
//发送私有消息数据格式:To|用户名|内容

type client struct {
	C    chan string
	Name string //用户的用户名
}

//登入
var login = make(chan client)

//广播消息
var message = make(chan string)

//退出
var logout = make(chan client)

var onlineClientMap map[client]bool

//通过用户名来获取对应的结构体 第三步
var onlineClientNameMap map[string]client

//监听三个管道
func Manager() {

	//要把客户端的信息存储到map中
	//onlineClientMap := map[client]bool{}

	onlineClientMap = make(map[client]bool) //给map空间

	onlineClientNameMap = make(map[string]client) //第四步

	for {
		select {
		case cli := <-login:
			//有一个新用户登入
			onlineClientMap[cli] = true //向map中添加一个cli

			///第五步
			onlineClientNameMap[cli.Name] = cli
		case cli := <-logout:
			//有一个用户退出
			delete(onlineClientMap, cli) //从map中删除一个cli

			delete(onlineClientNameMap, cli.Name) //第六步

		case msg := <-message:
			//有用户发送消息,要向每个在线的用户广播消息
			for cli, isOnline := range onlineClientMap {
				if isOnline == true {
					cli.C <- msg
				}
			}

		}
	}
}

func ReadFromMycli(conn net.Conn, myCli client) {
	//如果有广播的消息过来，将消息回显给客户端
	for msg := range myCli.C {
		conn.Write([]byte(msg + "\n"))
	}
}
//处理查询指令
func dealWhoCmd(conn net.Conn) {
	for cli := range onlineClientMap {
		fmt.Fprintln(conn, "name=", cli.Name)
	}

	fmt.Fprintln(conn, "目前在线人数:", len(onlineClientMap))

}

//处理发送私有消息的业务 第二步
func dealToCmd(name string, connMsg string) {

	targatClientName := strings.Split(connMsg, "|")[1] //Split是分割，去掉connMsg中的|
	fmt.Println("targaT=", targatClientName)

	//第七步
	onlineClientNameMap[targatClientName].C <- name + "对你说:[" + strings.Split(connMsg, "|")[2] + "]"

}
func handleCon(conn net.Conn) {
	defer conn.Close()

	///////添加的代码////////
	cliTimeOutTicker := time.NewTicker(time.Second * 60)
	//当前用户是否发送消息channel
	revcData := make(chan interface{})
	//对端是否关闭，用户不是超时关闭而是主动关闭
	removeClose := make(chan interface{})

	/////添加的代码////////

	//以客户端的ip加端口号为用户名
	myName := conn.RemoteAddr().String()

	//创建一个client,代表当前客户端和Message的channel通信
	//myCli := make(client)
	fmt.Println(myName)
	myCli := client{make(chan string), myName}

	//开个go程来接收消息
	go ReadFromMycli(conn, myCli)

	//登入动作
	login <- myCli //把myCil写入登入管道

	message <- myName + "已经登入"

	//第八步
	myCli.C <- "我是" + myName

	go func() {
		//阻塞，等待conn用户输入信息
		input_scanf := bufio.NewScanner(conn)

		for input_scanf.Scan() {
			connMsg := input_scanf.Text()
	
			/*
			switch {
			case connMsg == "who":
				fmt.Println("查询用户名和在线人数")
				fmt.Fprintln(conn, myName)
				//查询后要返回用户的用户名，用户名存储到map中
				//修改client结构体

				//用户查询指令业务
				dealWhoCmd(conn)

			default:
				fmt.Println(connMsg)
				myexp := regexp.MustCompile(`\d+\.\d+`)
				result := myexp.FindAllStringSubmatch(connMsg, 1)
				fmt.Println(result)
				fmt.Fprintln(conn, connMsg)
			}*/
			if connMsg == "who" {
				//需要查询当前在线人数和在线用户名
				fmt.Println("查询用户名和在线人数")
				fmt.Fprintln(conn, myName)
				//查询后要返回用户的用户名，用户名存储到map中
				//修改client结构体

				//用户查询指令业务
				dealWhoCmd(conn)
				////第一步, connMsg[:2]表示是字符串的前2个字符,字符串必须大于2个字符
			} else if len(connMsg) >= 2 && connMsg[:2] == "To" {
				fmt.Println("私发消息")
				dealToCmd(myName, connMsg)

			} else {
				//客户端已经发送数据过来，将这个数据进行广播
				message <- myName + ":" + connMsg
			}

			//用户正常输入了数据，更新定时器
			revcData <- struct{}{}
		}

		//用户正常退出
		removeClose <- struct{}{}
	}()

	for {
		select {
		case <-cliTimeOutTicker.C:
			//用户已经超时
			//退出
			logout <- myCli
			fmt.Fprintln(conn, "由于你长时间没有活跃，已断开连接")
			message <- myName + "已经超时，被退出"
			return
		case <-revcData:
			//用户正常数据
			cliTimeOutTicker.Stop()
			cliTimeOutTicker = time.NewTicker(time.Second * 60)
		case <-removeClose:
			//用户主动关闭
			//退出
			logout <- myCli
			message <- myName + "主动退出"
			return
		}
	}

}
func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ln.Close()

	go Manager() //监听管道

	for { //监听客户端
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		//开辟一个go程，去处理这个客户端的业务
		go handleCon(conn)
	}
}
