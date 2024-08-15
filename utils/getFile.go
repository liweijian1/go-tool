package utils

import (
	"github.com/pkg/sftp"
	"os"
	"io"
	"fmt"
	"log"
	"golang.org/x/crypto/ssh"
	"path/filepath"
)

//下载单个文件
func GetFile(clientInterface *ssh.Client,remotePath,localPath string) {
	//类型断言
//   client,ok := clientInterface.(*ssh.Client)
//    if !ok {
// 	fmt.Errorf("client is not of type *ssh.Client")
//    }

	//创建sftp客户端
	sftpClient,err := sftp.NewClient(clientInterface)
	if err != nil {
		log.Fatalf("Failed to create SFTP client: %s", err)
	}
	defer sftpClient.Close()

	//打开远程文件
	remoteFilePath := remotePath
	remoteFile,err := sftpClient.Open(remoteFilePath)
	if err != nil {
		log.Fatalf("Failed to open remote file: %s", err)
	}
	defer remoteFile.Close()

	//创建本地文件
	localFilePath := localPath
	localFile,err := os.Create(localFilePath)
	if err != nil {
		log.Fatalf("Failed to create local file: %s", err)
	}
	defer localFile.Close()

	//复制远程文件
	if _,err := io.Copy(localFile,remoteFile);err != nil {
		log.Fatalf("Failed to copy file content: %s", err)
	}

	fmt.Println("文件保存成功")

}

//下载文件夹
func downloadFolder(client *ssh.Client,remotePath,localPath string){
	//确认本地目录存在
	if err := os.MkdirAll(localPath,os.ModePerm); err != nil {
		log.Fatalf("Failed to create local directory: %s", err)
	}
	sftpClient,err := sftp.NewClient(client)
	//打开远程目录
	remoteDir , err := sftpClient.ReadDir(remotePath)
    if err != nil {
		log.Fatalf("Failed to open remote directory: %s", err)
	}
	//遍历远程目录
	for _,file := range remoteDir {
		remoteFilePath := filepath.Join(remotePath,file.Name())
		localFilePath := filepath.Join(localPath,file.Name())

		if file.IsDir() {
			//递归子文件夹
			downloadFolder(client,remoteFilePath,localFilePath)

		}else {
			//下载文件
			GetFile(client,remoteFilePath,localFilePath)
		
		}
	}
}
