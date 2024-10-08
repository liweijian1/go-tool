package utils

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
	"log"
	"net/http"
	"strings"
)

func Connaction(c *gin.Context, server,user, password string) {
	// user := "root"
	// password := "921226LWJlwj"
	//SSH配置
	config := &ssh.ClientConfig{
		User: strings.TrimSpace(user),
		Auth: []ssh.AuthMethod{
			ssh.Password(strings.TrimSpace(password)),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	//连接到服务器
	client, err := ssh.Dial("tcp", server, config)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "连接失败",
			"error":   err.Error(),
		})
		log.Println("连接失败: %v", err)
		return

	}
	defer client.Close()

	//打开一个会话
	// seesion, err := client.NewSession()
	// if err != nil {
	// 	log.Fatalf("创建会话失败: %v", err)

	// }
	// defer client.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": "连接成功",
	})

	// remotePath := "/data/frontend/dist"
	// localPath := "/Users/doususu/frontend/dist"
	// downloadFolder(client,remotePath,localPath)

	// //执行远程命令
	// output, err := seesion.CombinedOutput(GetCli())
	// if err != nil {
	// 	log.Fatalf("执行命令失败: %v", err)

	// }
	// //打印命令输出
	// fmt.Println(string(output))
}
