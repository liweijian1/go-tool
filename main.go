package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
)

func main(){
	server := "150.158.140.243:22"
	user := "root"
	password := "921226LWJlwj"

	//SSH配置
	config := &ssh.config{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	
	//连接到服务器
	client , err := ssh.Detail("tcp",server,config)
    if err != nill {
		
	}

}