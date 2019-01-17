package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"

	"github.com/pkg/sftp"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
)

const (
	User     = "root"
	Password = "*******"
	Host     = "127.0.0.1"
	Port     = "22"
)

//配置文件名数组
var FileNameArray = []string{
	"game2",
	"golang_game_merchant",
	"golang_game_admin",
	"game_agent",
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
	fmt.Println("upload success")
	return nil
}

func runCommand(commandArray []string, filePath, dstPath string) {
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
		command := commandArray[i]
		_, err := stdIn.Write([]byte(command + "\n"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	session.Wait()
}

func localCmdRun(localCmd []string) error {
	/*for i := range localCmd {
		status.ErrMerchantCodeExist
	}*/
	for i := range localCmd {
		cmd := exec.Command("cmd", "/C", localCmd[i])
		buf, err := cmd.Output()
		if err != nil {
			fmt.Println("cmd.Run", err)
			return err
		}
		fmt.Fprintf(os.Stdout, "Result: %s", buf)

	}
	//err := cmd.Run()
	fmt.Println("tar czvf success")
	return nil
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

	fmt.Println(FileNameArray[index])
	v := viper.New()
	v.SetConfigName(FileNameArray[index])
	v.AddConfigPath("config")
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//本地执行命令组
	localCmd := v.GetStringSlice("localcmd")
	fmt.Println(localCmd)
	if len(localCmd) > 0 {
		if err := localCmdRun(localCmd); err != nil {
			return
		}
	}
	//源文件路径
	filePath := v.GetString("upload.srcfile")
	fmt.Println(filePath)
	//目标路径
	dstPath := v.GetString("upload.dstpath")
	runCommand(v.GetStringSlice("cmd"), filePath, dstPath)
	fmt.Println("tar xzvf success")
}
