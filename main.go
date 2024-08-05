package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func main(){
	server := "150.158.140.243:22"
	user := "root"
	password := "921226LWJlwj"

	//SSH配置
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	
	//连接到服务器
	client , err := ssh.Dial("tcp",server,config)
    if err != nil {
		log.Fatalf("连接失败: %v",err)
	}
	defer client.Close()

	//打开一个会话
	seesion, err := client.NewSession()
	if err != nil {
		log.Fatalf("创建会话失败: %v",err)
	
	}
	defer client.Close()

	//执行远程命令
	output,err := seesion.CombinedOutput("ls")
	if err != nil {
		log.Fatalf("执行命令失败: %v",err)
	
	}
	 //打印命令输出
	 fmt.Println(string(output))
}