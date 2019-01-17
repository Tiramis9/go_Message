package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"regexp"
	"strconv"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const (
	User     = "root"                          // 服务器的账号
	Password = "******"                        // 服务器的密码
	Host     = "127.0.0.1"                     // 服务器的地址
	Port     = "22"                            // 服务器的端口号
	FilePath = "D:/schedule/order/itemNow/src" // 本机 windows环境工作项目的路径
	DstPath  = "/home/go/src/"                 // 服务器 linx下工作项目的路径
)

var FileNameArray = []string{
	"game2",                // 项目程序1
	"golang_game_merchant", // 项目程序2
	"golang_game_admin",    // 项目程序3
	"game_agent",           // 项目程序4
}

var commandArray = []string{
	"cd /home/go/src/{file}",  //进入工作目录
	"tar -xzvf {file}.tar.gz", //解压文件
	"cd global",               //进入global目录
	"sed -i \"s/LOCAL = true/LOCAL = false/g\" const.go", //修改配置文件
}

func authenticate(host, username, password string) (ssh.Client, error) {
	// Create client config
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: time.Second * 5,
	}

	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", host+":"+Port, config)
	if err != nil {
		fmt.Println(err)
	}
	return *client, err
}

func uploadFile(client ssh.Client, filePath, dstPath string) error {
	var sftpClient *sftp.Client
	srcFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer srcFile.Close()
	if sftpClient, err = sftp.NewClient(&client); err != nil {
		fmt.Println(err)
		return err
	}
	defer sftpClient.Close()

	var remoteFileName = path.Base(filePath)
	dstFile, err := sftpClient.Create(path.Join(dstPath, remoteFileName))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer dstFile.Close()
	buf := make([]byte, 1024*1024*1024)
	for {
		n, err := srcFile.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("conn.Read error", err)
			} else {
				fmt.Println("conn.Read error", err)
			}
			break
		}
		dstFile.Write(buf[:n])
	}
	return nil
}

func runCommand(file, filePath, dstPath string) {
	client, err := authenticate(Host, User, Password)
	if err != nil {
		fmt.Printf("unable to connect: %v", err)
		return
	}
	defer client.Close()

	//<!--上传文件begin-->
	//上传文件太慢了
	err = uploadFile(client, filePath, dstPath)
	if err != nil {
		fmt.Printf("unable upload file: %v", err)
		return
	}
	//<!--上传文件end-->

	// Create a session
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Failed to create session: ", err)
		return
	}
	defer session.Close()

	stdIn, err := session.StdinPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	if err := session.Shell(); err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(commandArray); i++ {
		reg := regexp.MustCompile("{file}")
		command := reg.ReplaceAllString(commandArray[i], file)
		_, err := stdIn.Write([]byte(command + "\n"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	session.Wait()
}

func main() {
	var err error
	index := 0
	if len(os.Args) > 1 {
		index, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
	}
	if index < -1 || index > 3 {
		fmt.Println("index err")
	}
	if index == 1 { //当为商户管理后台时
		commandArray = append(commandArray, "sed -i \"s/ReportLocal = true/ReportLocal = false/g\" const.go")
	}
	file := FileNameArray[index] //game2
	//源文件路径
	filePath := FilePath + FileNameArray[index] + "/" + FileNameArray[index] + ".tar.gz" //D:/go_workspace/src/game2/game2.tar.gz
	//目标路径
	dstPath := DstPath + "" + FileNameArray[index] ///home/go/src/game2
	runCommand(file, filePath, dstPath)
}
