//Author  :     knight
//Date    :     2020/07/12 13:26:08
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     使用阿里云oss保存图片

package dao

import (
	"bufio"
	"log"
	"os"
	"holl/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var client *oss.Client
var bucket *oss.Bucket

//初始化oss连接
func InitOss() error {
	var err error

	endPoint := "http://oss-cn-beijing.aliyuncs.com"
	accessKey, accessSecret := readLine()

	client, err = oss.New(endPoint, accessKey, accessSecret)
	if err != nil {
		log.Println(err)
		return err
	}

	bucket, err = client.Bucket("holl")
	if err != nil {
		log.Println(err)
		return err
	}
	
	log.Println("Oss Init Success")
	return nil
}

func readLine() (string, string) {
	strings := make([]string, 0)
	
	f, err := os.Open(config.GetOssConfigFilePath())
	
	if err != nil {
		return "", ""
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		strings = append(strings, s.Text())
	}

	return strings[0], strings[1]
}

//OssUploadFile 上传oss上的图片
func OssUploadFile(objectName, localFileName string) error {
	err := bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		return err
	}
	log.Println("上传成功")
	return nil
}

//OssDownloadFile 下载oss上的图片
func OssDownloadFile(objectName, localFileName string) error {
	err := bucket.GetObjectToFile(objectName, localFileName)
	if err != nil {
		return err
	}
	log.Println("下载成功")
	return nil
}

//OssDeleteFile 删除oss上的图片
func OssDeleteFile(objectName string) error {
	err := bucket.DeleteObject(objectName)
	if err != nil {
		return err
	}
	log.Printf("删除%s成功\n", objectName)
	return nil
}
